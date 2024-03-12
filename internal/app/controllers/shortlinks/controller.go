package shortlinks

import (
	"net/http"
	"shortlinks/internal/pkg/firebase"
	"shortlinks/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

func GET(c *gin.Context) {
	_, client, ctx := firebase.Firebase()

	defer client.Close()

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
}

func POST(c *gin.Context) {
	_, client, ctx := firebase.Firebase()

	defer client.Close()

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
}
