package info_trigger

import "elichika/internal/server"

func init() {
	// the request / response is exactly the same, there's only different in the end point
	server.AddHandler("/", "POST", "/infoTrigger/readEventMarathonResult", read)
}
