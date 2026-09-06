package reviews

type ReplyStatus int

const (
	ReplyStatusPending     ReplyStatus = 1
	ReplyStatusApproved    ReplyStatus = 5
	ReplyStatusNotApproved ReplyStatus = 8
)

const (
	ReplyPending     = ReplyStatusPending
	ReplyApproved    = ReplyStatusApproved
	ReplyNotApproved = ReplyStatusNotApproved
)

func (s ReplyStatus) String() string {
	switch s {
	case ReplyStatusPending:
		return "Pending"
	case ReplyStatusApproved:
		return "Approved"
	case ReplyStatusNotApproved:
		return "NotApproved"
	default:
		return "Unknown"
	}
}

func (s ReplyStatus) IsValid() bool {
	switch s {
	case ReplyStatusPending, ReplyStatusApproved, ReplyStatusNotApproved:
		return true
	default:
		return false
	}
}

func ParseReplyStatus(v int) (ReplyStatus, bool) {
	s := ReplyStatus(v)
	return s, s.IsValid()
}
