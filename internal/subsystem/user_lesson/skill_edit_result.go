package user_lesson

import (
	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_card"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func SkillEditResult(session *userdata.Session, req request.SkillEditResultRequest) {
	session.UserStatus.LessonResumeStatus = enum.TopPriorityProcessStatusNone

	_, err := session.Db.Table("u_lesson").Where("user_id = ?", session.UserId).Delete(&response.LessonResultResponse{})
	utils.CheckErr(err)

	for cardMasterId, selectedSkills := range req.SelectedSkillIds.Map {
		card := user_card.GetUserCard(session, cardMasterId)
		for i, skillId := range selectedSkills.Slice {
			switch i {
			case 0:
				card.AdditionalPassiveSkill1Id = skillId
			case 1:
				card.AdditionalPassiveSkill2Id = skillId
			case 2:
				card.AdditionalPassiveSkill3Id = skillId
			case 3:
				card.AdditionalPassiveSkill4Id = skillId
			}
		}
		user_card.UpdateUserCard(session, card) // this is always updated even if no skill change happen
	}

}
