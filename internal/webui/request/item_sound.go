package request

type WebUIItemSoundRequest struct {
	VoiceId      int32 `form:"id"`
	ReleaseRoute int32 `form:"route"`
	ReleaseValue int32 `form:"value"`
}
