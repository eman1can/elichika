package user_present

import (
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func FetchPresentCount(session *userdata.Session) int32 {
	count, err := session.Db.Table("u_present_item").Where("user_id = ?", session.UserId).Count()
	utils.CheckErr(err)
	return int32(count)
}
