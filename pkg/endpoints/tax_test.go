package endpoints_test

import (
	"errors"

	"github.com/go-kit/kit/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tscolari/mservice/pkg/endpoints"
	"github.com/tscolari/mservice/pkg/services/servicesfakes"
	"golang.org/x/net/context"
)

var _ = Describe("Tax", func() {
	var (
		subject endpoints.Tax
		logger  log.Logger
	)

	BeforeEach(func() {
		logger = log.NewLogfmtLogger(GinkgoWriter)
		subject = endpoints.Tax{}
	})

	Describe("Add", func() {
		It("uses the given AddEndpoint to calculate", func() {
			called := false
			subject.AddEndpoint = func(_ context.Context, req interface{}) (interface{}, error) {
				called = true
				Expect(req).To(Equal(endpoints.AddRequest{Value: 100}))
				return endpoints.AddResponse{Value: 200}, nil
			}

			value, err := subject.Add(context.Background(), 100)
			Expect(err).NotTo(HaveOccurred())
			Expect(value).To(Equal(200.0))
			Expect(called).To(BeTrue())
		})
	})

	Describe("Sub", func() {
		It("uses the given SubEndpoint to calculate", func() {
			called := false
			subject.SubEndpoint = func(_ context.Context, req interface{}) (interface{}, error) {
				called = true
				Expect(req).To(Equal(endpoints.SubRequest{Value: 100}))
				return endpoints.SubResponse{Value: 250}, nil
			}

			value, err := subject.Sub(context.Background(), 100)
			Expect(err).NotTo(HaveOccurred())
			Expect(value).To(Equal(250.0))
			Expect(called).To(BeTrue())
		})
	})

	Context("NewTax", func() {
		var service *servicesfakes.FakeTax

		BeforeEach(func() {
			service = new(servicesfakes.FakeTax)
			subject = endpoints.NewTax(logger, service)
		})

		Describe("AddEndpoint", func() {
			BeforeEach(func() {
				service.AddReturns(500, nil)
			})

			It("uses the service provided to process the request", func() {
				resp, err := subject.AddEndpoint(context.Background(), endpoints.AddRequest{Value: 100})
				Expect(err).NotTo(HaveOccurred())

				Expect(service.AddCallCount()).To(Equal(1))
				_, argValue := service.AddArgsForCall(0)
				Expect(argValue).To(Equal(100.0))

				Expect(resp).To(Equal(endpoints.AddResponse{Value: 500.0, Err: nil}))
			})

			Context("when the service returns an error", func() {
				BeforeEach(func() {
					service.AddReturns(500, errors.New("failed"))
				})

				It("adds it to the response", func() {
					resp, err := subject.AddEndpoint(context.Background(), endpoints.AddRequest{Value: 100})
					Expect(err).NotTo(HaveOccurred())

					Expect(service.AddCallCount()).To(Equal(1))
					_, argValue := service.AddArgsForCall(0)
					Expect(argValue).To(Equal(100.0))

					Expect(resp).To(Equal(endpoints.AddResponse{Value: 0, Err: errors.New("failed")}))

				})
			})
		})

		Describe("SubEndpoint", func() {
			BeforeEach(func() {
				service.AddReturns(500, nil)
			})

			It("uses the service provided to process the request", func() {
				resp, err := subject.AddEndpoint(context.Background(), endpoints.AddRequest{Value: 100})
				Expect(err).NotTo(HaveOccurred())

				Expect(service.AddCallCount()).To(Equal(1))
				_, argValue := service.AddArgsForCall(0)
				Expect(argValue).To(Equal(100.0))

				Expect(resp).To(Equal(endpoints.AddResponse{Value: 500.0, Err: nil}))
			})

			Context("when the service returns an error", func() {
				BeforeEach(func() {
					service.AddReturns(500, errors.New("failed"))
				})

				It("adds it to the response", func() {
					resp, err := subject.AddEndpoint(context.Background(), endpoints.AddRequest{Value: 100})
					Expect(err).NotTo(HaveOccurred())

					Expect(service.AddCallCount()).To(Equal(1))
					_, argValue := service.AddArgsForCall(0)
					Expect(argValue).To(Equal(100.0))

					Expect(resp).To(Equal(endpoints.AddResponse{Value: 0, Err: errors.New("failed")}))
				})
			})
		})
	})
})
