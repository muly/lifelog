package logrMain

import (
	"appengine"
	"appengine/datastore"
	"appengine/search"
	//"appengine/user"
	//"errors"
	"fmt"
	"html/template"
	"modl"
	"net/http"
	"strconv"
	//"net/url"
	//"strings"
	"time"
)

type ActivityNameOnly struct {
	ActivityName string
}

func handleLabIcon(w http.ResponseWriter, r *http.Request) {

	a := modl.ActivityGetIconsData()

	t := template.Must(template.ParseFiles(
		"html/test.html",
		"html/_ActivityList.html",
	))

	if err := t.Execute(w, a); err != nil {
		panic(err)
	}

	//fmt.Fprintln(w, UsefulMaterialIcons)
}

func handleLabTest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	var recSet []ActivityNameOnly
	parentKey := modl.GetActivityTableKeyByUser(c)

	activeRecs := datastore.NewQuery("activityRecord").Ancestor(parentKey)

	//activeRecs = activeRecs.Filter("Status=", modl.ActivityStatusStarted)
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

func handleActivityIndexLab(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	s := strconv.Itoa(time.Now().Hour()) + strconv.Itoa(time.Now().Minute())
	s = s + " Google"

	i, _ := search.Open("ActivityLog")
	/*
				a := ActivityNameOnly{s}
				id, err := i.Put(c, "", &a)
				if err != nil {
					panic(err)
				}
				fmt.Fprintln(w, "Put: ", id)

				dst := ActivityNameOnly{}
				i.Get(c, id, &dst)
				fmt.Fprintln(w, "Get : ", dst)




			for t := i.Search(c, `ActivityName:index`, nil); ; {
				fmt.Fprintln(w, "Search: ", t.Count())

				dst := ActivityNameOnly{}
				_, err := t.Next(&dst)
				if err == search.Done {
					break
				}
				fmt.Fprintln(w, "Search: ", dst)
			}


		src := "index"
		t := i.Search(c, `ActivityName:`+src, nil)

		dst := make([]ActivityNameOnly, 10)
		for i := 0; ; i++ { //TODO: is there a better way to handle the loop for search output? need to research
			_, err := t.Next(&dst[i])
			if err == search.Done {
				break
			}
		}
		fmt.Fprintln(w, "Search f: ", dst)
	*/
	dst, _ := modl.FullTextSearchActivity(c, "index", w)
	fmt.Fprintln(w, "Search f: ", dst)

	for t := i.List(c, nil); ; {
		dst := ActivityNameOnly{}
		_, err := t.Next(&dst)
		if err == search.Done {
			break
		}
		fmt.Fprintln(w, "List: ", dst)
	}

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
