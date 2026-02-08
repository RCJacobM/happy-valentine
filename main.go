package main

import (
	"valentoins/controllers"
	"valentoins/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.DbConnect()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/create", func(ctx *gin.Context) {
		ctx.HTML(200, "create.html", nil)
	})
	r.POST("/create", controllers.CreateValentine)

	r.GET("/card", controllers.GetValentineCard)
	r.POST("/card", controllers.Result)

	r.Run("0.0.0.0:5001") // 127.0.0.1
}
