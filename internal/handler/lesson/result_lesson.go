package handler

import (
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_lesson"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func resultLesson(ctx *gin.Context) {
	// there is no request body

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_lesson.ResultLesson(session)

	common.JsonResponse(ctx, resp)
}

func init() {
	server.AddHandler("/", "POST", "/lesson/resultLesson", resultLesson)
}
