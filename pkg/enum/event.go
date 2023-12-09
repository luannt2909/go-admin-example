package enum

type ReminderEvent int

const (
	RECreate ReminderEvent = iota + 1
	REUpdate
	REDelete
)

type UserEvent int

const (
	UECreate UserEvent = iota + 1
	UEUpdate
	UEDelete
)
