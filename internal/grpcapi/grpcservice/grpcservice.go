package grpcservice

import (
	"context"

	pb "github.com/shipherman/go-metrics/internal/grpcapi/protometrics/v1/gen"

	"github.com/bufbuild/protovalidate-go"
	"github.com/shipherman/go-metrics/internal/storage"
)

type GServiceServer struct {
	pb.UnimplementedMetricsServiceServer
	Storage *storage.MemStorage
}

func (gs *GServiceServer) GetGauge(ctx context.Context, in *pb.GaugeRequest) (*pb.GaugeResponse, error) {
	var response pb.GaugeResponse

	// create validator
	v, err := protovalidate.New()
	if err != nil {
		return &response, err
	}

	m, err := gs.Storage.Get(in.Gauge.Name)
	if err != nil {
		return &response, err
	}
	g := pb.Gauge{
		Name:  in.Gauge.Name,
		Value: float64(m.(storage.Gauge)),
	}
	response.Gauge = &g

	if err = v.Validate(&response); err != nil {
		return &response, err
	}
	return &response, nil
}

func (gs *GServiceServer) UpdateGauge(ctx context.Context, in *pb.GaugeRequest) (*pb.GaugeResponse, error) {
	var response pb.GaugeResponse
	var g pb.Gauge

	// create validator
	v, err := protovalidate.New()
	if err != nil {
		return &response, err
	}

	gs.Storage.UpdateGauge(in.Gauge.Name, storage.Gauge(in.Gauge.Value))
	res, err := gs.Storage.Get(in.Gauge.Name)
	if err != nil {
		return &response, err
	}

	g.Name = in.Gauge.Name
	g.Value = float64(res.(storage.Gauge))

	response.Gauge = &g

	if err = v.Validate(&response); err != nil {
		return &response, err
	}

	return &response, nil
}

func (gs *GServiceServer) GetCounter(ctx context.Context, in *pb.CounterRequest) (*pb.CounterResponse, error) {
	var response pb.CounterResponse

	// create validator
	v, err := protovalidate.New()
	if err != nil {
		return &response, err
	}

	m, err := gs.Storage.Get(in.Counter.Name)
	if err != nil {
		return &response, err
	}
	g := pb.Counter{
		Name:  in.Counter.Name,
		Delta: uint32(m.(storage.Counter)),
	}
	response.Counter = &g

	if err = v.Validate(&response); err != nil {
		return &response, err
	}

	return &response, nil
}

func (gs *GServiceServer) UpdateCounter(ctx context.Context, in *pb.CounterRequest) (*pb.CounterResponse, error) {
	var response pb.CounterResponse
	var g pb.Counter

	// create validator
	v, err := protovalidate.New()
	if err != nil {
		return &response, err
	}

	gs.Storage.UpdateCounter(in.Counter.Name, storage.Counter(in.Counter.Delta))
	res, err := gs.Storage.Get(in.Counter.Name)
	if err != nil {
		return &response, err
	}

	g.Name = in.Counter.Name
	g.Delta = uint32(res.(storage.Counter))

	response.Counter = &g

	if err = v.Validate(&response); err != nil {
		return &response, err
	}

	return &response, nil
}
