package gamedata

import (
	"log"

	"elichika/internal/utils"

	"xorm.io/xorm"
)

type NaviVoice struct {
	Id                    int32  `xorm:"pk 'id'"`
	MemberMId             int32  `xorm:"'member_m_id'"`
	Name                  string `xorm:"'name'"`
	NaviVoiceReleaseRoute int32  `xorm:"'navi_voice_release_route'"`
	NaviVoiceReleaseValue int32  `xorm:"'navi_voice_release_value'"`
	DisplayOrder          int32  `xorm:"'display_order'"`
	ListType              int32  `xorm:"'list_type'" enum:"navi_voice_list_type"`
}

func loadNaviVoice(gamedata *Gamedata) {
	log.Println("Loading NaviVoice")
	gamedata.NaviVoice = make(map[int32]*NaviVoice)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_navi_voice").Find(&gamedata.NaviVoice)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadNaviVoice)
}
