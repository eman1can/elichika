package gamedata

import (
	"log"

	"elichika/internal/generic/drop"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

type Lesson struct {
	ItemAmount map[int32]*drop.WeightedDropList[int32] `xorm:"-"`

	DefaultSkillDrop map[int32]*drop.WeightedDropList[int32] `xorm:"-"`
	SkillPosition    *drop.WeightedDropList[int32]           `xorm:"-"`
}

func (lm *Lesson) populate(gamedata *Gamedata) {
	type LessonEnhancingItemDropRate struct {
		LessonEnhancingItemId int32
		TargetRarity          int32
		MagnificationWeight   int32
	}
	var err error
	var enhancingItems []LessonEnhancingItemDropRate
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_lesson_enhancing_item_effect_drop_rate").Find(&enhancingItems)
	})
	utils.CheckErr(err)

	type LessonSkillContent struct {
		SkillMasterId int32
		Rarity        int32
		DropType      int32
		LessonMenuId1 int32
		LessonMenuId2 int32
	}

	var lessonSkillContent []LessonSkillContent
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_lesson_skill_content").Find(&lessonSkillContent)
	})
	utils.CheckErr(err)

	type SkillRarityDropWeight struct {
		Rarity int32
		Weight int32
	}

	var skillRarityDropWeight []SkillRarityDropWeight
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_lesson_skill_rarity").Find(&skillRarityDropWeight)
	})
	utils.CheckErr(err)

	rarityWeights := map[int32]int32{}
	for _, rarity := range skillRarityDropWeight {
		rarityWeights[rarity.Rarity] = rarity.Weight
	}

	lm.DefaultSkillDrop = map[int32]*drop.WeightedDropList[int32]{}
	for lm1 := 1; lm1 <= 9; lm1++ {
		for lm2 := 1; lm2 <= 9; lm2++ {
			for lm3 := 1; lm3 <= 9; lm3++ {
				key := int32(lm1*100 + lm2*10 + lm3)
				if lm.DefaultSkillDrop[key] == nil {
					lm.DefaultSkillDrop[key] = &drop.WeightedDropList[int32]{}
				}

				// Calculate the dropped skills and the rarity distribution for this lesson combination.
				rarities := map[int32]int32{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
				skills := map[int32]int32{}
				for _, skill := range lessonSkillContent {
					doDropSkill := false
					switch skill.DropType {
					case 1: // Pure: all lessons are ID 1
						doDropSkill = skill.LessonMenuId1 == int32(lm1) && lm1 == lm2 && lm2 == lm3
						break
					case 2: // Mixed: at least one lesson matches ID 1
						doDropSkill = skill.LessonMenuId1 == int32(lm1)
						doDropSkill = doDropSkill || skill.LessonMenuId1 == int32(lm2)
						doDropSkill = doDropSkill || skill.LessonMenuId1 == int32(lm3)
						break
					case 3: // Skill always drops
						doDropSkill = true
						break
					case 4: // Majority / Minority: Lessons match AAB, ABA, BAA
						doDropSkill = skill.LessonMenuId1 == int32(lm1) && lm1 == lm2 && skill.LessonMenuId2 == int32(lm3)
						doDropSkill = doDropSkill || (skill.LessonMenuId1 == int32(lm1) && skill.LessonMenuId2 == int32(lm2) && lm1 == lm3)
						doDropSkill = doDropSkill || (skill.LessonMenuId2 == int32(lm1) && lm2 == lm3 && skill.LessonMenuId1 == int32(lm3))
						break
					default: // Other Drop Types Not Implemented (There should be no others)
						doDropSkill = false
					}

					if doDropSkill {
						rarities[skill.Rarity] += 1
						skills[skill.SkillMasterId] = skill.Rarity
					}
				}

				// Add a "no skill drop" item with the appropriate weight
				lm.DefaultSkillDrop[key].AddItem(0, rarityWeights[0])

				// Add each dropped skill to the drop list with the appropriate weight
				for skillMasterId, rarity := range skills {
					weight := rarityWeights[rarity] / rarities[rarity]
					lm.DefaultSkillDrop[key].AddItem(skillMasterId, weight)
				}
			}
		}
	}

	// TODO: Test / Implement enhancing items for skill drops
	type SkillMemberChance struct {
		PositionId int32
		Weight     int32
	}

	var skillMemberChance []SkillMemberChance
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_lesson_skill_member_chance").Find(&skillMemberChance)
	})
	utils.CheckErr(err)

	lm.SkillPosition = &drop.WeightedDropList[int32]{}
	for _, skillMemberChance := range skillMemberChance {
		lm.SkillPosition.AddItem(skillMemberChance.PositionId, skillMemberChance.Weight)
	}

	type LessonDropAmount struct {
		ItemId int32
		Count  int32
		Weight int32
	}

	var lessonDropAmount []LessonDropAmount
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_lesson_drop_amount").Find(&lessonDropAmount)
	})
	utils.CheckErr(err)

	lm.ItemAmount = map[int32]*drop.WeightedDropList[int32]{}
	for _, dropAmount := range lessonDropAmount {
		if lm.ItemAmount[dropAmount.ItemId] == nil {
			lm.ItemAmount[dropAmount.ItemId] = &drop.WeightedDropList[int32]{}
		}
		lm.ItemAmount[dropAmount.ItemId].AddItem(dropAmount.Count, dropAmount.Weight)
	}
}

func loadLesson(gamedata *Gamedata) {
	log.Println("Loading Lesson")
	lesson := Lesson{}
	lesson.populate(gamedata)
	gamedata.Lesson = &lesson
}

func init() {
	addLoadFunc(loadLesson)
}
