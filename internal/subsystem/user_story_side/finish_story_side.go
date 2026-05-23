package user_story_side

import (
	"log"

	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func FinishStorySide(session *userdata.Session, storySideMasterId int32) {
	userStorySide := client.UserStorySide{}
	exist, err := session.Db.Table("u_story_side").Where("user_id = ? AND story_side_master_id = ?",
		session.UserId, storySideMasterId).Get(&userStorySide)
	utils.CheckErr(err)
	if !exist {
		log.Panic("side story must exist to be read")
	}
	if !userStorySide.IsNew { // already read
		return
	}
	userStorySide.IsNew = false
	session.UserModel.UserStorySideById.Set(storySideMasterId, userStorySide)
}
