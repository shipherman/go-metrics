package main

import (
	"context"

	protometrics "github.com/shipherman/go-metrics/internal/grpcapi/protometrics/v1/gen"
	"github.com/shipherman/go-metrics/internal/storage"
)

func SendGRPC(m storage.MemStorage) error {
	client := protometrics.NewMetricsServiceClient(ConnGrpc)

	for k, v := range m.CounterData {
		_, err := client.UpdateCounter(context.Background(), &protometrics.UpdateCounterRequest{
			Counter: &protometrics.Counter{
				Name:  k,
				Delta: uint32(v),
			},
		})
		if err != nil {
			return err
		}
	}

	for k, v := range m.GaugeData {
		_, err := client.UpdateGauge(context.Background(), &protometrics.UpdateGaugeRequest{
			Gauge: &protometrics.Gauge{
				Name:  k,
				Value: float64(v),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
