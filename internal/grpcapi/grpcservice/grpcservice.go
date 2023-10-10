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

func (gs *GServiceServer) GetGauge(ctx context.Context, in *pb.GetGaugeRequest) (*pb.GetGaugeResponse, error) {
	var response pb.GetGaugeResponse
	m, err := gs.Storage.Get(in.Name)
	if err != nil {
		return &response, err
	}
	g := pb.Gauge{
		Name:  in.Name,
		Value: float64(m.(storage.Gauge)),
	}
	response.Gauge = &g
	return &response, nil
}

func (gs *GServiceServer) UpdateGauge(ctx context.Context, in *pb.UpdateGaugeRequest) (*pb.UpdateGaugeResponse, error) {
	var response pb.UpdateGaugeResponse
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

func (gs *GServiceServer) GetCounter(ctx context.Context, in *pb.GetCounterRequest) (*pb.GetCounterResponse, error) {
	var response pb.GetCounterResponse
	m, err := gs.Storage.Get(in.Name)
	if err != nil {
		return &response, err
	}
	g := pb.Counter{
		Name:  in.Name,
		Delta: uint32(m.(storage.Counter)),
	}
	response.Counter = &g
	return &response, nil
}

func (gs *GServiceServer) UpdateCounter(ctx context.Context, in *pb.UpdateCounterRequest) (*pb.UpdateCounterResponse, error) {
	var response pb.UpdateCounterResponse
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
