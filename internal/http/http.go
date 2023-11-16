package http

import (
	"net/http"

	mqtt "github.com/Ubn-Jr/hirob-be-core/internal/mqtt"

	"github.com/gin-gonic/gin"
)

func setupHTTPServer() *gin.Engine {
	r := gin.Default()

	r.POST("/api/movement", func(c *gin.Context) {
		movementData, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}

		mqtt.Publish(string(movementData))
		c.String(http.StatusOK, "Movement request szsdfdsfent to MQTT")
	})

	return r
}
