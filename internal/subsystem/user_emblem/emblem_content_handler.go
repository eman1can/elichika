package user_emblem

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

// TODO(emblem): This doesn't set the params, and maybe there should be a function that manually add with params
// But the params doesn't seems to do anything
func emblemContentHandler(session *userdata.Session, content *client.Content) any {
	content.ContentAmount = 0
	exists, err := session.Db.Table("u_emblem").
		Where("user_id = ? AND emblem_m_id = ?", session.UserId, content.ContentId).Exist(&client.UserEmblem{})
	utils.CheckErr(err)
	if !exists {
		userdata.GenericDatabaseInsert(session, "u_emblem", client.UserEmblem{
			EmblemMId:  content.ContentId,
			IsNew:      true,
			AcquiredAt: session.Time.Unix(),
		})
	}
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeEmblem, emblemContentHandler)
}
