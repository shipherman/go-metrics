package storage

import (
    "fmt"
)

type Counter int64
type Gauge float64

type MemStorage struct {
    Data map[string]interface{}
}

type Storage interface {
    Update(mtype string, metric string, value interface{}) error
    Get(metric string) error
    GetAll() map[string]interface{}
}


func (m *MemStorage) Get(metric string) (interface{}, error) {
    if v, ok := m.Data[metric]; ok {
        return v, nil
    }
    return "No such metric in memstorage", fmt.Errorf("Not found")

}

func (m *MemStorage) GetAll() map[string]interface{} {
    return m.Data
}

func (m *MemStorage) Update(mtype string, metric string, value interface{}) error {
    switch mtype {
        case "counter":
            if m.Data[metric] == nil {
                m.Data[metric] = value.(Counter)
                return nil
            }
            m.Data[metric] = m.Data[metric].(Counter) + value.(Counter)
        case "gauge":
            m.Data[metric] = value.(Gauge)
        default:
            panic(fmt.Errorf("Wrong type of metric"))
    }
    return nil
}
