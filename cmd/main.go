package main

import (
	"github.com/gin-gonic/gin"

	"shortlinks/internal/app/controllers"
)

func main() {

	r := gin.Default()

	controllers.Init(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
