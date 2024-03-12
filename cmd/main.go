package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shortlinks/internal/pkg/firebase"
	"shortlinks/internal/pkg/utils"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

func main() {
	_, client, ctx := firebase.Firebase()

	defer client.Close()

	r := gin.Default()

	r.GET("/g/:slug", func(c *gin.Context) {
		short := c.Param("slug")

		dsnap, err := client.Collection("shortlinks").Doc(short).Get(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		shortlink := dsnap.Data()
		if shortlink == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.Redirect(http.StatusMovedPermanently, shortlink["original"].(string))
	})

	r.POST("/g", func(c *gin.Context) {
		var request ShortenRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		randomString, err := utils.GenerateRandomString(4)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		_, err = client.Collection("shortlinks").Doc(randomString).Set(ctx, map[string]interface{}{
			"short":    randomString,
			"original": request.URL,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"short":    randomString,
			"original": request.URL,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
