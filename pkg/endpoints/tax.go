package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/tscolari/mservice/pkg/services"
	"golang.org/x/net/context"
)

type AddRequest struct {
	Value float64 `json:"value"`
}

type AddResponse struct {
	Value float64 `json:"value"`
	Err   error   `json:"-"`
}

type SubRequest struct {
	Value float64 `json:"value"`
}

type SubResponse struct {
	Value float64 `json:"value"`
	Err   error   `json:"-"`
}

type Tax struct {
	AddEndpoint endpoint.Endpoint
	SubEndpoint endpoint.Endpoint
}

// NewTax returns the Tax service endpoint, that wraps the real service
// implementation and wires any existing middleware
func NewTax(logger log.Logger, svc services.Tax) Tax {
	return Tax{
		AddEndpoint: loggingMiddleware(log.With(logger, "method", "add"))(makeTaxAddEndpoint(svc)),
		SubEndpoint: loggingMiddleware(log.With(logger, "method", "sub"))(makeTaxSubEndpoint(svc)),
	}
}

// Add returns the value with taxes added to it
func (t Tax) Add(ctx context.Context, value float64) (float64, error) {
	resp, err := t.AddEndpoint(ctx, AddRequest{Value: value})
	if err != nil {
		return 0, err
	}
	response := resp.(AddResponse)
	return response.Value, response.Err
}

// Sub return the value with the taxes subtracted from it
func (t Tax) Sub(ctx context.Context, value float64) (float64, error) {
	resp, err := t.SubEndpoint(ctx, SubRequest{Value: value})
	if err != nil {
		return 0, err
	}
	response := resp.(SubResponse)
	return response.Value, response.Err
}

func makeTaxAddEndpoint(svc services.Tax) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		value, err := svc.Add(ctx, req.Value)
		if err != nil {
			return AddResponse{0, err}, nil
		}

		return AddResponse{value, nil}, nil
	}
}

func makeTaxSubEndpoint(svc services.Tax) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SubRequest)
		value, err := svc.Sub(ctx, req.Value)
		if err != nil {
			return SubResponse{0, err}, nil
		}

		return SubResponse{value, nil}, nil
	}
}
