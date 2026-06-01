package user

import (
	"net/http"

	"elichika/internal/enum"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_card"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/subsystem/user_training_tree"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

var MaxCardLevelByMemberId = make(map[int32]map[int32]int32)

func isCardRarityEffect(bonusType int32) bool {
	switch bonusType {
	case enum.MemberLovePanelEffectTypeRLevel:
		fallthrough
	case enum.MemberLovePanelEffectTypeSrLevel:
		fallthrough
	case enum.MemberLovePanelEffectTypeUrLevel:
		return true
	default:
	}
	return false
}

func toCardRarity(bonusType int32) int32 {
	if bonusType == enum.MemberLovePanelEffectTypeRLevel {
		return enum.CardRarityTypeRare
	} else if bonusType == enum.MemberLovePanelEffectTypeSrLevel {
		return enum.CardRarityTypeSRare
	}

	return enum.CardRarityTypeURare
}

func getMaxCardLevel(session *userdata.Session, memberId int32) map[int32]int32 {
	var result = make(map[int32]int32)

	for _, rarity := range enum.AllCardRarities {
		result[rarity] = session.Gamedata.CardRarity[rarity].MaxLevel
	}

	masterPanel := session.Gamedata.MemberFirstLovePanel[memberId]
	for masterPanel != nil {
		panel := user_member.GetMemberLovePanel(session, memberId, masterPanel.Id)

		for ix, cellId := range masterPanel.CellIds {
			if panel.Unlocked(ix) {
				cellContent := session.Gamedata.MemberLovePanelCell[cellId]
				if !isCardRarityEffect(cellContent.BonusType) {
					continue
				}

				result[toCardRarity(cellContent.BonusType)] += cellContent.BonusValue

				if panel.AllUnlocked() {
					for _, bonus := range masterPanel.Bonuses {
						result[toCardRarity(bonus.BonusType)] += cellContent.BonusValue
					}
				}
			}
		}

		if !panel.AllUnlocked() {
			break
		}
		masterPanel = masterPanel.NextPanel
	}

	return result
}

type WebUIPracticeCardRequest struct {
	CardMasterIds []int32 `form:"card_master_ids" json:"card_master_ids"`
}

func practiceCard(ctx *gin.Context) {
	req := WebUIPracticeCardRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)

	for _, cardMasterId := range req.CardMasterIds {
		masterCard := session.Gamedata.Card[cardMasterId]
		memberId := masterCard.Member.Id

		maxLevel, ok := MaxCardLevelByMemberId[memberId]
		if !ok {
			maxLevel = getMaxCardLevel(session, memberId)
			MaxCardLevelByMemberId[memberId] = maxLevel
		}

		card := user_card.GetUserCard(session, cardMasterId)
		card.Level = maxLevel[masterCard.CardRarityType]

		if card.IsAllTrainingActivated {
			continue
		}

		trainingTree := user_training_tree.GetUserTrainingTree(session, cardMasterId)
		hasCellId := map[int32]bool{}
		for _, cell := range trainingTree.Slice {
			hasCellId[cell.CellId] = true
		}

		var addedCells []int32
		for ix, cell := range masterCard.TrainingTree.TrainingTreeMapping.TrainingTreeCellContents {
			if ix == 0 || cell.RequiredGrade > card.Grade || hasCellId[cell.CellId] {
				continue
			}
			addedCells = append(addedCells, cell.CellId)
		}

		if len(addedCells) > 0 {
			user_training_tree.ActivateTrainingTreeCells(session, card.CardMasterId, addedCells)
		}

		card.IsAllTrainingActivated = true

		user_card.UpdateUserCard(session, card)
	}

	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/practice_card", practiceCard)
}
