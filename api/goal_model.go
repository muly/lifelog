package lifelog

import (
	"time"

	//"github.com/muly/lifelog/util"
	//"types"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type (
	Goal struct {
		Name       string
		Notes      string `json:"Notes,omitempty"`
		CreatedOn  time.Time
		ModifiedOn time.Time `json:"ModifiedOn,omitempty"`
	}
	Goals []Goal
)

// Put saves the goal record to database. In this case to Google Appengine Datastore. If already exists, the record will be overwritten.
func (goal *Goal) Put(c context.Context) error {

	// generate the key
	key := datastore.NewKey(c, "Goal", StringKey(goal.Name), 0, nil)

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

// Get retrieves the record based on the provided key.
//
func (goal *Goal) Get(c context.Context) (err error) {
	key := datastore.NewKey(c, "Goal", StringKey(goal.Name), 0, nil)

	err = datastore.Get(c, key, goal)
	if err != nil && err.Error() == "datastore: no such entity" {
		err = ErrorNoMatch
	}

	return
}

// Delete deletes the record based on the provided key.
//
func (goal *Goal) Delete(c context.Context) (err error) {
	// TODO: need to check for existance before deleting. if NOT exists, then throw ErrorNoMatch error (err = ErrorNoMatch)
	if err = goal.Get(c); err == ErrorNoMatch {
		return
	}

	key := datastore.NewKey(c, "Goal", StringKey(goal.Name), 0, nil)

	err = datastore.Delete(c, key)

	return
}

// Get gets all the goal records matching the given criteria.
// Fields on which query supported are Name, Notes
//
func (goals *Goals) Get(c context.Context, filter Goal, offset int, limit int) (err error) {
	q := datastore.NewQuery("Goal")

	if filter.Name != "" {
		q = q.Filter("Name =", filter.Name)
	}

	if filter.Notes != "" {
		q = q.Filter("Notes =", filter.Notes)
	}

	q = q.Offset(offset).Limit(limit).Order("Name")

	_, err = q.GetAll(c, goals)
	if err != nil {
		return
	}

	return
}
