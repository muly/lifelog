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
		Name       string
		Notes      string
		CreatedOn  time.Time
		ModifiedOn time.Time
	}
	Goals []Goal
)

func (goal *Goal) Put(c context.Context) error {

	// generate the key
	key := datastore.NewKey(c, "Goal", util.StringKey(goal.Name), 0, nil)

	// put the record into the database and capture the key
	key, err := datastore.Put(c, key, goal)
	if err != nil {
		return err
	}

	// read from database into the same variable
	if err = datastore.Get(c, key, goal); err != nil {
		return err
	}

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

func (goals *Goals) Get(c context.Context) (err error) {
	_, err = datastore.NewQuery("Goal").GetAll(c, goals)

	return
}
