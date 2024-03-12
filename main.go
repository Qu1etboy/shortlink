package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

var shortURLs = make(map[string]string)

func generateRandomString(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buffer)[:length], nil
}

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/g/:slug", func(c *gin.Context) {
		short := c.Param("slug")

		if url, ok := shortURLs[short]; ok {
			c.Redirect(http.StatusMovedPermanently, url)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		}
	})

	r.POST("/g", func(c *gin.Context) {
		var request ShortenRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		randomString, err := generateRandomString(4)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		shortURLs[randomString] = request.URL

		c.JSON(http.StatusOK, gin.H{
			"short":    randomString,
			"original": request.URL,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
