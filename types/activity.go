package types

import (
	"time"
)

//TODO: need to add json tags for column names and to ignore blank fields
type (
	ActivityLog struct {
		Name        string
		Notes       string `json:"Notes,omitempty"`
		StartTime   time.Time
		EndTime     time.Time
		CreatedDate time.Time
		ModifiedOn  time.Time `json:"ModifiedOn,omitempty"`
	}
)
