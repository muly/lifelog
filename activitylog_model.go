package lifelog

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"

	"time"
	//"github.com/muly/lifelog/util"
	//"types"
)

//TODO: need to add json tags for column names and to ignore blank fields
type (
	ActivityLog struct {
		Name       string
		Notes      string `json:"Notes,omitempty"`
		StartTime  time.Time
		EndTime    time.Time
		CreatedOn  time.Time `json:"CreatedOn,omitempty"`
		ModifiedOn time.Time `json:"ModifiedOn,omitempty"`
	}
	ActivityLogs []ActivityLog
)

// Get
func (al *ActivityLog) Get(c context.Context) (err error) {
	key := datastore.NewKey(c, "ActivityLog", StringKey(al.Name), 0, nil)

	err = datastore.Get(c, key, al)
	if err != nil && err.Error() == "datastore: no such entity" {
		err = ErrorNoMatch
	}

	return
}

// Put (same for Post)
func (al *ActivityLog) Put(c context.Context) error {
	key := datastore.NewKey(c, "ActivityLog", StringKey(al.Name), 0, nil)

	// put the record into the database and capture the key
	key, err := datastore.Put(c, key, al)
	if err != nil {
		return err
	}

	// read from database into the same variable
	if err = datastore.Get(c, key, al); err != nil {
		return err
	}

	return err
}

// Get(s)

func (als *ActivityLogs) Get(c context.Context, filter ActivityLog, offset int, limit int) (err error) {
	q := datastore.NewQuery("ActivityLog")

	if filter.Name != "" {
		q = q.Filter("Name =", filter.Name)
	}

	if filter.Notes != "" {
		q = q.Filter("Notes =", filter.Notes)
	}

	q = q.Offset(offset).Limit(limit).Order("Name")

	_, err = q.GetAll(c, als)
	if err != nil {
		return
	}

	return
}

// Delete
func (al *ActivityLog) Delete(c context.Context) (err error) {
	// TODO: need to check for existance before deleting. if NOT exists, then throw ErrorNoMatch error (err = ErrorNoMatch)
	if err = al.Get(c); err == ErrorNoMatch {
		return
	}

	key := datastore.NewKey(c, "ActivityLog", StringKey(al.Name), 0, nil)

	err = datastore.Delete(c, key)

	return
}

func (a *ActivityLog) SetDefaults() {
	a.CreatedOn = time.Now()
}
