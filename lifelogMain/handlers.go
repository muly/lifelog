package ActivityLoggerMain

import (
	"appengine"
	//	"appengine/datastore"
	"appengine/user"
	//"errors"
	"fmt"
	"helpers"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//[TODO: need to document this function]
func handleRoot(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	var f []helpers.Filter // filter slice to store the number filters
	var OrderBy string

	// retrieve the Started activities
	f = []helpers.Filter{{"Status=", helpers.ActivityStatusStarted}}
	OrderBy = "-TimeStamp"
	actiLogS := helpers.GetActivity(c, f, OrderBy)
	//............................................

	// retrieve the Paused activities
	// since the datastore query doesn't support filter on multiple values (like an IN operator or OR operator in SQL), we are doing it in 2 passes and storing the results in the same variable (i.e merging the result sets).
	f = []helpers.Filter{{"Status=", helpers.ActivityStatusPaused}}
	OrderBy = "-TimeStamp"
	actiLogP := helpers.GetActivity(c, f, OrderBy)

	// merge the started and paused activities
	actiLog := append(actiLogS, actiLogP...)

	actiLogIcn := helpers.AddIconsToActivityLog(actiLog)

	t := template.Must(template.ParseFiles(
		"html/home.html",
		"html/_ActivityList.html",
		"html/_SvgButtons.html",
		"html/_mdl.html",
		"html/_footer.html",
		"html/_header.html",
	))

	//[TODO: need to merge the actiLog with the activity icon information before passing to template]

	// prepare the final data structure to pass to templates: add the user name to the activities list.
	a := helpers.HomePgData{user.Current(c).String(), len(actiLog), actiLogIcn}

	// execute the template while passing the required data to be rendered.
	if err := t.Execute(w, a); err != nil {
		panic(err)
	}
}

//[TODO: need to document this function]
func handleHistory(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	actiLog := helpers.GetActivity(c, nil, "-TimeStamp")

	t := template.Must(template.ParseFiles(
		"html/history.html",
		"html/_ActivityList.html",
		"html/_SvgButtons.html",
		"html/_mdl.html",
		"html/_footer.html",
		"html/_header.html",
	))

	actiLogIcn := helpers.AddIconsToActivityLog(actiLog)
	a := helpers.HomePgData{user.Current(c).String(), len(actiLog), actiLogIcn}

	if err := t.Execute(w, a); err != nil {
		panic(err)
	}
}

func handleActivitySearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	c := appengine.NewContext(r)

	var t *template.Template
	var actiLog []helpers.ActivityLog //TODO: need to rename actiLog to actiLogs as it can hold more than one

	if r.Method == "GET" {

		actiLog = helpers.GetRecomm(c)

		t = template.Must(template.ParseFiles(
			"html/SearchActivity.html",
			"html/_ActivityList.html",
			"html/_SvgButtons.html",
			"html/_header.html",
			"html/_mdl.html",
			"html/activityForms.html",
		))

	} else if r.Method == "POST" {

		a, _ := helpers.FullTextSearchActivity(c, r.FormValue("activity"), w)
		actiLog = make([]helpers.ActivityLog, len(a))
		for i := range a { //TODO: looping and retrieving for each activity as I'm not sure of direct way as of now
			actiLog[i].ActivityName = a[i].ActivityName
		}
		t = template.Must(template.ParseFiles(
			"html/searchResults.html",
			"html/_ActivityList.html",
			"html/_SvgButtons.html",
			"html/_mdl.html",
			"html/_footer.html",
			"html/_header.html",
			"html/activityForms.html",
		))

		//fmt.Fprintln(w, "Search: ", a)

	}

	actiLogIcn := helpers.AddIconsToActivityLog(actiLog)
	a := helpers.HomePgData{user.Current(c).String(), len(actiLog), actiLogIcn}
	if err := t.Execute(w, a); err != nil {
		panic(err)
	}

}

func handleActivityAddByForm(w http.ResponseWriter, r *http.Request) {

	ActivityName := r.FormValue("activity")
	SubTask := r.FormValue("subtask")

	handleActivityAdd(w, r, ActivityName, SubTask)

}

//[TODO: need to document this function]
func handleActivityAddByURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	//[TODO: query param approach is not the effecient way to handle, as the parameters values in the url are visible to anyone,a nd it could pose security issues. so, need to explore the way in which we can pass the values as part of the HTTP header/body instead of URL]
	i := strings.Index(r.RequestURI, "?")       // since url.ParseQuery is not able to retrieve the first key (of the query string) correctly, find the position of ? in the url and
	qs := r.RequestURI[i+1 : len(r.RequestURI)] //  substring it and then

	m, _ := url.ParseQuery(qs) // parse it

	ActivityName := m["ActivityName"][0]
	handleActivityAdd(w, r, ActivityName, "")

}

func handleActivityAdd(w http.ResponseWriter, r *http.Request, ActivityName string, SubTask string) {
	c := appengine.NewContext(r)
	a := helpers.ActivityLog{
		ActivityName: ActivityName,
		SubTask:      SubTask,
		TimeStamp:    time.Now(),
		StartTime:    time.Now(),
		Status:       helpers.ActivityStatusStarted,
		UserId:       user.Current(c).String(),
	}

	helpers.InsertActivity(c, a)
	helpers.AddActivityToIndex(c, a, w)

	http.Redirect(w, r, "/", http.StatusFound)
}

// handleActivityUpdate() handles the logic for "/activity/update" route.
func handleActivityUpdate(w http.ResponseWriter, r *http.Request) {
	//[TODO: query param approach is not the effecient way to handle, as the parameters values in the url are visible to anyone,a nd it could pose security issues. so, need to explore the way in which we can pass the values as part of the HTTP header/body instead of URL]
	i := strings.Index(r.RequestURI, "?")       // since url.ParseQuery is not able to retrieve the first key (of the query string) correctly, find the position of ? in the url and
	qs := r.RequestURI[i+1 : len(r.RequestURI)] //  substring it and then

	m, _ := url.ParseQuery(qs) // parse it
	c := appengine.NewContext(r)

	err := helpers.UpdateActivity(c, m["ActivityName"][0], m["StartTime"][0], m["Status"][0], m["NewStatus"][0])

	if err != nil {
		http.Error(w, "Error while changing the status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
