package notifications

type UserNotificationState int

const (
	UserNotificationStateUnread UserNotificationState = 0
	UserNotificationStateRead   UserNotificationState = 1
)

func (s UserNotificationState) String() string {
	switch s {
	case UserNotificationStateUnread:
		return "Unread"
	case UserNotificationStateRead:
		return "Read"
	default:
		return "Unknown"
	}
}

func (s UserNotificationState) IsValid() bool {
	switch s {
	case UserNotificationStateUnread, UserNotificationStateRead:
		return true
	default:
		return false
	}
}

func ParseUserNotificationState(v int) (UserNotificationState, bool) {
	s := UserNotificationState(v)
	return s, s.IsValid()
}
