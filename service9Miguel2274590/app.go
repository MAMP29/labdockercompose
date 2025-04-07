package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "MAMP sends greetings from Go!"})
	})

	r.Run(":5009")
}
