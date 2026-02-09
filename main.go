package main

import (
	"net/http"
	"os"
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
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, os.Getenv("base")+"create")
	})

	r.GET("/create", func(ctx *gin.Context) {
		ctx.HTML(200, "create.html", gin.H{"base": os.Getenv("base")})
	})
	r.POST("/create", controllers.CreateValentine)

	r.GET("/card", controllers.GetValentineCard)
	r.POST("/card", controllers.Result)

	r.Run("127.0.0.1:5001") // 127.0.0.1
}
