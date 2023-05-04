package storage

import (
    "fmt"
//     "strconv"
)

type Counter int64
type Gauge float64

type MemStorage struct {
    Data map[string]interface{}
}

type Storage interface {
    Update(mtype string, metric string, value interface{}) error
    Get(metric string) error
}

func (m *MemStorage) Get(metric string) (interface{}, error) {
    if v, ok := m.Data[metric]; ok {
        return v, nil
    }
    return "No such metric in MemStorage", fmt.Errorf("Not Found")

}

func (m *MemStorage) Update(mtype string, metric string, value interface{}) error {
    switch mtype {
        case "counter":
            if m.Data[metric] == nil {
                m.Data[metric] = value
                return nil
            }
            m.Data[metric] = m.Data[metric].(Counter) + value.(Counter)
        case "gauge":
            m.Data[metric] = value
        default:
            return fmt.Errorf("Wrong type of metric")
    }
    return nil
}
