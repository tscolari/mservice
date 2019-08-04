package integration_test

import (
	"math/rand"
	"os/exec"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Client (and Server)", func() {
	var (
		sess *gexec.Session
		port string
	)

	BeforeEach(func() {
		port = strconv.Itoa(8000 + rand.Intn(100))
		cmd := exec.Command(serverBinPath,
			"--port", port,
			"--tax-value", "0.25",
			"--tls-cert", "./assets/certs/server.crt",
			"--tls-key", "./assets/certs/server.key",
			"--ca-cert", "./assets/certs/ca.crt")
		var err error
		sess, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		waitForServer(port)
	})

	AfterEach(func() {
		sess.Kill()
		Eventually(sess).Should(gexec.Exit())
	})

	Describe("Add Method", func() {
		It("returns the correct values for the request", func() {
			cmd := exec.Command(clientBinPath,
				"--addr", "127.0.0.1:"+port,
				"--method", "add",
				"--value", "1000",
				"--tls-cert", "./assets/certs/client.crt",
				"--tls-key", "./assets/certs/client.key",
				"--ca-cert", "./assets/certs/ca.crt")

			buffer := gbytes.NewBuffer()
			cli, err := gexec.Start(cmd, buffer, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(cli, 1*time.Second).Should(gexec.Exit(0))
			Eventually(buffer).Should(gbytes.Say("Result: 1250.00"))
		})
	})

	Describe("Sub Method", func() {
		It("returns the correct values for the request", func() {
			cmd := exec.Command(clientBinPath,
				"--addr", "127.0.0.1:"+port,
				"--method", "sub",
				"--value", "1250",
				"--tls-cert", "./assets/certs/client.crt",
				"--tls-key", "./assets/certs/client.key",
				"--ca-cert", "./assets/certs/ca.crt")

			buffer := gbytes.NewBuffer()
			cli, err := gexec.Start(cmd, buffer, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(cli, 1*time.Second).Should(gexec.Exit(0))
			Eventually(buffer).Should(gbytes.Say("Result: 1000.00"))
		})
	})

	Context("when using invalid certificates", func() {
		Describe("Add Method", func() {
			It("returns the correct values for the request", func() {
				cmd := exec.Command(clientBinPath,
					"--addr", "127.0.0.1:"+port,
					"--method", "add",
					"--value", "1000",
					"--tls-cert", "./assets/certs_v2/client.crt",
					"--tls-key", "./assets/certs_v2/client.key",
					"--ca-cert", "./assets/certs/ca.crt")

				buffer := gbytes.NewBuffer()
				cli, err := gexec.Start(cmd, GinkgoWriter, buffer)
				Expect(err).NotTo(HaveOccurred())

				Eventually(cli, 1*time.Second).Should(gexec.Exit(1))
				Eventually(buffer).Should(gbytes.Say("handshake failed: remote error: tls: bad certificate"))
			})
		})

		Describe("Sub Method", func() {
			It("returns the correct values for the request", func() {
				cmd := exec.Command(clientBinPath,
					"--addr", "127.0.0.1:"+port,
					"--method", "sub",
					"--value", "1250",
					"--tls-cert", "./assets/certs_v2/client.crt",
					"--tls-key", "./assets/certs_v2/client.key",
					"--ca-cert", "./assets/certs/ca.crt")

				buffer := gbytes.NewBuffer()
				cli, err := gexec.Start(cmd, GinkgoWriter, buffer)
				Expect(err).NotTo(HaveOccurred())

				Eventually(cli, 1*time.Second).Should(gexec.Exit(1))
				Eventually(buffer).Should(gbytes.Say("handshake failed: remote error: tls: bad certificate"))
			})
		})
	})
})
