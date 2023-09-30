package wsclient_test

import (
	"context"

	configpkg "github.com/happilymarrieddad/ws-api/internal/config"
	"github.com/happilymarrieddad/ws-api/internal/wsclient"
	"github.com/happilymarrieddad/ws-api/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("WSClient Tester", func() {

	var client wsclient.WSClient
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()

		config, err := configpkg.GetConfig()
		Expect(err).NotTo(HaveOccurred())
		Expect(config).NotTo(BeNil())

		client, err = wsclient.NewWSClient(config, nil) // just use the default http client
		Expect(err).NotTo(HaveOccurred())
		Expect(config).NotTo(BeNil())
	})

	Context("GetWeatherAtLongLat", func() {
		It("should return an err when an invalid lat long is passed in", func() {
			res, err := client.GetWeatherAtLongLat(ctx, -111.773613, 43.629398, utils.Ref(wsclient.Fahrenheit))
			Expect(err).NotTo(BeNil())
			Expect(res).To(BeNil())
		})

		It("should return a valid weather at a specific lat long", func() {
			res, err := client.GetWeatherAtLongLat(ctx, 43.629398, -111.773613, utils.Ref(wsclient.Fahrenheit))
			Expect(err).To(BeNil())

			Expect(res.City).To(Equal("Rigby"))
		})
	})

	Context("GetWeatherConditonFromTempType", func() {
		It("should return an error with an invalid temp type", func() {
			_, err := client.GetWeatherConditonFromTempType("GARBAGE", 83.56)
			Expect(err).To(HaveOccurred())
		})

		It("should return the correct weather type for Fahrenheit", func() {
			condition, err := client.GetWeatherConditonFromTempType(wsclient.Fahrenheit, 55.5)
			Expect(err).NotTo(HaveOccurred())
			Expect(condition).To(Equal(wsclient.Moderate))

			condition, err = client.GetWeatherConditonFromTempType(wsclient.Fahrenheit, 44.5)
			Expect(err).NotTo(HaveOccurred())
			Expect(condition).To(Equal(wsclient.Cold))

			condition, err = client.GetWeatherConditonFromTempType(wsclient.Fahrenheit, 84.5)
			Expect(err).NotTo(HaveOccurred())
			Expect(condition).To(Equal(wsclient.Hot))
		})

		It("should return the correct weather type for Celcius", func() {
			condition, err := client.GetWeatherConditonFromTempType(wsclient.Celcius, 10)
			Expect(err).NotTo(HaveOccurred())
			Expect(condition).To(Equal(wsclient.Moderate))

			condition, err = client.GetWeatherConditonFromTempType(wsclient.Celcius, 7)
			Expect(err).NotTo(HaveOccurred())
			Expect(condition).To(Equal(wsclient.Cold))

			condition, err = client.GetWeatherConditonFromTempType(wsclient.Celcius, 22)
			Expect(err).NotTo(HaveOccurred())
			Expect(condition).To(Equal(wsclient.Hot))
		})

		It("should return the correct weather type for Kelvin", func() {
			condition, err := client.GetWeatherConditonFromTempType(wsclient.Kelvin, 286)
			Expect(err).NotTo(HaveOccurred())
			Expect(condition).To(Equal(wsclient.Moderate))

			condition, err = client.GetWeatherConditonFromTempType(wsclient.Kelvin, 280)
			Expect(err).NotTo(HaveOccurred())
			Expect(condition).To(Equal(wsclient.Cold))

			condition, err = client.GetWeatherConditonFromTempType(wsclient.Kelvin, 295)
			Expect(err).NotTo(HaveOccurred())
			Expect(condition).To(Equal(wsclient.Hot))
		})
	})

})
