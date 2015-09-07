package ctrl

import (
	"appengine"
	"appengine/user"
	"fmt"
	"modl"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HandleActivityAddByForm(w http.ResponseWriter, r *http.Request) {

	ActivityName := r.FormValue("activity")
	SubTask := r.FormValue("subtask")

	activityAdd(w, r, ActivityName, SubTask)

}

//[TODO: need to document this function]
func HandleActivityAddByURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	//[TODO: query param approach is not the effecient way to handle, as the parameters values in the url are visible to anyone,a nd it could pose security issues. so, need to explore the way in which we can pass the values as part of the HTTP header/body instead of URL]
	i := strings.Index(r.RequestURI, "?")       // since url.ParseQuery is not able to retrieve the first key (of the query string) correctly, find the position of ? in the url and
	qs := r.RequestURI[i+1 : len(r.RequestURI)] //  substring it and then

	m, _ := url.ParseQuery(qs) // parse it

	ActivityName := m["ActivityName"][0]
	activityAdd(w, r, ActivityName, "")

}

// handleActivityUpdate() handles the logic for "/activity/update" route.
func HandleActivityUpdate(w http.ResponseWriter, r *http.Request) {
	//[TODO: query param approach is not the effecient way to handle, as the parameters values in the url are visible to anyone,a nd it could pose security issues. so, need to explore the way in which we can pass the values as part of the HTTP header/body instead of URL]
	i := strings.Index(r.RequestURI, "?")       // since url.ParseQuery is not able to retrieve the first key (of the query string) correctly, find the position of ? in the url and
	qs := r.RequestURI[i+1 : len(r.RequestURI)] //  substring it and then

	m, _ := url.ParseQuery(qs) // parse it
	c := appengine.NewContext(r)

	err := modl.ActivityUpdate(c, m["ActivityName"][0], m["SubTask"][0], m["StartTime"][0], m["Status"][0], m["NewStatus"][0])

	if err != nil {
		http.Error(w, "Error while changing the status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

//TODO: this function is not actually a route handler,so the name has to be changed to reflect the meaning correctly
func activityAdd(w http.ResponseWriter, r *http.Request, ActivityName string, SubTask string) {
	c := appengine.NewContext(r)
	a := modl.ActivityLog{
		ActivityName: ActivityName,
		SubTask:      SubTask,
		TimeStamp:    time.Now(),
		StartTime:    time.Now(),
		Status:       modl.ActivityStatusStarted,
		UserId:       user.Current(c).String(),
	}

	modl.ActivityInsert(c, a)
	modl.ActivityAddToIndex(c, a, w)

	http.Redirect(w, r, "/", http.StatusFound)
}
