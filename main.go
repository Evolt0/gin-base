package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Parker-Yang/gin-base/global"
	"github.com/Parker-Yang/gin-base/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println(global.Conf.Client)
	gin.SetMode(global.Conf.App.Mode)
	engine := gin.Default()
	routes.SetRoutes(engine)
	go func() {
		err := engine.Run(global.Conf.App.Port)
		logrus.Fatalf("failed to run project! %v", err)
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	select {
	case <-c:
		logrus.Println("exit...")
	}
}
