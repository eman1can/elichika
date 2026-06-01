package user

import (
	"net/http"

	"elichika/internal/enum"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_card"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUILimitBreakRequest struct {
	CardMasterIds []int32 `form:"card_master_ids" json:"card_master_ids"`
	Amount        int32   `form:"amount" json:"amount"`
}

func limitBreak(ctx *gin.Context) {
	req := WebUILimitBreakRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)

	for _, cardMasterId := range req.CardMasterIds {
		masterCard := session.Gamedata.Card[cardMasterId]
		card := user_card.GetUserCard(session, cardMasterId)
		if card.Grade > enum.CardNotOwned {
			for ix := card.Grade; ix < min(enum.CardMaxGrade, card.Grade+req.Amount); ix++ {
				card.Grade++
				user_member.IncreaseMemberLoveLevelLimit(session, masterCard.Member.Id, masterCard.Rarity.PlusLevel)
				user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountSchoolIdolGrade, nil, nil, user_mission.AddProgressHandler, int32(1))
			}
			user_card.UpdateUserCard(session, card)
		}
	}

	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/limit_break", limitBreak)
}
