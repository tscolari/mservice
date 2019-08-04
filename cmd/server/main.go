package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/pkg/errors"
	"github.com/tscolari/mservice/pkg/endpoints"
	"github.com/tscolari/mservice/pkg/pb"
	"github.com/tscolari/mservice/pkg/services"
	"github.com/tscolari/mservice/pkg/transports"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port, p",
			Value:  8080,
			Usage:  "TCP Port to listen to",
			EnvVar: "PORT",
		},
		cli.Float64Flag{
			Name:   "tax-value, t",
			Value:  0.20,
			Usage:  "Tax value used by the service",
			EnvVar: "TAX_VALUE",
		},
		cli.IntFlag{
			Name:   "health-check-port",
			Value:  8081,
			Usage:  "Health check port",
			EnvVar: "HEALTH_CHECK_PORT",
		},
	}

	app.Action = func(c *cli.Context) error {
		logger := log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)

		if !c.IsSet("tax-value") {
			err := errors.New("No --tax-value given")
			level.Error(logger).Log("event", "server.loading", "err", err)
			return err
		}

		level.Info(logger).Log("event", "server.loading", "port", c.Int("port"), "tax_value", c.Float64("tax-value"))
		taxService := services.NewTax(c.Float64("tax-value"))
		endpoints := endpoints.NewTax(log.With(logger, "service", "tax"), taxService)
		grpcServer := transports.NewTaxGRPCServer(endpoints)
		addr := fmt.Sprintf(":%d", c.Int("port"))

		grpcListener, err := net.Listen("tcp", addr)
		if err != nil {
			level.Error(logger).Log("event", "listen", "err", err)
			return errors.Wrap(err, "grpc.listening")
		}
		defer grpcListener.Close()

		baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
		pb.RegisterTaxServer(baseServer, grpcServer)
		level.Info(logger).Log("event", "server.started")
		defer level.Info(logger).Log("event", "server.finished")
		go startHealthCheck(logger, c.Int("health-check-port"))
		return baseServer.Serve(grpcListener)
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
