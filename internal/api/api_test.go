package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/happilymarrieddad/ws-api/internal/api"
	"github.com/happilymarrieddad/ws-api/internal/api/middleware"
	"github.com/happilymarrieddad/ws-api/internal/wsclient"
	"github.com/happilymarrieddad/ws-api/internal/wsclient/mocks"
	"github.com/happilymarrieddad/ws-api/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTTP Testing", func() {

	var (
		router       *gin.Engine
		ctrl         *gomock.Controller
		mockWsClient *mocks.MockWSClient
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		Expect(ctrl).NotTo(BeNil())

		mockWsClient = mocks.NewMockWSClient(ctrl)
		Expect(mockWsClient).NotTo(BeNil())

		gin.SetMode(gin.TestMode)
		router = gin.Default()
		router.Use(middleware.HTTPWSClientInjector(mockWsClient))
		router.GET("/weather", api.GetWeather)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("/weather", func() {
		var (
			rec  *httptest.ResponseRecorder
			lat  = float64(43.629398)
			long = float64(-111.773613)
		)

		BeforeEach(func() {
			rec = httptest.NewRecorder()
		})

		It("should complain when a bad request is passed in", func() {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/weather?lat=%f&long=%f&tempType=garbage", lat, long), nil)

			router.ServeHTTP(rec, req)

			Expect(rec.Code).To(Equal(http.StatusBadRequest))
		})

		It("should successfully return response", func() {
			mockWsClient.EXPECT().GetWeatherDataAtLongLat(gomock.Any(), lat, long, gomock.Any()).Return(&wsclient.WeatherResponse{
				TempFeelsLike: wsclient.Moderate,
				Temperature:   38.48,
				Conditions: []*types.Weather{
					{
						ID:          1,
						Main:        "clouds",
						Description: "some description",
					},
				},
			}, nil).Times(1)

			req, _ := http.NewRequest("GET", fmt.Sprintf("/weather?lat=%f&long=%f&tempType=imperial", lat, long), nil)

			router.ServeHTTP(rec, req)

			Expect(rec.Code).To(Equal(http.StatusOK))
		})
	})
})
