package main

import (
	"suber/fuck"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/:file", fuck.Get)

	r.POST("/:auth/upload", fuck.Post)

	r.Run(":8080")
}
