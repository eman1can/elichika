package serverdata

import (
	"log"
	"os"
	"path/filepath"

	"elichika/internal/client"
	"elichika/internal/config"
	"elichika/internal/generic"
	"elichika/internal/parser"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

type LoginBonus struct {
	LoginBonusId            int32   `xorm:"pk 'login_bonus_id'"`
	LoginBonusType          int32   `xorm:"login_bonus_type"`
	StartAt                 int64   `xorm:"start_at"`
	EndAt                   int64   `xorm:"end_at"`
	BackgroundId            int32   `xorm:"background_id"`
	WhiteboardTextureAsset  *string `xorm:"varchar(3)"`
	DotUnderText            string  `json:"dot_under_text"`
	LoginBonusHandler       string  `json:"login_bonus_handler"`
	LoginBonusHandlerConfig string  `json:"login_bonus_handler_config"`
}

type LoginBonusRewardDay struct {
	LoginBonusId int32 `xorm:"pk 'login_bonus_id'"`
	Day          int32 `xorm:"pk 'day'"`
	ContentGrade int32 `enum:"LoginBonusContentGrade"`
}

type LoginBonusRewardContent struct {
	LoginBonusId int32
	Day          int32
	Content      client.Content `xorm:"extends"`
}

type LoginBonusRewardContentJson struct {
	Day           int32 `xorm:"pk"`
	Grade         int32 `enum:"LoginBonusContentGrade"`
	ContentType   int32 `xorm:"'content_type'" json:"content_type" enum:"ContentType"`
	ContentId     int32 `xorm:"'content_id'" json:"content_id"`
	ContentAmount int32 `xorm:"'content_amount'" json:"content_amount"`
}

type LoginBonusJson struct {
	Id                      int32 `json:"login_bonus_id"`
	LoginBonusType          int32 `json:"login_bonus_type"`
	StartAt                 int64
	EndAt                   int64
	BackgroundId            int32                         `json:"background_id"`
	WhiteboardTextureAsset  *string                       `json:"whiteboard_texture_asset"`
	DotUnderText            string                        `json:"dot_under_text"`
	LoginBonusHandler       string                        `json:"login_bonus_handler"`
	LoginBonusHandlerConfig string                        `json:"login_bonus_handler_config"`
	Rewards                 []LoginBonusRewardContentJson `json:"rewards"`
}

func LoadLoginBonus(path string, loginBonus *LoginBonus, loginBonusRewardDay *generic.List[LoginBonusRewardDay], loginBonusRewardContent *generic.List[LoginBonusRewardContent]) {

}

func InsertLoginBonus(session *xorm.Session, path string) {
	log.Println("Parsing login bonus file:", path)
	var loginBonus = new(LoginBonus)

	var loginBonusJson = new(LoginBonusJson)
	parser.ParseJson(path, loginBonusJson)

	loginBonus.LoginBonusId = loginBonusJson.Id
	loginBonus.LoginBonusType = loginBonusJson.LoginBonusType
	loginBonus.BackgroundId = loginBonusJson.BackgroundId
	loginBonus.WhiteboardTextureAsset = loginBonusJson.WhiteboardTextureAsset
	loginBonus.DotUnderText = loginBonusJson.DotUnderText
	loginBonus.LoginBonusHandler = loginBonusJson.LoginBonusHandler
	loginBonus.LoginBonusHandlerConfig = loginBonusJson.LoginBonusHandlerConfig

	if loginBonusJson.StartAt != 0 && loginBonusJson.EndAt != 0 {
		loginBonus.StartAt = loginBonusJson.StartAt
		loginBonus.EndAt = loginBonusJson.EndAt
	} else {
		// Assume that login bonus is always valid
		loginBonus.StartAt = 0
		loginBonus.EndAt = 1<<31 - 1
	}

	_, err := session.Table("s_login_bonus").Insert(loginBonus)
	utils.CheckErr(err)

	for _, reward := range loginBonusJson.Rewards {
		var day = LoginBonusRewardDay{}
		day.LoginBonusId = loginBonus.LoginBonusId
		day.Day = reward.Day
		day.ContentGrade = reward.Grade

		_, err = session.Table("s_login_bonus_reward_day").Insert(day)
		utils.CheckErr(err)

		var content = LoginBonusRewardContent{}
		content.LoginBonusId = loginBonus.LoginBonusId
		content.Day = reward.Day
		content.Content = client.Content{
			ContentType:   reward.ContentType,
			ContentId:     reward.ContentId,
			ContentAmount: reward.ContentAmount,
		}

		_, err = session.Table("s_login_bonus_reward_content").Insert(content)
		utils.CheckErr(err)
	}
}

func InitializeLoginBonus(session *xorm.Session) {
	path := config.ServerInitJsons + "login_bonus"

	entries, err := os.ReadDir(path)
	utils.CheckErr(err)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		InsertLoginBonus(session, filepath.Join(path, entry.Name()))
	}
}

func init() {
	addTable("s_login_bonus", LoginBonus{}, InitializeLoginBonus)
	addTable("s_login_bonus_reward_day", LoginBonusRewardDay{}, nil)
	addTable("s_login_bonus_reward_content", LoginBonusRewardContent{}, nil)
}
