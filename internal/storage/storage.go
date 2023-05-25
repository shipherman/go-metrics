package storage

import (
    "fmt"
)

type Counter int64
type Gauge float64

type MemStorage struct {
    CounterData map[string]Counter
    GaugeData map[string]Gauge
}


func New() (MemStorage) {
    return MemStorage{
        CounterData: map[string]Counter{},
        GaugeData: map[string]Gauge{},
    }
}

func (m *MemStorage) Get(metric string) (interface{}, error) {
    if v, ok := m.CounterData[metric]; ok {
        return v, nil
    }
    if v, ok := m.GaugeData[metric]; ok {
        return v, nil
    }
    return "No such metric in memstorage", fmt.Errorf("metric not found")

}

func (m *MemStorage) GetAllCounters() map[string]Counter {
    return m.CounterData
}

func (m *MemStorage) GetAllGauge() map[string]Gauge {
    return m.GaugeData
}

func (m *MemStorage) UpdateGauge(metric string, value Gauge) {
    m.GaugeData[metric] = value
}

func (m *MemStorage) UpdateCounter(metric string, value Counter) {
//     if m.CounterData[metric] == nil {
//         m.CounterData[metric] = value
//         return
//     }
    m.CounterData[metric] = m.CounterData[metric] + value
}
