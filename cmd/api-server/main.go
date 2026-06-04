//go:build !embedded

package main

import (
	"log"
	"runtime"

	"elichika/internal/clientdb"
	"elichika/internal/config"
	"elichika/internal/handler"
	"elichika/internal/locale"
	"elichika/internal/server"
	"elichika/internal/serverdata"
	"elichika/internal/serverstate"
	"elichika/internal/subsystem"
	"elichika/internal/userdata"
	"elichika/internal/webui"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	clientdb.Init()
	config.Init()
	serverstate.Init()
	serverdata.Init()
	userdata.Init()
	locale.Init()
	handler.Register()
	subsystem.Register()
	webui.Register()

	runtime.GC()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Session-Key"},
		AllowCredentials: true,
	}))
	server.Router(r)
	log.Println("Server Address: ", *config.Conf.ServerAddress)

	go func() {
		err := r.Run(*config.Conf.ServerAddress)
		if err != nil {
			log.Fatal(err)
		}
	}()
	server.ReceiveFinalSignal()
}
