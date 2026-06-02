package user_lesson

import (
	"math"
	"reflect"
	"sort"

	"elichika/internal/client"
	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/config"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/item"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/subsystem/user_lesson_deck"
	"elichika/internal/subsystem/user_member_guild"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/subsystem/user_status"
	"elichika/internal/subsystem/user_subscription_status"
	"elichika/internal/userdata"
)

func ExecuteLesson(session *userdata.Session, req request.ExecuteLessonRequest) response.ExecuteLessonResponse {
	resp := response.ExecuteLessonResponse{
		UserModelDiff: &session.UserModel,
	}

	result := response.LessonResultResponse{
		SelectedDeckId: req.SelectedDeckId,
	}

	deck := user_lesson_deck.GetUserLessonDeck(session, req.SelectedDeckId)
	repeatCount := int32(1)
	if req.IsThreeTimes {
		repeatCount = 3
	}
	if config.Conf.ResourceConfig().ConsumeAp {
		user_status.AddUserActivityPoints(session, -repeatCount)
	}
	user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountLesson, nil, nil,
		func(session *userdata.Session, missionList []any, _ ...any) {
			for _, mission := range missionList {
				user_mission.AddMissionProgress(session, mission, repeatCount)
			}
		})

	session.UserStatus.LessonResumeStatus = enum.TopPriorityProcessStatusLesson

	enhancingItems := map[int32]*client.Content{}

	for _, itemId := range req.ConsumedContentIds.Slice {
		enhancingItem := user_content.GetUserContent(session, enum.ContentTypeLessonEnhancingItem, itemId)
		enhancingItems[itemId] = &enhancingItem
	}

	resp.IsSubscription = user_subscription_status.HasSubscription(session)

	for lesson := int32(1); lesson <= 4; lesson++ {
		actions := generic.List[client.LessonMenuAction]{}
		for i := 1; i <= 9; i++ {
			cardMasterId := reflect.ValueOf(deck).Field(i + 1).Interface().(generic.Nullable[int32]).Value
			actions.Append(client.LessonMenuAction{
				CardMasterId: cardMasterId,
				Position:     int32(i),
			})
		}
		resp.LessonMenuActions.Set(lesson%4, actions)
		resp.LessonDropRarityList.Set(lesson%4, generic.List[int32]{})
	}

	isMemberGuildRankingPeriod := user_member_guild.IsMemberGuildRankingPeriod(session)

	for repeat := int32(1); repeat <= repeatCount; repeat++ {
		var usedItems []int32
		for _, itemId := range req.ConsumedContentIds.Slice {
			if enhancingItems[itemId].ContentAmount > 0 {
				enhancingItems[itemId].ContentAmount--
				usedItems = append(usedItems, itemId)
			}
		}

		key := req.ExecuteLessonIds.Slice[0]*100 + req.ExecuteLessonIds.Slice[1]*10 + req.ExecuteLessonIds.Slice[2]

		dropCount := session.Gamedata.Lesson.ItemAmount[enum.NormalDropType].GetRandomItem()

		// Calculate Item Drops
		for lesson := int32(1); lesson <= 3; lesson++ {
			lessonMenu := session.Gamedata.LessonMenu[req.ExecuteLessonIds.Slice[lesson-1]]
			dropList := lessonMenu.DefaultItemDrop
			for _, enhancedItemId := range usedItems {
				enhancedDropList, exist := lessonMenu.ItemDrop[enhancedItemId]
				if exist {
					dropList = enhancedDropList
				}
			}

			dropRarityList := resp.LessonDropRarityList.GetOnly(lesson)
			var gainedItems []client.LessonDropItem

			// Calculate Normal Item Drops
			var lessonDropCount int32
			if lesson == 3 {
				lessonDropCount = dropCount - int32(math.Round(float64(dropCount)/3.0))*2
			} else {
				lessonDropCount = int32(math.Round(float64(dropCount) / 3.0))
			}

			for i := int32(0); i < lessonDropCount; i++ {
				dropItem := dropList.GetRandomItem()
				gainedItems = append(gainedItems, dropItem)
			}

			// Calculate Megaphone Drops if ranking is on
			if isMemberGuildRankingPeriod {
				megaphoneCount := session.Gamedata.Lesson.ItemAmount[enum.MegaphoneDropType].GetRandomItem()
				for i := int32(0); i < megaphoneCount; i++ {
					gainedItems = append(gainedItems, client.LessonDropItem{
						ContentType:   item.RallyMegaphone.ContentType,
						ContentId:     item.RallyMegaphone.ContentId,
						ContentAmount: item.RallyMegaphone.ContentAmount,
						DropRarity:    4, // this field is not enum
					})
				}
			}

			// Award items to user
			for _, content := range gainedItems {
				user_content.AddContent(session, client.Content{
					ContentType:   content.ContentType,
					ContentId:     content.ContentId,
					ContentAmount: content.ContentAmount,
				})
				result.DropItemList.Append(content)
				dropRarityList.Append(min(2, content.DropRarity))
			}

			// If subscription, double items
			if resp.IsSubscription {
				for _, content := range gainedItems {
					content.IsSubscription = true

					user_content.AddContent(session, client.Content{
						ContentType:   content.ContentType,
						ContentId:     content.ContentId,
						ContentAmount: content.ContentAmount,
					})
					result.DropItemList.Append(content)
					dropRarityList.Append(min(2, content.DropRarity))
				}
			}
		}

		// Calculate Skill Drop
		// TODO: Enhancing items?
		skillMasterId := session.Gamedata.Lesson.DefaultSkillDrop[key].GetRandomItem()
		if skillMasterId != 0 {
			memberPositionId := session.Gamedata.Lesson.SkillPosition.GetRandomItem()
			result.DropSkillList.Append(client.LessonResultDropPassiveSkill{
				Position:       memberPositionId,
				PassiveSkillId: skillMasterId,
			})
		}

		if (repeat == 1) && (repeat < repeatCount) {
			sort.Slice(req.ExecuteLessonIds.Slice, func(i, j int) bool {
				return req.ExecuteLessonIds.Slice[i] < req.ExecuteLessonIds.Slice[j]
			})
		}
	}

	for _, enhancingItem := range enhancingItems {
		user_content.UpdateUserContent(session, *enhancingItem)
	}

	userdata.GenericDatabaseInsert(session, "u_lesson", result)

	return resp
}
