package config_test

import (
	"os"

	configpkg "github.com/happilymarrieddad/ws-api/internal/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config Tester", func() {

	var config *configpkg.Config

	BeforeEach(func() {
		var err error

		config, err = configpkg.GetConfig()
		Expect(err).NotTo(HaveOccurred())
		Expect(config).NotTo(BeNil())
	})

	It("should return a valid config", func() {
		Expect(config.OpenWeatherAPIKey).NotTo(HaveLen(0))
	})

	It("should err when env is not present", func() {
		os.Setenv("OPEN_WEATHER_API_KEY", "")

		c, e := configpkg.GetConfig()
		Expect(e).To(HaveOccurred())
		Expect(e.Error()).To(Equal("missing env var OPEN_WEATHER_API_KEY"))
		Expect(c).To(BeNil())
	})

})
