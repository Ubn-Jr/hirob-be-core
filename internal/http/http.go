package http

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"mqtt"  //from mqtt.go
)

func setupHTTPServer() *gin.Engine {
	r := gin.Default()

	r.POST("/api/movement", func(c *gin.Context) {
		movementData, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}

		mqtt.Publish(movementData)
		c.String(http.StatusOK, "Movement request sent to MQTT")
	})

	return r
}
