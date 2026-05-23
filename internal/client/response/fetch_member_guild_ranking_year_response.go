package response

import "elichika/internal/client"

type FetchMemberGuildRankingYearResponse struct {
	MemberGuildRanking client.MemberGuildRanking `json:"member_guild_ranking"`
}
