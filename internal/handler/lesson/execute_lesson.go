package handler

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_lesson"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func executeLesson(ctx *gin.Context) {
	req := request.ExecuteLessonRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_lesson.ExecuteLesson(session, req)

	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/lesson/executeLesson", executeLesson)
}
