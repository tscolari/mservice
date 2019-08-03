package integration_test

import (
	"math/rand"
	"os/exec"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Server", func() {
	var (
		sess *gexec.Session
		port string
	)

	BeforeEach(func() {
		port = strconv.Itoa(8000 + rand.Intn(100))
		cmd := exec.Command(serverBinPath, "--port", port, "--tax-value", "0.25")
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
})
