package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/happilymarrieddad/ws-api/internal/wsclient"
)

const (
	ContextKeyWSClient = "mw:WSClient"
)

func HTTPWSClientInjector(wsc wsclient.WSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ContextKeyWSClient, wsc)
		c.Next()
	}
}

func HTTPRetrieveWSClient(c *gin.Context) wsclient.WSClient {
	if iface, exists := c.Get(ContextKeyWSClient); exists {
		if wsc, ok := iface.(wsclient.WSClient); ok {
			return wsc
		}
	}
	return nil
}
