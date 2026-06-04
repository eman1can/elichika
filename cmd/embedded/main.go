//go:build embedded

package main

import (
	"log"

	"elichika/internal/config"
	"elichika/internal/handler"
	"elichika/internal/locale"
	"elichika/internal/server"
	"elichika/internal/serverdata"
	"elichika/internal/serverstate"
	"elichika/internal/subsystem"
	"elichika/internal/userdata"
	"elichika/internal/webui"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("Entered main init")
	config.Init()
	serverstate.Init()
	serverdata.Init()
	log.Println("Start loading userdata")
	userdata.Init()
	locale.Init()
	handler.Register()
	subsystem.Register()
	webui.Register()

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
