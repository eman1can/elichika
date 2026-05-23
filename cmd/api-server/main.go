//go:build !embedded

package main

import (
	"log"

	"elichika/internal/config"
	_ "elichika/internal/handler"
	"elichika/internal/server"
	_ "elichika/internal/subsystem"
	"elichika/internal/subsystem/user_training_tree"
	"elichika/internal/userdata"
	_ "elichika/internal/webui"

	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

// with some cli, we keep the server open
// return true to keep open
func checkCli() bool {
	if os.Args[1] == "rebuild_assets" {
		if len(os.Args) > 2 && os.Args[2] == "keep_alive" {
			return true
		}
	}
	if os.Args[1] == "fix_training_trees" {
		user_training_tree.FixUsersTrainingTrees()
	}
	log.Println("CLI is reserved for special behaviour, the server will now exit, start it again without any argument!")
	return false
}

func main() {
	if len(os.Args) > 1 {
		if !checkCli() {
			return
		}
	}
	userdata.Init()
	runtime.GC()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	server.Router(r)
	log.Println("Server Address: ", *config.Conf.ServerAddress)

	go func() {
		r.Run(*config.Conf.ServerAddress)
	}()
	server.ReceiveFinalSignal()
}
