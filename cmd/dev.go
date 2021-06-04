package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/xiusin/logger"
	"github.com/xiusin/pinecms/cmd/util"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var (
	devCmd = &cobra.Command{
		Use:   "dev",
		Short: "启动开发服务器",
		RunE:  devCommand,
	}
	rebuildNotifier   = make(chan struct{})
	types, ignoreDirs []string
	rootDir           string
	buildName         = "pinecms-dev-build"
	delay, limit      int32
	watcher           *fsnotify.Watcher
	counter           int32
)

func init() {
	serveCmd.AddCommand(devCmd)
	devCmd.Flags().StringSlice("ignoreDirs", []string{"vendor", ".git", ".idea", "node_modules"}, "忽略变动监听的目录")
	devCmd.Flags().StringSlice("types", []string{".go", ".yml"}, "需要监听的文件类型, .*为监听任意文件")
	devCmd.Flags().String("root", util.AppPath(), "监听的根目录")
	devCmd.Flags().Int32("delay", 3, "每次构建进程的延迟时间单位：秒")
	devCmd.Flags().Int32("limit", 500, "监听文件的最大数量")
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
}

func devCommand(cmd *cobra.Command, args []string) error {
	closeCh := make(chan os.Signal)
	signal.Notify(closeCh, os.Interrupt, os.Kill)
	if runtime.GOOS == "windows" {
		buildName += ".exe"
	}
	_ = os.Remove(buildName)
	defer func() { _ = watcher.Close() }()
	if err := getCommandFlags(cmd); err != nil {
		return err
	}
	if err := build(); err != nil {
		return err
	}
	if err := registerFileToWatcher(); err != nil {
		panic(err)
	}
	go eventNotify()
	go serve()
	<-closeCh
	return nil
}

func serve() {
	var nextEventCh = make(chan struct{})
	for {
		ctx, cancel := context.WithCancel(context.Background())
		process := exec.CommandContext(ctx, fmt.Sprintf("./%s", buildName), "serve", "start", "--banner=false")
		process.Stdout = os.Stdout
		process.Stderr = os.Stdout
		go func() {
			<-rebuildNotifier
			cancel()
			nextEventCh <- struct{}{}
		}()
		if err := process.Start(); err != nil {
			logger.Error(err)
		}
		excludeErrors := "signal:"
		if err := process.Wait(); err != nil && !strings.Contains(err.Error(), excludeErrors) {
			logger.Error(err)
		}
		<-nextEventCh
	}
}
func build() error {
	start := time.Now()
	cmd := exec.Command("go", "build", "-o", buildName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Dir = util.AppPath()
	if err := cmd.Run(); err != nil {
		return err
	}

	logger.Print("构建耗时: " + time.Now().Sub(start).String())

	return nil
}

func registerFileToWatcher() error {
	files, err := util.ScanDir(rootDir, ignoreDirs)
	if err != nil {
		return err
	}
	for _, file := range files {
		if counter > limit {
			logger.Error("监听文件已达上限")
			break
		}
		if len(types) > 0 && !util.InSlice(".*", types) && !file.IsDir {
			suffixPartial := strings.Split(file.Path, ".")
			if !util.InSlice("."+suffixPartial[len(suffixPartial)-1], types) {
				continue
			}
		}
		// 忽略构建生成的文件
		if !file.IsDir && strings.HasSuffix(file.Path, buildName) {
			continue
		}
		if err := watcher.Add(file.Path); err != nil {
			return err
		} else if !file.IsDir {
			atomic.AddInt32(&counter, 1)
		}
	}
	return nil
}

func isIgnoreAction(event *fsnotify.Event) bool {
	// 忽略jb的临时文件, 以及修改文件权限的动作
	return strings.HasSuffix(event.Name, "__") || event.Op.String() == "CHMOD"
}

func eventNotify() {
	var lockerTimestamp time.Time
	var building = false
	for {
		select {
		case event, _ := <-watcher.Events:
			if isIgnoreAction(&event) {
				continue
			}
			if time.Now().Sub(lockerTimestamp) > time.Second*time.Duration(delay) && !building {
				name := util.Replace(event.Name, util.AppPath(), "")
				fileInfo := strings.Split(name, ".")
				if !util.InSlice(".*", types) && !util.InSlice("."+strings.TrimRight(fileInfo[len(fileInfo)-1], "~"), types) {
					continue
				}
				lockerTimestamp, building = time.Now(), true
				if event.Op&fsnotify.Create == fsnotify.Create {
					_ = watcher.Add(event.Name)
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					_ = watcher.Remove(event.Name)
				}
				logger.Warningf("%s event %s", name, strings.ToLower(event.Op.String()))
				go func() {
					if err := build(); err != nil {
						logger.Error("构建错误", err)
						building = false
					}
					rebuildNotifier <- struct{}{}
					building = false
				}()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logger.Error("watcher error: %s", err)
		}
	}
}

func getCommandFlags(cmd *cobra.Command) (err error) {
	ignoreDirs, err = cmd.Flags().GetStringSlice("ignoreDirs")
	if err != nil {
		return
	}
	types, err = cmd.Flags().GetStringSlice("types")
	if err != nil {
		return
	}
	rootDir, err = cmd.Flags().GetString("root")
	if err != nil {
		return
	}
	delay, err = cmd.Flags().GetInt32("delay")
	if err != nil {
		return
	}
	limit, err = cmd.Flags().GetInt32("limit")
	if err != nil {
		return
	}
	return
}
