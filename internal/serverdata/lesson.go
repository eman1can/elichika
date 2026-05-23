package serverdata

type LessonSkillDrop struct {
	SkillMasterId int32 `xorm:"pk 'skill_master_id'"`
	GroupId       int32 `xorm:"'group_id'"`
	Rarity        int32 `xorm:"'rarity'"`
}

type LessonSkillDropGroup struct {
	GroupId   int32 `xorm:"pk 'group_id'"`
	GroupType int32 `xorm:"'type'"`
	MenuId1   int32 `xorm:"'lesson_menu_id_1'"`
	MenuId2   int32 `xorm:"'lesson_menu_id_2'"`
}

type LessonSkillDropMemberChance struct {
	Position int32 `xorm:"pk 'member_position_id'"`
	Weight   int32 `xorm:"'member_weight'"`
}

type LessonItemDropAmount struct {
	LessonType int32 `xorm:"pk 'drop_type'"`
	Amount     int32 `xorm:"'amount'"`
	Weight     int32 `xorm:"'weight'"`
}

func init() {
	loadTable("s_lesson_skill_drop_group", LessonSkillDropGroup{})
	loadTable("s_lesson_skill_drop", LessonSkillDrop{})
	loadTable("s_lesson_skill_member_chance", LessonSkillDropMemberChance{})
	loadTable("s_lesson_item_drop_amounts", LessonItemDropAmount{})
}
