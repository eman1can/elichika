package userdata

import (
	"elichika/internal/userdata/database"

	"xorm.io/xorm"
)

var (
	Engine *xorm.Engine
)

func Init() {
	database.Init()
	Engine = database.Engine
	database.InitTables(Engine)
}
