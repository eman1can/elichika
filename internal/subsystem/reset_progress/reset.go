package reset_progress

import (
	"elichika/internal/item"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func RemoveUserProgress(session *userdata.Session, table string) {
	count, err := session.Db.Table(table).Where("user_id = ?", session.UserId).Delete()
	utils.CheckErr(err)
	if table == "u_story_event_history" {
		user_content.AddContent(session, item.MemoryKey.Amount(int32(count)))
	}
}

func MarkIsNew(session *userdata.Session, table string, isNew bool) {
	_, err := session.Db.Table(table).Where("user_id = ?", session.UserId).Update(map[string]interface{}{"is_new": isNew})
	utils.CheckErr(err)

}
