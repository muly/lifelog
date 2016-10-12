package model

import (
	"time"
)

type Activity struct {
	Name        string
	GoalID      string
	CreatedDate time.Time
	ModifiedOn  time.Time `json:"ModifiedOn,omitempty"`
}

/*
func ActivityPost(c context.Context,) {
}
func ActivityGet(c context.Context,) {
}
func ActivityDelete(c context.Context,) {
}
*/
