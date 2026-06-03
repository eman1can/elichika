package gamedata

import (
	"fmt"
	"log"
	"sort"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/serverdata"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

// marathon event loading process:
// - load the main structure from s_event_marathon
// - load the board pictures and memos from s_event_marathon_board_thing
// - load the event stories from gamedata.StoryEventHistory
// - load the rewards from s_event_marathon_reward:
//   - this use gamedata.EventMarathonRewardGroups
//
// - load the event_ranking_topic_reward_info and event_total_topic_reward_info from the rewards:
//   - the event UR go first
//   - then the event SR that is not in final ranking
//   - then the other event SR
//
// - event_marathon_bonus_popup_order_card_mater_rows:
//   - the order seems to be gacha UR
//   - event UR
//   - gacha UR
//   - gacha UR
//   - event SR
//   - event SR
//   - gacha SR
//
// - event_marathon_bonus_popup_order_member_mater_rows is sorted by member id
type EventMarathon struct {
	serverdata.EventMarathon

	Name string

	// this is the top status template, COPY before use
	TopStatus client.EventMarathonTopStatus

	// relevant data that is can changed based one what user have done
	BoardMemos    []client.EventMarathonBoardMemorialThingsMasterRow
	BoardPictures []client.EventMarathonBoardMemorialThingsMasterRow

	// bonus mapping
	CardBonus   map[int32][]int32
	MemberBonus map[int32]int32

	// We don't need this right because we will only be using old event stories.// BoardStory []EventMarathonBoardStory

	// TODO(extra): Check if this data is available when the event start or only later on
	// TODO(extra): check and implement loop rewards
}

func (em *EventMarathon) GetNextReward(gamedata *Gamedata, eventPoint int32) (generic.Nullable[int32], generic.Nullable[client.Content]) {
	slice := em.TopStatus.EventMarathonPointRewardMasterRows.Slice
	idx := sort.Search(len(slice), func(i int) bool {
		return slice[i].RequiredPoint > eventPoint
	})

	if idx < len(slice) {
		content := gamedata.EventMarathonReward[slice[idx].RewardGroupId][0]
		return generic.NewNullable(slice[idx].RequiredPoint), generic.NewNullableFromPointer(content)
	}
	return generic.Nullable[int32]{}, generic.Nullable[client.Content]{}
}

func (em *EventMarathon) GetRankingReward(rank int32) int32 {
	for _, reward := range em.TopStatus.EventMarathonRankingRewardMasterRows.Slice {
		if (!reward.LowerRank.HasValue) || (reward.LowerRank.Value >= rank) {
			return reward.RewardGroupId
		}
	}
	log.Panic("wrong ranking reward")
	return 0
}

func (em *EventMarathon) populate(gamedata *Gamedata) {
	event := gamedata.Event[em.EventId]

	// TODO: Look into what TextureStruktur and SoundStruktur provide
	em.TopStatus = client.EventMarathonTopStatus{
		EventId: em.EventId,
		TitleImagePath: client.TextureStruktur{
			V: generic.NewNullableFromPointer(em.TitleImagePath),
		},
		BackgroundImagePath: client.TextureStruktur{
			V: generic.NewNullableFromPointer(em.BackgroundImagePath),
		},
		BgmAssetPath: client.SoundStruktur{
			V: generic.NewNullableFromPointer(em.BgmAssetPath),
		},
		GachaMasterId: event.GachaMasterId,
	}

	em.TopStatus.BoardStatus.BoardBaseImagePath.V = generic.NewNullableFromPointer(em.BoardBaseImagePath)
	em.TopStatus.BoardStatus.BoardDecoImagePath.V = generic.NewNullableFromPointer(em.BoardDecoImagePath)

	var err error
	{
		gamedata.ServerdataDb.Do(func(session *xorm.Session) {
			err = session.Table("s_event_marathon_board_thing").Where("event_id = ? AND event_marathon_board_position_type = ?", event.EventId, enum.EventMarathonBoardPositionTypeMemo).OrderBy("priority").Find(&em.BoardMemos)
		})
		utils.CheckErr(err)
		gamedata.ServerdataDb.Do(func(session *xorm.Session) {
			err = session.Table("s_event_marathon_board_thing").Where("event_id = ? AND event_marathon_board_position_type = ?", event.EventId, enum.EventMarathonBoardPositionTypePicture).OrderBy("priority").Find(&em.BoardPictures)
		})
		utils.CheckErr(err)
	}

	{
		eventStoryIds := []int32{}
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_story_event_history_detail").Where("event_master_id = ?", event.EventId).OrderBy("story_number DESC").Cols("story_event_id").Find(&eventStoryIds)
		})
		utils.CheckErr(err)
		for _, storyId := range eventStoryIds {
			em.TopStatus.StoryStatus.Stories.Append(gamedata.EventStory[storyId].GetEventMarathonStory())
		}
	}

	{
		topicRewards := []serverdata.EventTopicReward{}
		gamedata.ServerdataDb.Do(func(session *xorm.Session) {
			err = session.Table("s_event_marathon_total_topic_reward").Where("event_id = ?", event.EventId).OrderBy("display_order").Find(&topicRewards)
		})
		utils.CheckErr(err)
		for _, topicReward := range topicRewards {
			member := gamedata.Card[topicReward.RewardCardId].Member
			em.TopStatus.EventTotalTopicRewardInfo.Append(client.EventTopicReward{
				DisplayOrder: topicReward.DisplayOrder,
				RewardContent: client.Content{
					ContentType:   enum.ContentTypeCard,
					ContentId:     topicReward.RewardCardId,
					ContentAmount: topicReward.RewardCardAmount,
				},
				MainNameTopAssetPath:    member.MainNameTopAssetPath,
				MainNameBottomAssetPath: member.MainNameBottomAssetPath,
				SubNameTopAssetPath:     member.SubNameTopAssetPath,
				SubNameBottomAssetPath:  member.SubNameBottomAssetPath,
			})
		}
	}

	{
		topicRewards := []serverdata.EventTopicReward{}
		gamedata.ServerdataDb.Do(func(session *xorm.Session) {
			err = session.Table("s_event_marathon_ranking_topic_reward").Where("event_id = ?", event.EventId).OrderBy("display_order").Find(&topicRewards)
		})
		utils.CheckErr(err)
		for _, topicReward := range topicRewards {
			member := gamedata.Card[topicReward.RewardCardId].Member
			em.TopStatus.EventRankingTopicRewardInfo.Append(client.EventTopicReward{
				DisplayOrder: topicReward.DisplayOrder,
				RewardContent: client.Content{
					ContentType:   enum.ContentTypeCard,
					ContentId:     topicReward.RewardCardId,
					ContentAmount: topicReward.RewardCardAmount,
				},
				MainNameTopAssetPath:    member.MainNameTopAssetPath,
				MainNameBottomAssetPath: member.MainNameBottomAssetPath,
				SubNameTopAssetPath:     member.SubNameTopAssetPath,
				SubNameBottomAssetPath:  member.SubNameBottomAssetPath,
			})
		}
	}

	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_marathon_point_reward").Where("event_id = ?", event.EventId).OrderBy("required_point").Find(&em.TopStatus.EventMarathonPointRewardMasterRows.Slice)
	})
	utils.CheckErr(err)

	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_marathon_ranking_reward").Where("event_id = ?", event.EventId).OrderBy("ranking_reward_master_id").Find(&em.TopStatus.EventMarathonRankingRewardMasterRows.Slice)
	})
	utils.CheckErr(err)

	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_marathon_reward").Where("event_id = ?", event.EventId).OrderBy("reward_group_id").OrderBy("display_order").Find(&em.TopStatus.EventMarathonRewardMasterRows.Slice)
	})
	utils.CheckErr(err)

	{
		assetPaths := []string{}

		gamedata.ServerdataDb.Do(func(session *xorm.Session) {
			err = session.Table("s_event_marathon_rule_description_page").Where("event_id = ?", event.EventId).OrderBy("page").Cols("image_asset_path").Find(&assetPaths)
		})
		utils.CheckErr(err)
		totalPage := len(assetPaths)
		for i, assetPath := range assetPaths {
			var title string
			if gamedata.Language == "ko" {
				title = fmt.Sprintf(gamedata.Dictionary.ServerResolve("event_rule")+" %d", i+1)
			} else {
				title = fmt.Sprintf(gamedata.Dictionary.ServerResolve("event_rule")+" %d/%d", i+1, totalPage)
			}
			em.TopStatus.EventMarathonRuleDescriptionPageMasterRows.Append(
				client.EventMarathonRuleDescriptionPageMasterRow{
					Page: int32(i + 1),
					Title: client.LocalizedText{
						DotUnderText: title,
					},
					ImageAssetPath: client.TextureStruktur{
						V: generic.NewNullable[string](assetPath),
					},
				})
		}
	}

	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_marathon_bonus_popup_order_card_mater").Where("event_id = ?", event.EventId).Find(&em.TopStatus.EventMarathonBonusPopupOrderCardMaterRows.Slice)
	})
	utils.CheckErr(err)

	{
		em.CardBonus = map[int32][]int32{}
		type cardBonusValue struct {
			CardMasterId int32 `xorm:"'card_master_id'"`
			Grade        int32 `xorm:"'grade'"`
			Value        int32 `xorm:"'value'"`
		}
		bonuses := []cardBonusValue{}
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_event_marathon_bonus_card").Where("event_marathon_master_id = ?", event.EventId).Find(&bonuses)
		})
		utils.CheckErr(err)
		for _, bonus := range bonuses {
			_, exist := em.CardBonus[bonus.CardMasterId]
			if !exist {
				em.CardBonus[bonus.CardMasterId] = make([]int32, 6)
			}
			em.CardBonus[bonus.CardMasterId][bonus.Grade] = bonus.Value
		}
	}

	{
		em.MemberBonus = map[int32]int32{}
		type memberBonusValue struct {
			MemberMasterId int32 `xorm:"'member_master_id'"`
			Value          int32 `xorm:"'value'"`
		}
		bonuses := []memberBonusValue{}
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_event_marathon_bonus_member").Where("event_marathon_master_id = ?", event.EventId).Find(&bonuses)
		})
		utils.CheckErr(err)
		for _, bonus := range bonuses {
			em.MemberBonus[bonus.MemberMasterId] = bonus.Value
		}
	}

	// generate the event_marathon_bonus_popup_order_member_mater_rows field, which are always sorted on member_matser_id (typo)
	for memberId := range em.MemberBonus {
		em.TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Append(client.EventMarathonBonusPopupOrderMemberMaterRow{
			MemberMatserId: memberId,
			DisplayLine:    3,
			DisplayOrder:   memberId,
		})
	}
	sort.Slice(em.TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Slice, func(i, j int) bool {
		return em.TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Slice[i].DisplayOrder <
			em.TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Slice[j].DisplayOrder
	})
}

func loadEventMarathon(gamedata *Gamedata) {
	gamedata.EventMarathon = make(map[int32]*EventMarathon)

	var err error
	var events []serverdata.EventMarathon
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_marathon").Find(&events)
	})
	utils.CheckErr(err)

	for _, event := range events {
		eventMarathon := EventMarathon{
			EventMarathon: event,
			Name:          fmt.Sprintf("m.event_marathon_title_%d", event.EventId),
		}
		eventMarathon.populate(gamedata)
		gamedata.EventMarathon[event.EventId] = &eventMarathon
	}
}

func init() {
	addLoadFunc(loadEventMarathon)
	addPrequisite(loadEventMarathon, loadCard)
	addPrequisite(loadEventMarathon, loadEvent)
	addPrequisite(loadEventMarathon, loadEventStory)
}
