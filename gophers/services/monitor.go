package main

import (
	"os"
	"runtime"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/laher/gophertron/gophers"
	"github.com/thoas/stats"
)

type MonitoringService struct {
	WebStats            *stats.Stats
	IntervalSeconds     int
	IsMonitoring        bool
	timeout             time.Duration
	runtimeStatsRunning bool
	hostname            string
}

func (ms *MonitoringService) RecordStartup() {
	ms.Init()
	maxprocs := runtime.GOMAXPROCS(0)

	ms.Log().Debugf("[REPORT] Max procs: %d. NumCPU: %d", maxprocs, runtime.NumCPU())
	goroutines := runtime.NumGoroutine()
	ms.Log().Debugf("[REPORT] Goroutines: %d", goroutines)

}

func (ms *MonitoringService) Init() {
	hs, err := os.Hostname()
	if err != nil {
		ms.Log().Errorf("Unknown hostname. Error: %v", err)
		hs = "unknown-host"
	}
	ms.hostname = hs
}

func (ms *MonitoringService) Monitor() {
	if ms.IntervalSeconds == 0 {
		ms.IntervalSeconds = 60 //1 minutes
	}
	if ms.runtimeStatsRunning {
		ms.Log().Errorf("Runtime stats is already running\n")
		return
	}

	ms.runtimeStatsRunning = true
	go gophers.Recover(ms.reportRuntimeStats, "Monitor")
}

type SnapshotReport struct {
	NumGoroutine           int
	MemoryAllocated        uint64
	MemoryAllocations      uint64
	MemoryFrees            uint64
	MemoryGcTotalPause     float64
	MemoryHeap             uint64
	MemoryStackInuse       uint64
	MemoryGcPausePerSecond float64
	MemoryGcPerSecond      float64
}

func (ms *MonitoringService) Log() *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"host": ms.hostname,
		"cat":  "MONITOR",
	})
}

func (ms *MonitoringService) reportRuntimeStats() {

	sleep := time.Duration(ms.IntervalSeconds) * time.Second
	memStats := &runtime.MemStats{}
	for ms.runtimeStatsRunning {
		runtime.ReadMemStats(memStats)
		report := &SnapshotReport{}
		report.NumGoroutine = runtime.NumGoroutine()
		report.MemoryAllocated = memStats.Alloc
		report.MemoryAllocations = memStats.Mallocs
		report.MemoryFrees = memStats.Frees
		report.MemoryGcTotalPause = float64(memStats.PauseTotalNs) / float64(time.Millisecond)
		report.MemoryHeap = memStats.HeapAlloc
		report.MemoryStackInuse = memStats.StackInuse

		//just log for now. TODO: store in the database
		ms.Log().Debugf("[REPORT:ProcStats] %+v", report)
		ms.Log().Debugf("[REPORT:WebStats] %+v", ms.WebStats.Data())

		time.Sleep(sleep)
	}
}
