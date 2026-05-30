package user_voice

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetUserVoice(session *userdata.Session, naviVoiceMasterId int32) client.UserVoice {
	ptr, exist := session.UserModel.UserVoiceByVoiceId.Get(naviVoiceMasterId)
	if exist {
		return *ptr
	}

	voice := client.UserVoice{}
	exist, err := session.Db.Table("u_voice").Where("user_id = ? AND navi_voice_master_id = ?", session.UserId, naviVoiceMasterId).Get(&voice)
	utils.CheckErr(err)

	if !exist {
		voice = client.UserVoice{
			NaviVoiceMasterId: naviVoiceMasterId,
			IsNew:             true,
		}
	}

	return voice
}
