//go:build embedded

package main

import (
	"log"

	"elichika/internal/config"
	_ "elichika/internal/handler"
	"elichika/internal/server"
	_ "elichika/internal/subsystem"
	"elichika/internal/userdata"
	_ "elichika/internal/webui"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("Entered main init")
	log.Println("Start loading userdata")
	userdata.Init()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	server.Router(r)
	log.Println("server address: ", *config.Conf.ServerAddress)
	log.Println("WebUI address: ", *config.Conf.ServerAddress+"/webui/...")
	go func() {
		r.Run(*config.Conf.ServerAddress)
	}()
	log.Println("Exit main init, server should be live")
	embedded.SendLoadedSignal()
	log.ElichikaReady = true
}

func main() {}
