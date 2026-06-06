package auth

import (
	"time"

	"elichika/internal/serverstate"

	"github.com/alexedwards/scs/v2"
	"xorm.io/xorm"
)

// SessionStore is a minimal scs.Store backed by the userdata DB.
type SessionStore struct{}

func initSessionStore() scs.Store {
	return &SessionStore{}
}

func (s *SessionStore) Find(token string) ([]byte, bool, error) {
	var record serverstate.WebUISession

	var data []byte
	var exist bool
	var err error
	serverstate.Database.Do(func(session *xorm.Session) {
		exist, err = session.Table("s_webui_sessions").Where("token = ? AND expiry > ?", token, time.Now().UTC()).Get(&record)
	})

	if err == nil && exist {
		data = record.Data
	}

	return data, exist, err
}

func (s *SessionStore) Commit(token string, b []byte, expiry time.Time) error {
	var err error
	serverstate.Database.Do(func(session *xorm.Session) {
		exists, err := session.Table("s_webui_sessions").Where("token = ?", token).Exist()
		if err != nil {
			return
		}

		record := serverstate.WebUISession{Token: token, Data: b, Expiry: expiry}
		if exists {
			_, err = session.Table("s_webui_sessions").Where("token = ?", token).Update(&record)
		} else {
			_, err = session.Table("s_webui_sessions").Insert(&record)
		}
		if err != nil {
			return
		}
		err = session.Commit()
	})
	return err
}

func (s *SessionStore) Delete(token string) error {
	var err error
	serverstate.Database.Do(func(session *xorm.Session) {
		_, err := session.Table("s_webui_sessions").Where("token = ?", token).Delete(&serverstate.WebUISession{})
		if err != nil {
			return
		}
		err = session.Commit()
	})
	return err
}
