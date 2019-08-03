package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tscolari/mservice/pkg/services"
	"golang.org/x/net/context"
)

var _ = Describe("Tax", func() {
	Describe("Add", func() {
		It("adds the tax value to the base value", func() {
			taxService := services.NewTax(0.20)
			value, err := taxService.Add(context.Background(), 100)
			Expect(err).NotTo(HaveOccurred())
			Expect(value).To(Equal(120.0))
		})

		It("returns an error if the value is negative", func() {
			taxService := services.NewTax(0.20)
			_, err := taxService.Add(context.Background(), -100)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("negative input"))
		})
	})

	Describe("Sub", func() {
		It("subtracts the tax value from the total value", func() {
			taxService := services.NewTax(0.20)
			value, err := taxService.Sub(context.Background(), 120)
			Expect(err).NotTo(HaveOccurred())
			Expect(value).To(Equal(100.0))
		})

		It("returns an error if the value is negative", func() {
			taxService := services.NewTax(0.20)
			_, err := taxService.Sub(context.Background(), -100)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("negative input"))
		})
	})
})
