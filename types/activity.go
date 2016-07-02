package types

import (
	"time"
)

type (
	Activity struct {
		Name   string
		GoalID string
	}

	Goal struct {
		Name  string
		Notes string
	}

	ActivityLog struct {
		ActivityName string
		StartTime    time.Time
		EndTime      time.Time
	}
)
