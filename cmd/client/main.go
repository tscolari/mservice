package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tscolari/mservice/pkg/services"
	"github.com/tscolari/mservice/pkg/transports"
	"github.com/urfave/cli"

	"google.golang.org/grpc"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "addr, a",
			Value:  ":8080",
			Usage:  "Address to connect to",
			EnvVar: "ADDR",
		},
		cli.StringFlag{
			Name:  "method, m",
			Value: "",
			Usage: "Method to call on the server",
		},
		cli.Float64Flag{
			Name:  "value",
			Usage: "Value to pass to the method",
		},
	}

	app.Action = func(c *cli.Context) error {
		logger := log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger.Log("event", "starting")
		defer logger.Log("event", "exiting")

		conn, err := grpc.Dial(c.String("addr"), grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			level.Error(logger).Log("err", err)
			return errors.Wrap(err, "dialing connection")
		}
		defer conn.Close()

		var service services.Tax
		service = transports.NewTaxGRPCClient(conn)

		ctx := context.Background()
		var v float64

		switch c.String("method") {
		case "add":
			v, err = service.Add(ctx, c.Float64("value"))
		case "sub":
			v, err = service.Sub(ctx, c.Float64("value"))
		default:
			return errors.Errorf("invalid method: %s", c.String("method"))
		}

		if err != nil {
			level.Error(logger).Log("err", err)
			return errors.Wrap(err, "the server returned an error")
		}

		level.Info(logger).Log("resp", v)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
