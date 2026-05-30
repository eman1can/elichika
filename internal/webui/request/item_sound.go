package request

type WebUIItemSoundRequest struct {
	VoiceId  int32  `form:"id"`
	Language string `form:"l"`
}
