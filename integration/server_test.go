package integration_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"os/exec"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Server", func() {
	var (
		sess   *gexec.Session
		port   string
		hcPort string
	)

	BeforeEach(func() {
		port = strconv.Itoa(8000 + rand.Intn(100))
		hcPort = strconv.Itoa(9000 + rand.Intn(100))
		cmd := exec.Command(serverBinPath,
			"--port", port,
			"--tax-value", "0.25",
			"--health-check-port", hcPort,
			"--insecure") // Authenticated tests are done with the client
		var err error
		sess, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		waitForServer(port)
	})

	AfterEach(func() {
		sess.Kill()
		Eventually(sess).Should(gexec.Exit())
	})

	Describe("Add Endpoint", func() {
		It("returns the correct values for the request", func() {
			value, err := mkAddRequest(port, 1000)
			Expect(err).NotTo(HaveOccurred())
			Expect(value).To(Equal(1250.0))
		})
	})

	Describe("Sub Endpoint", func() {
		It("returns the correct values for the request", func() {
			value, err := mkSubRequest(port, 1250)
			Expect(err).NotTo(HaveOccurred())
			Expect(value).To(Equal(1000.0))
		})
	})

	Context("health check", func() {
		It("responds to health checks", func() {
			var resp *http.Response

			Eventually(func() error {
				var err error
				resp, err = http.Get(fmt.Sprintf("http://localhost:%s/health_check", hcPort))
				return err
			}).ShouldNot(HaveOccurred())

			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
