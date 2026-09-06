package reviews

type ReviewStatus int

const (
	ReviewStatusPending     ReviewStatus = 1
	ReviewStatusApproved    ReviewStatus = 5
	ReviewStatusNotApproved ReviewStatus = 8
)

const (
	Pending     = ReviewStatusPending
	Approved    = ReviewStatusApproved
	NotApproved = ReviewStatusNotApproved
)

func (s ReviewStatus) String() string {
	switch s {
	case ReviewStatusPending:
		return "Pending"
	case ReviewStatusApproved:
		return "Approved"
	case ReviewStatusNotApproved:
		return "NotApproved"
	default:
		return "Unknown"
	}
}

func (s ReviewStatus) IsValid() bool {
	switch s {
	case ReviewStatusPending, ReviewStatusApproved, ReviewStatusNotApproved:
		return true
	default:
		return false
	}
}

func ParseReviewStatus(v int) (ReviewStatus, bool) {
	s := ReviewStatus(v)
	return s, s.IsValid()
}
