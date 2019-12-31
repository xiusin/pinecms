package cmd

import (
	"context"
	"fmt"
	"github.com/xiusin/iriscms/cmd/util"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/fatih/color"
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
	buildName         = "iriscms-dev-build"
	delay, limit      int32
	watcher           *fsnotify.Watcher
	counter           int32
)

func init() {
	serveCmd.AddCommand(devCmd)
	devCmd.Flags().StringSlice("ignoreDirs", []string{"vendor", ".git", ".idea", "node_modules"}, "忽略变动监听的目录")
	devCmd.Flags().StringSlice("types", []string{".go"}, "需要监听的文件类型, .*为监听任意文件")
	devCmd.Flags().String("root", util.AppPath()+"/src", "监听的根目录")
	devCmd.Flags().Int32("delay", 2, "每次构建进程的延迟时间单位：秒")
	devCmd.Flags().Int32("limit", 500, "监听文件的最大数量")
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
}

func devCommand(cmd *cobra.Command, args []string) error {
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
		return err
	}
	go eventNotify()
	go serve()
	select {}
}

func serve() {
	var nextLoop = make(chan struct{})
	for {
		ctx, cancel := context.WithCancel(context.Background())
		cmd := exec.CommandContext(ctx, fmt.Sprintf("./%s", buildName), "serve", "start", "--banner=false")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		go func() {
			<-rebuildNotifier
			cancel()
			nextLoop <- struct{}{}
		}()
		if err := cmd.Start(); err != nil {
			log.Println(fmt.Sprintf("%s %s", color.RedString("[ERRO]"), err.Error()))
		}
		log.Println(fmt.Sprintf("%s %s", color.YellowString("[INFO]"), "构建执行文件, 并且启动服务成功"))
		if err := cmd.Wait(); err != nil && err.Error() != "signal: killed" {
			log.Println(fmt.Sprintf("%s %s", color.RedString("[ERRO]"), err.Error()))
		}
		<-nextLoop
	}
}
func build() error {
	cmd := exec.Command("go", "build", "-o", buildName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Dir = util.AppPath()
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func registerFileToWatcher() error {
	files, err := util.ScanDir(rootDir, ignoreDirs)
	if err != nil {
		return err
	}
	for _, file := range files {
		if counter > limit {
			log.Println(fmt.Sprintf("%s %s", color.RedString("[ERRO]"), "监听文件已达上限"))
			break
		}
		if len(types) > 0 && !util.InSlice(".*", types) && !file.IsDir {
			suffixPartial := strings.Split(file.Path, ".")
			if len(suffixPartial) > 2 && !util.InSlice("."+suffixPartial[len(suffixPartial)-1], types) {
				continue
			}
		}
		// 忽略构建生成的文件
		if !file.IsDir && strings.HasSuffix(file.Path, buildName) {
			continue
		}
		if err := watcher.Add(file.Path); err != nil {
			return err
		} else {
			if !file.IsDir {
				atomic.AddInt32(&counter, 1)
				log.Println(fmt.Sprintf("%s %s", color.YellowString("[WATC]"), strings.Replace(file.Path, util.AppPath(), "", 1)))
			}
		}
	}
	return nil
}

func isIgnoreAction(event *fsnotify.Event) bool {
	// 忽略jb的临时文件, 以及修改文件权限的动作
	return strings.HasSuffix(event.Name, "__") || event.Op.String() == "CHMOD"
}

func eventNotify() {
	var lockerTimestamp int64
	var building = false
	for {
		select {
		case event, _ := <-watcher.Events:
			if time.Now().Unix()-lockerTimestamp > int64(delay) && !building {
				if isIgnoreAction(&event) {
					continue
				}
				lockerTimestamp, building = time.Now().Unix(), true
				name := util.Replace(event.Name, util.AppPath(), "")
				if event.Op&fsnotify.Create == fsnotify.Create {
					_ = watcher.Add(event.Name)
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					_ = watcher.Remove(event.Name)
				}
				log.Println(fmt.Sprintf("%s %s event %s", color.YellowString("[EVEN]"), name, strings.ToLower(event.Op.String())))
				if err := build(); err != nil {
					log.Println(fmt.Sprintf("\n%s build err", color.RedString("[ERRO]")))
					building = false
					continue
				}
				rebuildNotifier <- struct{}{}
				building = false
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println(fmt.Sprintf("%s watcher error: %s", color.RedString("[ERRO]"), err.Error()))
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
