package serverstate

import "time"

// WebUISession maps to the webui_sessions table.
type WebUISession struct {
	Token  string    `xorm:"pk varchar(43) 'token'"`
	Data   []byte    `xorm:"blob 'data'"`
	Expiry time.Time `xorm:"datetime 'expiry'"`
}

func init() {
	addTable("s_webui_sessions", WebUISession{}, nil)
}
