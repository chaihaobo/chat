package querytypes

import (
	"gorm.io/gorm/clause"
)

type RecentlyMessageQuery struct {
	Pagination
	FriendUserID *uint64
	UserID       *uint64
}

func (q RecentlyMessageQuery) ToClauses() []clause.Expression {
	var conditions []clause.Expression
	if q.FriendUserID != nil {
		conditions = append(conditions, clause.Or(clause.Eq{Column: "from", Value: q.FriendUserID}))
		conditions = append(conditions, clause.Or(clause.Eq{Column: "to", Value: q.FriendUserID}))
	}
	if q.UserID != nil {
		conditions = append(conditions, clause.Or(clause.Eq{Column: "from", Value: q.UserID}))
		conditions = append(conditions, clause.Or(clause.Eq{Column: "to", Value: q.UserID}))
	}
	return conditions
}
