package main

import (
	"io"
	"os"

	"github.com/QuvonchbekBobojonov/golang-gin-pic/controller"
	"github.com/QuvonchbekBobojonov/golang-gin-pic/middlewares"
	"github.com/QuvonchbekBobojonov/golang-gin-pic/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService   service.VideoService     = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutPut(){
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}


func main() {

	setupLogOutPut()

	server := gin.New()

	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
		gindump.Dump(),
	)

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
