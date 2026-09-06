package comments

type CommentStatus int

const (
	CommentStatusPending     CommentStatus = 1
	CommentStatusApproved    CommentStatus = 5
	CommentStatusNotApproved CommentStatus = 8
)

// Aliases curtos compatíveis com a spec original
// (Pending / Approved / NotApproved). Em Go não existe
// sintaxe Tipo.Membro como no C# (CommentStatus.Approved);
// o uso correto é pacote.Constante, ex: comments.Approved.
const (
	Pending     = CommentStatusPending
	Approved    = CommentStatusApproved
	NotApproved = CommentStatusNotApproved
)

// String retorna a representação textual do status.
func (s CommentStatus) String() string {
	switch s {
	case CommentStatusPending:
		return "Pending"
	case CommentStatusApproved:
		return "Approved"
	case CommentStatusNotApproved:
		return "NotApproved"
	default:
		return "Unknown"
	}
}

// IsValid indica se o valor é um status conhecido.
func (s CommentStatus) IsValid() bool {
	switch s {
	case CommentStatusPending, CommentStatusApproved, CommentStatusNotApproved:
		return true
	default:
		return false
	}
}

// ParseCommentStatus converte int para CommentStatus validado.
func ParseCommentStatus(v int) (CommentStatus, bool) {
	s := CommentStatus(v)
	return s, s.IsValid()
}
