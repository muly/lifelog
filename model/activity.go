package model

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"

	"time"

	"types"
	"util"
)

//TODO: need to add json tags for column names and to ignore blank fields
type (
	Activity struct {
		Name       string
		GoalID     string
		CreatedOn  time.Time `json:"CreatedOn,omitempty"`
		ModifiedOn time.Time `json:"ModifiedOn,omitempty"`
	}
	Activities []Activity
)

// Get
func (act *Activity) Get(c context.Context) (err error) {
	key := datastore.NewKey(c, "Activity", util.StringKey(act.Name), 0, nil)

	err = datastore.Get(c, key, act)
	if err != nil && err.Error() == "datastore: no such entity" {
		err = types.ErrorNoMatch
	}

	return
}

// Put (same for Post)
func (act *Activity) Put(c context.Context) (err error) {
	key := datastore.NewKey(c, "Activity", util.StringKey(act.Name), 0, nil)

	// put the record into the database and capture the key

	key, err = datastore.Put(c, key, act)
	if err != nil {
		return err
	}

	// read from database into the same variable
	if err = datastore.Get(c, key, act); err != nil {
		return err
	}

	return err
}

// Get(s)

func (acts *Activities) Get(c context.Context, filter Activity, offset int, limit int) (err error) {
	q := datastore.NewQuery("Activity")

	if filter.Name != "" {
		q = q.Filter("Name =", filter.Name)
	}

	if filter.GoalID != "" {
		q = q.Filter("GoalID =", filter.GoalID)
	}

	q = q.Offset(offset).Limit(limit).Order("Name")

	_, err = q.GetAll(c, acts)
	if err != nil {
		return
	}

	return
}

// Delete
func (act *Activity) Delete(c context.Context) (err error) {
	// TODO: need to check for existance before deleting. if NOT exists, then throw ErrorNoMatch error (err = ErrorNoMatch)
	if err = act.Get(c); err == types.ErrorNoMatch {
		return
	}

	key := datastore.NewKey(c, "Activity", util.StringKey(act.Name), 0, nil)

	err = datastore.Delete(c, key)

	return
}

func (a *Activity) SetDefaults() {
	a.CreatedOn = time.Now()
}
