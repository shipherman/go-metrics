// Implements storing data in RAM
package storage

import (
	"fmt"
)

type Counter int64
type Gauge float64

type MemStorage struct {
	CounterData map[string]Counter
	GaugeData   map[string]Gauge
}

func New() MemStorage {
	return MemStorage{
		CounterData: map[string]Counter{},
		GaugeData:   map[string]Gauge{},
	}
}

// Get exact metric from memory storage
// To Do:
// Refactor to not to return interface{}
func (m *MemStorage) Get(metric string) (interface{}, error) {
	if v, ok := m.CounterData[metric]; ok {
		return v, nil
	}
	if v, ok := m.GaugeData[metric]; ok {
		return v, nil
	}
	return "No such metric in memstorage", fmt.Errorf("metric not found")

}

// Return all metrics that have type Counter
func (m *MemStorage) GetAllCounters() map[string]Counter {
	return m.CounterData
}

// Return metrics that have type Gauge
func (m *MemStorage) GetAllGauge() map[string]Gauge {
	return m.GaugeData
}

// Update value for gauge metric
func (m *MemStorage) UpdateGauge(metric string, value Gauge) {
	m.GaugeData[metric] = value
}

// Increment Counter metric with provided value
func (m *MemStorage) UpdateCounter(metric string, value Counter) {
	m.CounterData[metric] = m.CounterData[metric] + value
}
