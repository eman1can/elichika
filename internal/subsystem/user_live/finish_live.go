package user_live

import (
	"log"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func FinishLive(session *userdata.Session, req request.FinishLiveRequest) response.FinishLiveResponse {
	// this is pretty different for different type of live
	// for simplicity we just read the request and call different handlers, even though we might be able to save some extra work

	exist, live, startReq := LoadUserLive(session)
	utils.MustExist(exist)
	ClearUserLive(session)
	switch live.LiveType {
	case enum.LiveTypeManual:
		return liveTypeManualHandler(session, req, live, startReq)
	case enum.LiveTypeTower:
		return liveTypeTowerHandler(session, req, live, startReq)
	default:
		log.Panic("not handled")
		return response.FinishLiveResponse{}
	}
}
