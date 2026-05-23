package user_lesson

import (
	"elichika/internal/client/response"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func ResultLesson(session *userdata.Session) response.LessonResultResponse {
	resp := response.LessonResultResponse{}
	exists, err := session.Db.Table("u_lesson").Where("user_id = ?", session.UserId).Get(&resp)
	utils.CheckErrMustExist(err, exists)

	resp.UserModelDiff = &session.UserModel
	return resp
}
