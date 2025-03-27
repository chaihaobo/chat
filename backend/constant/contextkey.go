package constant

const (
	ContextKeyTrx ContextKey = iota + 1
	ContextKeyUserID
	ContextKeyUserName
	ContextKeyUserAvatar
)

type ContextKey int
