package model

import (
	"time"
	"types"
	"util"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

//TODO: need to add json tags for column names and to ignore blank fields
type (
	Goal struct {
		Name        string
		Notes       string
		createdDate time.Time
	}
	Goals []Goal
)

func (goal *Goal) Post(c context.Context) error {

	key := datastore.NewKey(c, "Goal", util.StringKey(goal.Name), 0, nil)

	_, err := datastore.Put(c, key, goal)

	return err
}

func (goal *Goal) Get(c context.Context) (err error) {
	key := datastore.NewKey(c, "Goal", util.StringKey(goal.Name), 0, nil)

	err = datastore.Get(c, key, goal)
	if err != nil && err.Error() == "datastore: no such entity" {
		err = types.ErrorNoMatch
	}

	return
}

func (goal *Goal) Delete(c context.Context) (err error) {
	// TODO: need to check for existance before deleting. if NOT exists, then throw ErrorNoMatch error (err = ErrorNoMatch)
	if err = goal.Get(c); err == types.ErrorNoMatch {
		return
	}

	key := datastore.NewKey(c, "Goal", util.StringKey(goal.Name), 0, nil)

	err = datastore.Delete(c, key)

	return
}

func (a *Goal) SetDefaults() {
	a.createdDate = time.Now()
}

// TODO:
func (goals *Goals) Get(c context.Context) (err error) {

	return
}
