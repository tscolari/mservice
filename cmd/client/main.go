package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tscolari/mservice/pkg/services"
	"github.com/tscolari/mservice/pkg/transports"
	"github.com/urfave/cli"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
		cli.StringFlag{
			Name:   "tls-cert",
			Usage:  "Path to the TLS certificate",
			EnvVar: "TLS_CERT_PATH",
		},
		cli.StringFlag{
			Name:   "tls-key",
			Usage:  "Path to the TLS key",
			EnvVar: "TLS_KEY_PATH",
		},
		cli.StringFlag{
			Name:   "ca-cert",
			Usage:  "Path to the CA cert",
			EnvVar: "CA_CERT_PATH",
		},
		cli.BoolFlag{
			Name:  "insecure",
			Usage: "Skip TLS authentication",
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

		var authOption grpc.DialOption

		if c.Bool("insecure") {
			authOption = grpc.WithInsecure()
		} else {
			tlsCreds, err := loadCerts(c.String("addr"), c.String("tls-cert"), c.String("tls-key"), c.String("ca-cert"))
			if err != nil {
				level.Error(logger).Log("event", "tls.failed", "err", err)
				return errors.Wrap(err, "loading TLS credentials")
			}

			authOption = grpc.WithTransportCredentials(tlsCreds)
		}

		conn, err := grpc.Dial(c.String("addr"), authOption, grpc.WithTimeout(time.Second))
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

func loadCerts(addr, certPath, keyPath, caPath string) (credentials.TransportCredentials, error) {
	crt, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, errors.Wrap(err, "could not load key pair from")
	}

	rawCaCrt, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, errors.Wrap(err, "could not load CA certificate from file")
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(rawCaCrt); !ok {
		return nil, errors.Wrap(err, "could not append CA certificate to the pool")
	}

	addr = strings.Split(addr, ":")[0]
	tlsCreds := credentials.NewTLS(&tls.Config{
		ServerName:   addr,
		Certificates: []tls.Certificate{crt},
		RootCAs:      certPool,
	})

	return tlsCreds, nil
}
