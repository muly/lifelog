package helpers // [TODO: eventually, the package name needs to be replaced with something more meaningful.]

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"

	"errors"
	"strings"
	"time"
)

//[TODO: need to document this section]
type ActivityLog struct {
	UserId       string //[TODO: not sure if I really need this here. if it can be retrieved from context, or via some other method, we should be good and secure by removing it from here]
	ActivityName string
	TimeStamp    time.Time //[TODO: need to evaluate if this field is required as it always will have the same value as that of StartTime]
	StartTime    time.Time
	EndTime      time.Time
	Status       string
}

//[TODO: need to document this section]
type HomePgData struct {
	UserName string
	Activity []ActivityLog
}

type Filter struct {
	Left  string
	Right string
}

//[TODO: need to document this section]
const ( //Note: the values of these constants will have impact in the templates (expecially "_ActivityList.html"), so beaware while changing these values
	ActivityStatusStarted = "S"
	ActivityStatusPaused  = "P"
	//ActivityStatusReStarted = "R" //"restart" is a special case of "start", where in the previous state is a brief "pause" // restart state removed to simplify the logic.
	ActivityStatusCompleted = "C"
)

func GetActivity(c appengine.Context, filters []Filter, orderBy string) []ActivityLog {
	parentKey := GetActivityTableKeyByUser(c)
	recSet := []ActivityLog{}

	activeRecs := datastore.NewQuery("activityRecord").Ancestor(parentKey)

	for _, f := range filters {
		activeRecs = activeRecs.Filter(f.Left, f.Right)
	}
	activeRecs = activeRecs.Order(orderBy)
	activeRecs.GetAll(c, &recSet)

	return recSet

}

func InsertActivity(c appengine.Context, rec ActivityLog) (err error) {
	if rec.ActivityName == "" { //server side data validations
		err = errors.New("ActivityName cannot be blank") //[TODO: this will also be handled at the client side, though we are checking at the server side.
		return
	}
	rec.ActivityName = strings.TrimSpace(rec.ActivityName)
	parentKey := GetActivityTableKeyByUser(c)
	childKey := datastore.NewIncompleteKey(c, "activityRecord", parentKey)

	_, err = datastore.Put(c, childKey, &rec)

	if err != nil {
		return
	}
	return
}

func UpdateActivity(c appengine.Context, ActivityName string, StartTime string, Status string, NewStatus string) (err error) {

	parentKey := GetActivityTableKeyByUser(c)

	q := datastore.NewQuery("activityRecord").Ancestor(parentKey)
	q = q.KeysOnly()

	t, _ := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", StartTime) // the first parameter of Parse should be in correct format that matches the way the data is stored in the datastore, otherwise, it will not retrieve correctly. [TODO: need to handle the error so that we know it in case if the parse fails because of format incompatibility]

	// considering the ActivityName, TimeStamp (user is already covered by the Context) as the keys suffecient enough to uniquely pull a record, they are applied to the query filters
	q = q.Filter("ActivityName=", ActivityName) //[TODO: need to see if there is any other way to retrieve the value instead of having to access ysing the array[0]]
	q = q.Filter("StartTime =", t)
	q = q.Limit(1) // just in case if the above keys result in more than 1 record, applying Limit(1) to set only 1 record [TODO: not sure if this makes sense]

	dst := ActivityLog{}

	k, err := q.GetAll(c, &dst)
	if err != nil {
		return
	}

	currStatus := Status
	newStatus := NewStatus

	currRec := ActivityLog{}
	newRec := ActivityLog{}

	currRec, newRec, err = HandleActivityStatusChange(c, k[0], currStatus, newStatus)
	if err != nil {
		return
	}

	_, err = datastore.Put(c, k[0], &currRec) //update the existing record. Note: we are updating even if no change is required, as it makes no harm and moreover it simplifies the code
	if err != nil {
		return
	}

	if newRec.ActivityName != "" { // if ActivityName is not blank, that means we have data in newRec that needs to be persisted, otherwise no need persist anything.
		//parentKey := GetActivityTableKeyByUser(c)
		newKey := datastore.NewIncompleteKey(c, "activityRecord", parentKey)
		_, err = datastore.Put(c, newKey, &newRec)
		if err != nil {
			return
		}
	}
	return
}

//HandleActivityStatusChange takes the current status and the new status of a task (key is passed), and will return the updated record and/or new record, based on the businesslogic/workflow. it takes the context as inout parameter inorder to identfy
func HandleActivityStatusChange(c appengine.Context, key *datastore.Key, currStatus string, newStatus string) (currRec ActivityLog, newRec ActivityLog, err error) {

	err = datastore.Get(c, key, &currRec)
	if err != nil {
		return
	}

	if (currStatus == ActivityStatusStarted && newStatus == ActivityStatusCompleted) || (currStatus == ActivityStatusStarted && newStatus == ActivityStatusPaused) { // “start” to “complete” , “Start” to “pause”:
		// prepare existing record set:
		currRec.Status = newStatus   //use the new status,
		currRec.EndTime = time.Now() //set the end time as now, rest of the fields unchanged
	} else if currStatus == ActivityStatusPaused && newStatus == ActivityStatusCompleted { // “pause” to “complete”
		// prepare existing record set:
		currRec.Status = newStatus //use the new status, rest of the fields unchanged
	} else if (currStatus == ActivityStatusCompleted && newStatus == ActivityStatusStarted) || (currStatus == ActivityStatusPaused && newStatus == ActivityStatusStarted) { //“complete” to “start”, “pause” to “start”
		// prepare new record set:
		newRec.ActivityName = currRec.ActivityName // copy the Activity Name from current record
		newRec.UserId = currRec.UserId             // copy the Userid from current record
		newRec.Status = newStatus                  // use the new status
		newRec.TimeStamp = time.Now()              // set the timestamp as now
		newRec.StartTime = newRec.TimeStamp        // set the start time as now
	} else { // for rest of all state changes, error out:
		err = errors.New("Error: unknown state change from " + currStatus + " to " + newStatus)
		return
	}

	if currStatus == ActivityStatusPaused && newStatus == ActivityStatusStarted { //additional logic for “pause” to “start”:
		// prepare existing record set:
		currRec.Status = ActivityStatusCompleted // also make the existing record status as "completed"
	}
	return
}

//[TODO: need to document this function]
func GetActivityTableKeyByUser(c appengine.Context) *datastore.Key {
	UserId := user.Current(c).String()
	return datastore.NewKey(c, "activity", UserId, 0, nil)

}

func add(a int, b int) (s int) {
	s = a + b
	return
}
