package client

import "elichika/internal/generic"

type OtherUserDeckDetail struct {
	Deck             OtherUserDeck                    `json:"deck"`
	MemberLoveLevels generic.Dictionary[int32, int32] `json:"member_love_levels"`
}
