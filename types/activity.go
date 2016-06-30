package types

import (
	"time"
)

type (
	Activity struct {
		Name      string
		GoalID    string
		StartTime string
		EndTime   string
	}

	Goal struct {
		Name  string
		Notes string
	}

	AcctivityLog struct {
		ActivityName string
		StartTime    time.Time
		EndTime      time.Time
	}
)
