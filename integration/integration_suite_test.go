package integration_test

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var (
	serverBinPath string
	clientBinPath string
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)

	SynchronizedBeforeSuite(func() []byte {
		sBinPath, err := gexec.Build("github.com/tscolari/mservice/cmd/server")
		Expect(err).NotTo(HaveOccurred())

		cBinPath, err := gexec.Build("github.com/tscolari/mservice/cmd/client")
		Expect(err).NotTo(HaveOccurred())

		return []byte(sBinPath + "|" + cBinPath)
	}, func(binPaths []byte) {
		paths := strings.Split(string(binPaths), "|")

		serverBinPath = string(paths[0])
		clientBinPath = string(paths[1])
	})

	SynchronizedAfterSuite(func() {
	}, func() {
		gexec.CleanupBuildArtifacts()
	})

	RunSpecs(t, "Integration Suite")
}
