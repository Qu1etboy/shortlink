package controllers

import (
	"shortlinks/internal/app/controllers/shortlinks"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.GET("/g/:slug", shortlinks.GET)
	r.POST("/g", shortlinks.POST)
}
