package client

import "elichika/internal/generic"

type MemberLovePanelList struct {
	MemberId               int32                `xorm:"pk" json:"member_id"`
	MemberLovePanelCellIds generic.Array[int32] `xorm:"-" json:"member_love_panel_cell_ids"`
}

type MemberLovePanel struct {
	MemberId int32 `xorm:"pk" json:"member_id"`
	PanelId  int32 `xorm:"pk" json:"panel_id"`
	Status   uint8 `xorm:"status default(0)" json:"status"`

	IsLastPanel bool `xorm:"-"`
}

func (panel *MemberLovePanel) Unlocked(index int) bool {
	mask := uint8(1 << index)
	return panel.Status&mask == mask
}

func (panel *MemberLovePanel) NoneUnlocked() bool {
	return panel.Status == 0
}

func (panel *MemberLovePanel) AllUnlocked() bool {
	return panel.Status == uint8(0x1F) // 0b00011111
}
