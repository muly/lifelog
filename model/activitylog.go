package model

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
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

func (a *ActivityLog) SetDefaults() {
	a.CreatedDate = time.Now()
}

func ActivityLogPost(c context.Context, activityLog *ActivityLog) error {
	// generate new key
	key := datastore.NewIncompleteKey(c, "ActivityLog", nil)
	//store in database
	_, err := datastore.Put(c, key, activityLog)

	return err
}
