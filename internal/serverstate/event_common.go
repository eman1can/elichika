package serverstate

type EventSchedule struct {
	EventId   int32 `xorm:"pk 'event_id'"`
	StartAt   int64 `xorm:"'start_at'"`
	EndAt     int64 `xorm:"'end_at'"`
	ResultAt  int64 `xorm:"'result_at'"`
	ExpiredAt int64 `xorm:"'expired_at'"`
}

func init() {
	addTable("s_event_schedule", EventSchedule{}, nil)
}
