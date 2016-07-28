package model

import (
	"types"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

func ActivityLogPost(c context.Context, activityLog *types.ActivityLog) error {
	// generate new key
	key := datastore.NewIncompleteKey(c, "ActivityLog", nil)
	//store in database
	_, err := datastore.Put(c, key, activityLog)

	return err
}

/*
func ActivityPost(c context.Context,) {
}
func ActivityGet(c context.Context,) {
}
func ActivityDelete(c context.Context,) {
}
*/
