package notifications

type NotificationSeverity byte

const (
	NotificationSeverityInfo    NotificationSeverity = 0
	NotificationSeveritySuccess NotificationSeverity = 1
	NotificationSeverityWarn    NotificationSeverity = 2
	NotificationSeverityError   NotificationSeverity = 3
	NotificationSeverityFatal   NotificationSeverity = 4
)

func (s NotificationSeverity) String() string {
	switch s {
	case NotificationSeverityInfo:
		return "Info"
	case NotificationSeveritySuccess:
		return "Success"
	case NotificationSeverityWarn:
		return "Warn"
	case NotificationSeverityError:
		return "Error"
	case NotificationSeverityFatal:
		return "Fatal"
	default:
		return "Unknown"
	}
}

func (s NotificationSeverity) IsValid() bool {
	switch s {
	case NotificationSeverityInfo, NotificationSeveritySuccess, NotificationSeverityWarn,
		NotificationSeverityError, NotificationSeverityFatal:
		return true
	default:
		return false
	}
}

func ParseNotificationSeverity(v byte) (NotificationSeverity, bool) {
	s := NotificationSeverity(v)
	return s, s.IsValid()
}
