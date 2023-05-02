package storage

import (
    "runtime"
    "math/rand"
)

type Counter int64
type Gauge float64
type MemStorage struct {
    CounterData map[string]Counter
    GaugeData map[string]Gauge
}



func UpdateMemStorage (m *MemStorage) {
    var stat runtime.MemStats
    runtime.ReadMemStats(&stat)
    m.GaugeData["Alloc"] = Gauge(stat.Alloc)
    m.GaugeData["BuckHashSys"] = Gauge(stat.BuckHashSys)
    m.GaugeData["Frees"] = Gauge(stat.Frees)
    m.GaugeData["GCCPUFraction"] = Gauge(stat.GCCPUFraction)
    m.GaugeData["GCSys"] = Gauge(stat.GCSys)
    m.GaugeData["HeapAlloc"] = Gauge(stat.HeapAlloc)
    m.GaugeData["HeapIdle"] = Gauge(stat.HeapIdle)
    m.GaugeData["HeapInuse"] = Gauge(stat.HeapInuse)
    m.GaugeData["HeapObjects"] = Gauge(stat.HeapObjects)
    m.GaugeData["HeapReleased"] = Gauge(stat.HeapReleased)
    m.GaugeData["HeapSys"] = Gauge(stat.HeapSys)
    m.GaugeData["LastGC"] = Gauge(stat.LastGC)
    m.GaugeData["Lookups"] = Gauge(stat.Lookups)
    m.GaugeData["MCacheInuse"] = Gauge(stat.MCacheInuse)
    m.GaugeData["MCacheSys"] = Gauge(stat.MCacheSys)
    m.GaugeData["MSpanInuse"] = Gauge(stat.MSpanInuse)
    m.GaugeData["MSpanSys"] = Gauge(stat.MSpanSys)
    m.GaugeData["Mallocs"] = Gauge(stat.Mallocs)
    m.GaugeData["NextGC"] = Gauge(stat.NextGC)
    m.GaugeData["NumForcedGC"] = Gauge(stat.NumForcedGC)
    m.GaugeData["NumGC"] = Gauge(stat.NumGC)
    m.GaugeData["OtherSys"] = Gauge(stat.OtherSys)
    m.GaugeData["PauseTotalNs"] = Gauge(stat.PauseTotalNs)
    m.GaugeData["StackInuse"] = Gauge(stat.StackInuse)
    m.GaugeData["StackSys"] = Gauge(stat.StackSys)
    m.GaugeData["Sys"] = Gauge(stat.Sys)
    m.GaugeData["TotalAlloc"] = Gauge(stat.TotalAlloc)
    m.GaugeData["RandomValue"] = Gauge(rand.Float32())
    m.CounterData["PollCount"] += 1
}
