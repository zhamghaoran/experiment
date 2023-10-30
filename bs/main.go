package main

import (
	"experiment/bs/controller"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	hertz := server.Default()
	hertz.POST("/add", controller.Add)
	hertz.GET("/show", controller.Show)
	hertz.POST("/delete", controller.Delete)
	hertz.POST("/modify", controller.Modify)
	hertz.Spin()
}
