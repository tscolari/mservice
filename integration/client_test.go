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

var _ = Describe("Client", func() {
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

	Describe("Add Method", func() {
		It("returns the correct values for the request", func() {
			cmd := exec.Command(clientBinPath, "--addr", ":"+port, "--method", "add", "--value", "1000")
			buffer := gbytes.NewBuffer()
			cli, err := gexec.Start(cmd, GinkgoWriter, buffer)
			Expect(err).NotTo(HaveOccurred())

			Eventually(cli, 1*time.Second).Should(gexec.Exit(0))
			Eventually(buffer).Should(gbytes.Say("resp=1250"))
		})
	})

	Describe("Sub Method", func() {
		It("returns the correct values for the request", func() {
			cmd := exec.Command(clientBinPath, "--addr", ":"+port, "--method", "sub", "--value", "1250")
			buffer := gbytes.NewBuffer()
			cli, err := gexec.Start(cmd, GinkgoWriter, buffer)
			Expect(err).NotTo(HaveOccurred())

			Eventually(cli, 1*time.Second).Should(gexec.Exit(0))
			Eventually(buffer).Should(gbytes.Say("resp=1000"))
		})
	})
})
