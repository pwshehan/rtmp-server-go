package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/auth", func(c *gin.Context) {
		streamKey := c.Query("name")

		validStreamKeys := map[string]bool{
			"validstreamkey1": true,
			"validstreamkey2": true,
		}

		if validStreamKeys[streamKey] {
			c.String(http.StatusOK, "OK")
		} else {
			c.String(http.StatusForbidden, "Invalid stream key")
		}
	})

	router.Run(":8081")
}
