package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchMemberGuildRankingResponse struct {
	MemberGuildRanking         client.MemberGuildRanking                   `json:"member_guild_ranking"`
	MemberGuildUserRankingList generic.List[client.MemberGuildUserRanking] `json:"member_guild_user_ranking_list"`
}
