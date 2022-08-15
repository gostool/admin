package status

import (
	"admin/internal/consts"
	"admin/internal/model"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"time"
)

type sStatus struct {
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterStatus(instance)
}

func New() *sStatus {
	return &sStatus{}
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// GetServerInfo 获取服务器信息
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *sStatus) GetServerInfo(ctx context.Context) (data *model.Server, err error) {
	server := model.Server{Os: s.InitOs(ctx)}
	if server.Cpu, err = s.InitCpu(ctx); err != nil {
		return nil, errors.Wrap(err, "获取CPU信息失败!")
	}
	if server.Rrm, err = s.InitRAM(ctx); err != nil {
		return nil, errors.Wrap(err, "获取ARM信息失败!")
	}
	if server.Disk, err = s.InitDisk(ctx); err != nil {
		return nil, errors.Wrap(err, "获取硬盘信息失败!")
	}

	return &server, nil
}

// InitOs 获取系统信息 组装数据为 response.Os
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *sStatus) InitOs(ctx context.Context) model.Os {
	return model.Os{
		GOOS:         runtime.GOOS,
		NumCPU:       runtime.NumCPU(),
		Compiler:     runtime.Compiler,
		GoVersion:    runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
	}
}

// InitCpu 获取CPU信息 组装数据为 response.Cpu
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *sStatus) InitCpu(ctx context.Context) (model.Cpu, error) {
	var _cpu model.Cpu
	cores, err := cpu.Counts(false)
	if err != nil {
		return _cpu, err
	}
	_cpu.Cores = cores
	_cpu.Cpus, err = cpu.Percent(time.Duration(200)*time.Millisecond, true)
	if err != nil {
		return _cpu, err
	}
	return _cpu, nil
}

// InitRAM 获取ARM信息 组装数据为 response.Rrm
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *sStatus) InitRAM(ctx context.Context) (model.Rrm, error) {
	var arm model.Rrm
	virtualMemoryStat, err := mem.VirtualMemory()
	if err != nil {
		return arm, err
	}
	arm.UsedMB = int(virtualMemoryStat.Used) / MB
	arm.TotalMB = int(virtualMemoryStat.Total) / MB
	arm.UsedPercent = int(virtualMemoryStat.UsedPercent)
	return arm, nil
}

// InitDisk 获取硬盘信息 组装数据为 response.Disk
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *sStatus) InitDisk(ctx context.Context) (model.Disk, error) {
	var _disk model.Disk
	usageStat, err := disk.Usage("/")
	if err != nil {
		return _disk, err
	}
	_disk.UsedMB = int(usageStat.Used) / MB
	_disk.UsedGB = int(usageStat.Used) / GB
	_disk.TotalMB = int(usageStat.Total) / MB
	_disk.TotalGB = int(usageStat.Total) / GB
	_disk.UsedPercent = int(usageStat.UsedPercent)
	return _disk, nil
}
