package storage

type Counter int64
type Gauge float64
type MemStorage struct {
    CounterData map[string]Counter
    GaugeData map[string]Gauge
}
