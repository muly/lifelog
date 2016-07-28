package types

import (
	"time"
)

//TODO: need to add json tags for column names and to ignore blank fields
type (
	Activity struct {
		Name        string
		GoalID      string
		createdDate time.Time
	}

	ActivityLog struct {
		Name        string //'json:Name'
		StartTime   time.Time
		EndTime     time.Time
		createdDate time.Time
	}
)

func (a *ActivityLog) SetDefaults() {
	a.createdDate = time.Now()
}
