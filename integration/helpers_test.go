package integration_test

import (
	"net"
	"time"

	. "github.com/onsi/gomega"
	"github.com/tscolari/mservice/pkg/services"
	"github.com/tscolari/mservice/pkg/transports"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func mkAddRequest(port string, value float64) (float64, error) {
	conn := mkConn(":" + port)
	defer conn.Close()

	var service services.Tax
	service = transports.NewTaxGRPCClient(conn)
	return service.Add(context.Background(), value)
}

func mkSubRequest(port string, value float64) (float64, error) {
	conn := mkConn(":" + port)
	defer conn.Close()

	var service services.Tax
	service = transports.NewTaxGRPCClient(conn)
	return service.Sub(context.Background(), value)
}

func mkConn(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	Expect(err).NotTo(HaveOccurred())

	return conn
}

func waitForServer(port string) {
	Eventually(func() error {
		_, err := net.DialTimeout("tcp", ":"+port, 5*time.Second)
		return err
	}).ShouldNot(HaveOccurred())
}
