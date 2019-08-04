package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
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
		logger := log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger.Log("event", "starting")
		defer logger.Log("event", "exiting")

		if !c.IsSet("method") || !c.IsSet("value") {
			return errors.New("No --method or --value argument present")
		}

		conn, err := grpc.Dial(c.String("addr"), grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
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
			return errors.Wrap(err, "the server returned an error")
		}

		fmt.Printf("Result: %.2f\n", v)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}
