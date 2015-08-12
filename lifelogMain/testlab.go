package ActivityLoggerMain

import (
	"appengine"
	"appengine/datastore"
	//"appengine/user"
	//"errors"
	"fmt"
	"helpers"
	//"html/template"
	"net/http"
	//"net/url"
	//"strings"
	//"time"
)

type ActivityNameOnly struct {
	ActivityName string
}

func handleTestLab(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	var recSet []ActivityNameOnly
	parentKey := helpers.GetActivityTableKeyByUser(c)

	activeRecs := datastore.NewQuery("activityRecord").Ancestor(parentKey)

	//activeRecs = activeRecs.Filter("Status=", helpers.ActivityStatusStarted)
	activeRecs = activeRecs.Project("ActivityName") //.Distinct() // pulling distinct activity names only

	activeRecs = activeRecs.Filter("ActivityName =", "search me")

	//	activeRecs = activeRecs.Order(orderBy)
	t := activeRecs.Run(c)

	for {
		//var recSet activityRecord
		_, err := t.Next(&recSet)
		if err == datastore.Done {
			break
		}
		if err != nil {
			c.Errorf("Running query: %v", err)
			break
		}
		fmt.Fprintln(w, "..")
		fmt.Fprintln(w, recSet)
	}

	//	fmt.Fprintln(w, recSet)
	// recSet

}

/*
//Note: had to create a new function to retrieve distinct activity names based on search criteria.
//		this did not fit in the GetActivity function because of .Project() returns only single column result set
//		where as the non projection query returns all fields, hence mismatch and resulting in issue. hence seperated them.
//Note: this code is NOT working, not returning records and so I need to revisit.
func GetActivityNames(c appengine.Context, filters []Filter, orderBy string) []string { //[TODO: need to return error]
	parentKey := GetActivityTableKeyByUser(c)
	recSet := []string{}
	fmt.Println("GetActivityNames")
	activeRecs := datastore.NewQuery("activityRecord").Ancestor(parentKey)

	for _, f := range filters {
		activeRecs = activeRecs.Filter(f.Left, f.Right)
	}

	activeRecs = activeRecs.Project("ActivityName").Distinct() // pulling distinct activity names only

	activeRecs = activeRecs.Order(orderBy)
	activeRecs.GetAll(c, &recSet)

	return recSet

}*/
