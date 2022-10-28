package main

import (
	"github.com/Azamjon99/gin/controller"
	"github.com/Azamjon99/gin/middlewares"
	"github.com/Azamjon99/gin/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOut() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(
		f,
		os.Stdout,
	)
}
func main() {
	setupLogOut()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
