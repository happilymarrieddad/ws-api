package middleware_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/happilymarrieddad/ws-api/internal/api/middleware"
	"github.com/happilymarrieddad/ws-api/internal/wsclient"
	"github.com/happilymarrieddad/ws-api/internal/wsclient/mocks"

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
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should be a function", func() {
		Expect(middleware.HTTPWSClientInjector).To(BeAssignableToTypeOf(func(wsclient.WSClient) gin.HandlerFunc { return nil }))
	})

	It("should add ws client to context then call the next handler", func() {
		router.GET("/", func(c *gin.Context) {
			iface, exists := c.Get(middleware.ContextKeyWSClient)
			Expect(exists).To(BeTrue())
			Expect(iface).To(BeEquivalentTo(mockWsClient))
		})
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(rec, req)
	})
})
