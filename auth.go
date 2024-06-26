package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/auth", func(c *gin.Context) {
		streamKey := c.PostForm("name")

		if isValidStreamKey(streamKey) {
			c.String(http.StatusOK, "OK")
		} else {
			c.String(http.StatusForbidden, "Invalid stream key")
		}
	})

	router.Run(":8081")
}

func isValidStreamKey(key string) bool {
	validStreamKeys := map[string]bool{
		"validstreamkey1": true,
		"validstreamkey2": true,
	}
	return validStreamKeys[key]
}
