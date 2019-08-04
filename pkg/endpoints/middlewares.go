package endpoints

import (
	"math/rand"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)

type contextKey int

const (
	keyLogger contextKey = iota
)

// loggingMiddleware wraps each request with logging instrumentation
// also makes the logger available down the chain
func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	rand := rand.New(rand.NewSource(time.Now().Unix()))

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger = log.With(logger, "request_id", rand.Int())
			logger.Log("event", "method.started")
			ctx = context.WithValue(ctx, keyLogger, logger)

			defer func(begin time.Time) {
				logger.Log("event", "method.finished", "transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}
