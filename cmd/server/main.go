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
		cli.StringFlag{
			Name:   "port, p",
			Value:  "8080",
			Usage:  "TCP Port to listen to",
			EnvVar: "PORT",
		},
		cli.Float64Flag{
			Name:   "tax-value, t",
			Value:  0.20,
			Usage:  "Tax value used by the service",
			EnvVar: "TAX_VALUE",
		},
	}

	app.Action = func(c *cli.Context) error {
		logger := log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
		level.Info(logger).Log("event", "server.loading", "port", c.String("port"))

		taxService := services.NewTax(c.Float64("tax-value"))
		endpoints := endpoints.NewTax(log.With(logger, "service", "tax"), taxService)
		grpcServer := transports.NewTaxGRPCServer(endpoints)
		addr := fmt.Sprintf(":%s", c.String("port"))

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
		return baseServer.Serve(grpcListener)
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}