package gamedata

import (
	"fmt"
	"log"
	"math/rand"
	"sort"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/serverdata"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

type EventMining struct {
	serverdata.EventMining

	Name string

	// this is the top status template, COPY before use
	TopStatus       client.EventMiningTopStatus
	EventAnimation  *client.EventMiningTopAnimationCellMasterRow
	GachaAnimations []client.EventMiningTopAnimationCellMasterRow

	// bonus mapping
	CardBonus map[int32][]int32

	// event music for bootstrap
	// note that .EndAt is not filled in, boot strap will fill it
	EventMusics generic.Array[client.LiveEventMiningVoltageMusicBadge]
	Trade       *client.Trade
}

func (em *EventMining) GetPointRankingReward(rank int32) int32 {
	for _, reward := range em.TopStatus.EventMiningPointRankingRewardMasterRows.Slice {
		if (!reward.LowerRank.HasValue) || (reward.LowerRank.Value >= rank) {
			return reward.RewardGroupId
		}
	}
	log.Panic("wrong ranking reward")
	return 0
}

func (em *EventMining) GetVoltageRankingReward(rank int32) int32 {
	for _, reward := range em.TopStatus.EventMiningVoltageRankingRewardMasterRows.Slice {
		if (!reward.LowerRank.HasValue) || (reward.LowerRank.Value >= rank) {
			return reward.RewardGroupId
		}
	}
	log.Panic("wrong ranking reward")
	return 0
}

func (em *EventMining) HasAnimation() bool {
	return em.EventAnimation != nil
}

func (em *EventMining) EventCharacterAnimation() client.EventMiningTopAnimationCellMasterRow {
	return *em.EventAnimation
}

func (em *EventMining) RandomGachaCharacterAnimation() client.EventMiningTopAnimationCellMasterRow {
	return em.GachaAnimations[rand.Intn(2)]
}

func (em *EventMining) populate(gamedata *Gamedata) {
	event := gamedata.Event[em.EventId]

	em.TopStatus = client.EventMiningTopStatus{
		EventId: event.EventId,
		TitleImagePath: client.TextureStruktur{
			V: generic.NewNullable(em.TitleImagePath),
		},
		BackgroundImagePath: client.TextureStruktur{
			V: generic.NewNullable(em.BackgroundImagePath),
		},
		BgmAssetPath: client.SoundStruktur{
			V: generic.NewNullable(em.BgmAssetPath),
		},
		GachaMasterId:      event.GachaMasterId,
		EventPointMasterId: event.EventId - 11000, // this need to be hard coded due to various reasons
		SelectionAmount:    em.EventCompetitionSelectionAmount,
		TradeMasterId:      generic.NewNullable[int32](event.EventId),
	}
	for _, liveId := range em.EventCompetitionLiveIds {
		em.TopStatus.EventMiningCompetitionMasterRows.Append(client.EventMiningCompetitionMasterRow{
			LiveId: liveId,
		})

		em.EventMusics.Append(client.LiveEventMiningVoltageMusicBadge{
			LiveMasterId: generic.NewNullable(liveId),
			AppealText: client.LocalizedText{
				// this is not packet capture, just eyeballing and comparing to old video
				// should be good enough
				DotUnderText: gamedata.Dictionary.ServerResolve("event_mining_music_appeal_text"),
				// DotUnderText: `<img src="Common/InlineImage/Icon/tex_inlineimage_icon_rank_vo_1" width="40px" height="40px" offsety="-12px"/>Voltage Ranking Song`,
			},
			EndAt: 0,
		})
	}

	var err error
	{
		eventStoryIds := []int32{}
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_story_event_history_detail").Where("event_master_id = ?", event.EventId).OrderBy("story_number DESC").Cols("story_event_id").Find(&eventStoryIds)
		})
		utils.CheckErr(err)
		for _, storyId := range eventStoryIds {
			em.TopStatus.StoryStatus.Stories.Append(gamedata.EventStory[storyId].GetEventMiningStory())
		}
	}

	topicRewards := []serverdata.EventTopicReward{}
	{
		gamedata.ServerdataDb.Do(func(session *xorm.Session) {
			err = session.Table("s_event_mining_point_ranking_topic_reward").Where("event_id = ?", event.EventId).OrderBy("display_order").Find(&topicRewards)
		})
		utils.CheckErr(err)
		for _, topicReward := range topicRewards {
			member := gamedata.Card[topicReward.RewardCardId].Member
			em.TopStatus.EventPointRankingTopicRewardInfo.Append(client.EventMiningTopicReward{
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
				RankingCategory:         enum.EventRankingCategoryPoint,
			})
		}
	}

	em.TopStatus.EventVoltageRankingTopicRewardInfo.Append(client.EventMiningTopicReward{
		DisplayOrder: 1,
		RewardContent: client.Content{
			ContentType:   12,
			ContentId:     1700,
			ContentAmount: 1,
		},
		RankingCategory: enum.EventRankingCategoryVoltage,
		SinglePictureAssetPath: client.TextureStruktur{
			V: generic.NewNullable(em.VoltageRankingTopicRewardAssetPath),
		},
	})

	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_mining_point_ranking_reward").Where("event_id = ?", event.EventId).OrderBy("ranking_reward_master_id").Find(&em.TopStatus.EventMiningPointRankingRewardMasterRows.Slice)
	})
	utils.CheckErr(err)
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_mining_voltage_ranking_reward").Where("event_id = ?", event.EventId).OrderBy("ranking_reward_master_id").Find(&em.TopStatus.EventMiningVoltageRankingRewardMasterRows.Slice)
	})
	utils.CheckErr(err)

	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_mining_reward").Where("event_id = ?", event.EventId).OrderBy("reward_group_id").OrderBy("display_order").Find(&em.TopStatus.EventMiningRewardMasterRows.Slice)
	})
	utils.CheckErr(err)

	// animations
	topAnimationCells := []serverdata.EventMiningTopAnimationCell{}
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_mining_top_animation_cell").Where("event_id = ?", event.EventId).Find(&topAnimationCells)
	})
	utils.CheckErr(err)
	for _, cell := range topAnimationCells {
		popupComment := generic.Nullable[client.LocalizedText]{}
		if cell.PopupComment != nil {
			popupComment = generic.NewNullable(client.LocalizedText{
				DotUnderText: *cell.PopupComment,
				// TODO(extra): if ths is ever used, it should be localised
			})
		}
		if cell.IsGacha {
			em.GachaAnimations = append(em.GachaAnimations, client.EventMiningTopAnimationCellMasterRow{
				EventMiningMasterId: event.EventId,
				ThumbnailCellId:     cell.ThumbnailCellId,
				MovieAssetPath: client.MovieStruktur{
					V: cell.MovieAssetPath,
				},
				IsPopup: cell.IsPopup,
				PopupMovieAssetPath: client.MovieStruktur{
					V: cell.PopupMovieAssetPath,
				},
				PopupComment: popupComment,
				IsBig:        cell.IsBig,
			})
		} else {
			em.EventAnimation = &client.EventMiningTopAnimationCellMasterRow{
				EventMiningMasterId: event.EventId,
				ThumbnailCellId:     cell.ThumbnailCellId,
				MovieAssetPath: client.MovieStruktur{
					V: cell.MovieAssetPath,
				},
				IsPopup: cell.IsPopup,
				PopupMovieAssetPath: client.MovieStruktur{
					V: cell.PopupMovieAssetPath,
				},
				PopupComment: popupComment,
				IsBig:        cell.IsBig,
			}
		}
	}

	// stills

	// EventMiningTopStillCellMasterRows
	// we will load all the cells in here, then filter out irrelevant data later
	topStillCells := []serverdata.EventMiningTopStillCell{}
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_mining_top_still_cell").Where("event_id = ?", event.EventId).OrderBy("add_story_number").Find(&topStillCells)
	})
	utils.CheckErr(err)
	for _, cell := range topStillCells {
		popupComment := generic.Nullable[client.LocalizedText]{}
		if cell.PopupComment != nil {
			popupComment = generic.NewNullable(client.LocalizedText{
				DotUnderText: *cell.PopupComment,
				// TODO(extra): if ths is ever used, it should be localised
			})
		}
		em.TopStatus.EventMiningTopStillCellMasterRows.Append(client.EventMiningTopStillCellMasterRow{
			EventMiningMasterId: event.EventId,
			ThumbnailCellId:     cell.ThumbnailCellId,
			AddStoryNumber:      cell.AddStoryNumber,
			Priority:            cell.Priority,
			ImageThumbnailAssetPath: client.TextureStruktur{
				V: generic.NewNullable[string](cell.ImageThumbnailAssetPath),
			},
			IsPopup: cell.IsPopup,
			PopupImageThumbnailAssetPath: client.TextureStruktur{
				V: generic.NewNullable[string](cell.PopupImageThumbnailAssetPath),
			},
			PopupComment: popupComment,
			IsBig:        cell.IsBig,
		})
	}

	// EventMiningTopStillSubCellMasterRows
	topStillSubCells := []serverdata.EventMiningTopStillSubCell{}
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_mining_top_still_sub_cell").Where("event_id = ?", event.EventId).OrderBy("thumbnail_cell_id").Find(&topStillSubCells)
	})

	for _, cell := range topStillSubCells {
		em.TopStatus.EventMiningTopStillSubCellMasterRows.Append(client.EventMiningTopStillSubCellMasterRow{
			EventMiningMasterId: event.EventId,
			PanelSetId:          cell.PanelSetId,
			ThumbnailCellId:     cell.ThumbnailCellId,
			Priority:            cell.Priority,
			ImageThumbnailAssetPath: client.TextureStruktur{
				V: generic.NewNullable[string](cell.ImageThumbnailAssetPath),
			},
			IsPopup: cell.IsPopup,
			PopupImageThumbnailAssetPath: client.TextureStruktur{
				V: generic.NewNullable[string](cell.PopupImageThumbnailAssetPath),
			},
		})
	}

	{
		assetPaths := []string{}
		gamedata.ServerdataDb.Do(func(session *xorm.Session) {
			err = session.Table("s_event_mining_rule_description_page").Where("event_id = ?", event.EventId).OrderBy("page").Cols("image_asset_path").Find(&assetPaths)
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
			em.TopStatus.EventMiningRuleDescriptionPageMasterRows.Append(
				client.EventMiningRuleDescriptionPageMasterRow{
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
		err = session.Table("s_event_mining_bonus_popup_order_card_mater").Where("event_id = ?", event.EventId).OrderBy("display_order").
			Find(&em.TopStatus.EventMiningBonusPopupOrderCardMaterRows.Slice)
	})
	utils.CheckErr(err)

	// Done with top status, loading helper data

	{
		em.CardBonus = map[int32][]int32{}
		type cardBonusValue struct {
			CardMasterId int32 `xorm:"'card_master_id'"`
			Grade        int32 `xorm:"'grade'"`
			Value        int32 `xorm:"'value'"`
		}
		bonuses := []cardBonusValue{}
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_event_mining_bonus_card").Where("event_mining_master_id = ?", event.EventId).Find(&bonuses)
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
	// EventMiningTopLikeMemerRows: the member that show up when you click on a cell
	// official server send this data for every event, which is good for us to study how it work:
	// - the thing is not very complicated
	// - with the exception of some specific event, all the people who's featured in an event will like all the cell that show up in that event
	// - event with special likeing: 31010, 31011, 31012, 31016, 31021
	// - it is impossible to determine which cell is which, and those event also follow the above rule, so we just keep the rule
	// we will also only send the relevant data, as sending more just waste network
	memberIds := []int{}
	for cardMasterId := range em.CardBonus {
		member := gamedata.Card[cardMasterId].Member
		memberIds = append(memberIds, int(member.Id))
	}
	sort.Ints(memberIds)
	for thumbnailCellId := int32(16); thumbnailCellId <= 26; thumbnailCellId++ {
		thumbnailCellIdWithEvent := thumbnailCellId
		if event.EventId > 31001 {
			thumbnailCellIdWithEvent = event.EventId*100 + thumbnailCellId
		}
		for _, memberId := range memberIds {
			em.TopStatus.EventMiningTopLikeMemerRows.Append(client.EventMiningTopLikeMemerRow{
				ThumbnailCellId: thumbnailCellIdWithEvent,
				MemberId:        int32(memberId),
			})
		}
	}

	// trades
	bannerImagePath := ""
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		_, err = session.Table("m_story_event_history").Where("id = ?", event.EventId).Cols("banner_image_asset_path").
			Get(&bannerImagePath)
		utils.CheckErr(err)
	})

	em.Trade = &client.Trade{
		TradeId: em.TopStatus.TradeMasterId.Value,
		BannerImagePath: client.TextureStruktur{
			V: generic.NewNullable(bannerImagePath),
		},
		SourceContentType: enum.ContentTypeExchangeEventPoint,
		SourceContentId:   em.TopStatus.EventPointMasterId,
		SourceThumbnailAssetPath: client.TextureStruktur{
			V: generic.NewNullable(""), // filled later
		},
		MonthlyReset: false,
	}
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		_, err = session.Table("m_exchange_event_point").Where("id = ?", em.Trade.SourceContentId).Cols("icon_asset_path").
			Get(&em.Trade.SourceThumbnailAssetPath.V.Value)
	})
	utils.CheckErr(err)
	tradeProducts := []serverdata.EventMiningTradeProduct{}
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_mining_trade_product").Where("event_id = ?", em.EventId).OrderBy("product_id").Find(&tradeProducts)
	})
	utils.CheckErr(err)
	for _, product := range tradeProducts {
		tradeProduct := client.TradeProduct{
			ProductId:    product.ProductId,
			TradeId:      em.Trade.TradeId,
			SourceAmount: product.SourceAmount,
			StockAmount:  generic.NewNullableFromPointer(product.StockAmount),
		}
		tradeProduct.Contents.Append(product.Content)
		gamedata.TradeProduct[tradeProduct.ProductId] = &tradeProduct
		em.Trade.Products.Append(tradeProduct)
	}
}

func loadEventMining(gamedata *Gamedata) {
	gamedata.EventMining = make(map[int32]*EventMining)

	var err error
	var events []serverdata.EventMining
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event_mining").Find(&events)
	})
	utils.CheckErr(err)

	for _, event := range events {
		eventMining := EventMining{
			EventMining: event,
			Name:        fmt.Sprintf("m.event_mining_title_%d", event.EventId),
		}
		eventMining.populate(gamedata)
		gamedata.EventMining[event.EventId] = &eventMining
	}
}

func init() {
	addLoadFunc(loadEventMining)
	addPrequisite(loadEventMining, loadCard)
	addPrequisite(loadEventMining, loadEvent)
	addPrequisite(loadEventMining, loadEventStory)
	addPrequisite(loadEventMining, loadTrade)
}
