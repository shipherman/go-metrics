package grpcservice

import (
	"context"

	pb "github.com/shipherman/go-metrics/internal/grpcapi/protometrics/v1/gen"

	"github.com/shipherman/go-metrics/internal/storage"
)

type GServiceServer struct {
	pb.UnimplementedMetricsServiceServer
	Storage *storage.MemStorage
}

func (gs *GServiceServer) GetGauge(ctx context.Context, in *pb.GaugeRequest) (*pb.GaugeResponse, error) {
	var response pb.GaugeResponse
	m, err := gs.Storage.Get(in.Gauge.Name)
	if err != nil {
		return &response, err
	}
	g := pb.Gauge{
		Name:  in.Gauge.Name,
		Value: float64(m.(storage.Gauge)),
	}
	response.Gauge = &g
	return &response, nil
}

func (gs *GServiceServer) UpdateGauge(ctx context.Context, in *pb.GaugeRequest) (*pb.GaugeResponse, error) {
	var response pb.GaugeResponse
	var g pb.Gauge

	gs.Storage.UpdateGauge(in.Gauge.Name, storage.Gauge(in.Gauge.Value))
	res, err := gs.Storage.Get(in.Gauge.Name)
	if err != nil {
		return &response, err
	}

	g.Name = in.Gauge.Name
	g.Value = float64(res.(storage.Gauge))

	response.Gauge = &g
	return &response, nil
}

func (gs *GServiceServer) GetCounter(ctx context.Context, in *pb.CounterRequest) (*pb.CounterResponse, error) {
	var response pb.CounterResponse
	m, err := gs.Storage.Get(in.Counter.Name)
	if err != nil {
		return &response, err
	}
	g := pb.Counter{
		Name:  in.Counter.Name,
		Delta: uint32(m.(storage.Counter)),
	}
	response.Counter = &g
	return &response, nil
}

func (gs *GServiceServer) UpdateCounter(ctx context.Context, in *pb.CounterRequest) (*pb.CounterResponse, error) {
	var response pb.CounterResponse
	var g pb.Counter

	gs.Storage.UpdateCounter(in.Counter.Name, storage.Counter(in.Counter.Delta))
	res, err := gs.Storage.Get(in.Counter.Name)
	if err != nil {
		return &response, err
	}

	g.Name = in.Counter.Name
	g.Delta = uint32(res.(storage.Counter))

	response.Counter = &g
	return &response, nil
}
