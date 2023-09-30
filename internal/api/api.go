package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/happilymarrieddad/ws-api/internal/api/middleware"
	"github.com/happilymarrieddad/ws-api/internal/config"
	"github.com/happilymarrieddad/ws-api/internal/wsclient"
)

type LatLong struct {
	Latitude  float64            `form:"lat"`
	Longitude float64            `form:"long"`
	TempType  *wsclient.TempType `form:"tempType"`
}

func GetWeather(c *gin.Context) {
	var latLong LatLong
	if err := c.ShouldBind(&latLong); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := wsclient.ValidateTempType(latLong.TempType); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := middleware.HTTPRetrieveWSClient(c).GetWeatherDataAtLongLat(c, latLong.Latitude, latLong.Longitude, latLong.TempType)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func Start(cfg *config.Config, client wsclient.WSClient, port int) {
	r := gin.Default()

	r.Use(middleware.HTTPWSClientInjector(client))

	r.GET("/weather", GetWeather)

	fmt.Printf("Server running on port %d\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
