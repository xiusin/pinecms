package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/cmd/version"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/common/helper"
	"runtime"
	"sync"
	"time"
)

var runningStart = time.Now()

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Server struct {
	Nets           []*Net `json:"nets"`
	Os             Os     `json:"os"`
	Cpu            Cpu    `json:"cpu"`
	Rrm            Rrm    `json:"ram"`
	Disk           Disk   `json:"disk"`
	RunningTime    int64  `json:"running_time"`
	StartTime      string `json:"start_time"`
	PineVersion    string `json:"pine_version"`
	PineCmsVersion string `json:"pine_cms_version"`
	XormVersion    string `json:"xorm_version"`
	MysqlVersion   string `json:"mysql_version"`
}

type Net struct {
	BytesRecv uint64 `json:"recv"`
	BytesSent uint64 `json:"send"`
	Name      string `json:"name"`
}

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	Cpus       []float64 `json:"cpus"`
	Cores      int       `json:"cores"`
	CpuPercent []float64 `json:"cpu_percent"`
}

type Rrm struct {
	UsedMB      int `json:"usedMb"`
	TotalMB     int `json:"totalMb"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct {
	UsedMB      int `json:"usedMb"`
	UsedGB      int `json:"usedGb"`
	TotalMB     int `json:"totalMb"`
	TotalGB     int `json:"totalGb"`
	UsedPercent int `json:"usedPercent"`
}

type StatController struct {
	pine.Controller
}

func (_ *StatController) InitOS() (o Os) {
	o.GOOS = runtime.GOOS
	o.NumCPU = runtime.NumCPU()
	o.Compiler = runtime.Compiler
	o.GoVersion = runtime.Version()
	o.NumGoroutine = runtime.NumGoroutine()
	return o
}

func (_ *StatController) InitCPU() (c Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Cpus = cpus
	}
	c.CpuPercent, _ = cpu.Percent(time.Duration(200)*time.Millisecond, false)
	return c, nil
}

func (_ *StatController) InitRAM() (r Rrm, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		r.UsedMB = int(u.Used) / MB
		r.TotalMB = int(u.Total) / MB
		r.UsedPercent = int(u.UsedPercent)
	}
	return r, nil
}

func (_ *StatController) InitDisk() (d Disk, err error) {
	if u, err := disk.Usage("/"); err != nil {
		return d, err
	} else {
		d.UsedMB = int(u.Used) / MB
		d.UsedGB = int(u.Used) / GB
		d.TotalMB = int(u.Total) / MB
		d.TotalGB = int(u.Total) / GB
		d.UsedPercent = int(u.UsedPercent)
	}
	return d, nil
}

func (_ *StatController) InitNet() (useages []*Net, err error) {
	nv, err := net.IOCounters(false)
	if err != nil {
		return
	}
	useages = make([]*Net, 0, len(nv))
	for _, status := range nv {
		useage := &Net{BytesRecv: status.BytesRecv, BytesSent: status.BytesSent, Name: status.Name}
		useages = append(useages, useage)
	}
	return
}

func (stat *StatController) GetData(orm *xorm.Engine, cacher cache.AbstractCache) {
	var s Server

	var wg sync.WaitGroup
	wg.Add(5)
	go func() { defer wg.Done(); s.Os = stat.InitOS() }()
	go func() { defer wg.Done(); s.Cpu, _ = stat.InitCPU() }()
	go func() { defer wg.Done(); s.Rrm, _ = stat.InitRAM() }()
	go func() { defer wg.Done(); s.Disk, _ = stat.InitDisk() }()
	go func() { defer wg.Done(); s.Nets, _ = stat.InitNet() }()
	wg.Wait()

	s.RunningTime = int64(time.Now().Sub(runningStart).Minutes())
	s.StartTime = runningStart.Format(helper.TimeFormat)
	s.PineVersion = pine.Version
	s.PineCmsVersion = version.Version
	s.XormVersion = xorm.Version

	cacher.Remember(controllers.CacheMysqlVersion, &s.MysqlVersion, func() (interface{}, error) {
		if version, err := orm.QueryString("SELECT VERSION() AS version"); err == nil {
			return version[0]["version"], nil
		} else {
			return "", err
		}
	}, 24*3600)
	helper.Ajax(s, 0, stat.Ctx())
}
