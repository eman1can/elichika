package login

import (
	"encoding/json"
	"fmt"

	"elichika/internal/client"
	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/config"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_live"
	"elichika/internal/userdata"
	utils2 "elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	req := request.LoginRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils2.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	ctx.Set("sign_key", session.AuthorizationKey())
	if session.AuthenticationData.AuthorizationCount+1 != req.AuthCount { // wrong authcount
		common.AlternativeJsonResponse(ctx, response.InvalidAuthCountResponse{
			AuthorizationCount: session.AuthenticationData.AuthorizationCount,
		})
		return
	} else {
		session.AuthenticationData.AuthorizationCount++
	}

	// log.Println("User logins: ", userId)

	resp := session.Login()
	resp.SessionKey = session.EncodedSessionKey(req.Mask)
	{
		exist, _, startLiveRequest := user_live.LoadUserLive(session)
		if exist {
			liveDifficulty := session.Gamedata.LiveDifficulty[startLiveRequest.LiveDifficultyId]
			if (liveDifficulty.UnlockPattern != enum.LiveUnlockPatternCoopOnly) &&
				(liveDifficulty.UnlockPattern != enum.LiveUnlockPatternTowerOnly) {
				resp.LiveResume = generic.NewNullable(client.LiveResume{
					LiveDifficultyId: startLiveRequest.LiveDifficultyId,
					DeckId:           startLiveRequest.DeckId,
					ConsumedLp:       liveDifficulty.ConsumedLP, // this thing is only to show how much lp is spent
				})
			} else { // just cancel this as it's not a relevant live (event and such)
				user_live.ClearUserLive(session)
			}
		}
	}
	resp.CheckMaintenance = true
	common.JsonResponse(ctx, &resp)

	{
		backupText, err := json.Marshal(resp)
		utils2.CheckErr(err)
		utils2.WriteAllText(fmt.Sprint(config.UserDataBackupPath, "login_", session.UserId, ".json"), string(backupText))
	}
}

func init() {
	server.AddHandler("/", "POST", "/login/login", login)
}
