package http

import (
	"net/http"

	mqtt "github.com/Ubn-Jr/hirob-be-core/internal/mqtt"

	"github.com/gin-gonic/gin"
)
//Defines a function named SetupHTTPServer that returns a pointer to a gin.Engine.
//The gin.Engine is a part of the Gin web framework, which is used for building web applications in Go.
func SetupHTTPServer() *gin.Engine {
//new instance of the gin.Engine is created with some default settings.
//gin.Default() sets up a Gin router with the logger and recovery middleware.
	r := gin.Default()
//a route for handling HTTP POST requests to the "/api/movement" endpoint.
//The second argument is an anonymous function that takes a gin.Context parameter, 
//representing the context of the current HTTP request.
	r.POST("/api/movement", func(c *gin.Context) {
		//Retrieves the raw request body data using c.GetRawData(). If an error occurs during this process,
		//it returns a JSON response with a 500 Internal Server Error status and an error message.
		movementData, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
		//The raw data is then converted to a string and published using the mqtt.Publish function. 
		mqtt.Publish(string(movementData))
		//After successfully processing the request and publishing the data,
		//the code sends a response to the client with a 200 OK status and the message.
		c.String(http.StatusOK, "Movement request sent to MQTT")
	})

	r.Run() //This line starts the HTTP server and makes it listen for incoming requests on the default address "0.0.0.0:8080".

	return r
}
