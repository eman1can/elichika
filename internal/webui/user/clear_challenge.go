package user

import (
	"elichika/internal/server"
	"elichika/internal/subsystem/user_beginner_challenge"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func clearBeginnerChallenges(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)
	challengeCells := user_beginner_challenge.GetBeginnerChallengeCells(session)

	for _, cell := range challengeCells {
		if cell.IsRewardReceived {
			continue
		}
		cell.IsRewardReceived = true
		user_beginner_challenge.UpdateChallengeCell(session, *cell)
	}
	session.Finalize()
}

func init() {
	server.AddHandler("/webui/user", "POST", "/clear_beginner_challenge", clearBeginnerChallenges)
}
