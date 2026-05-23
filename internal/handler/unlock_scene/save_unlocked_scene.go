package unlock_scene

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_unlock_scene"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func saveUnlockedScene(ctx *gin.Context) {
	req := request.SaveUnlockedSceneRequest1{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	for _, sceneType := range req.UnlockSceneTypes.Slice {
		user_unlock_scene.UnlockScene(session, sceneType, enum.UnlockSceneStatusOpened)
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/unlockScene/saveUnlockedScene", saveUnlockedScene)
}
