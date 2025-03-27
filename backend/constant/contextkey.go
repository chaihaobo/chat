package constant

const (
	ContextKeyTrx ContextKey = iota + 1
	ContextKeyUserID
	ContextKeyUserName
)

type ContextKey int
