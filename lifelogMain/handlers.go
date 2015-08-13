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
	var dst []helpers.ActivityLog // dst to store the query results
	var f []helpers.Filter        // filter slice to store the number filters
	var OrderBy string

	// retrieve the Started activities
	f = []helpers.Filter{{"Status=", helpers.ActivityStatusStarted}}
	OrderBy = "-TimeStamp"
	dst = helpers.GetActivity(c, f, OrderBy)
	//............................................

	// retrieve the Paused activities
	// since the datastore query doesn't support filter on multiple values (like an IN operator or OR operator in SQL), we are doing it in 2 passes and storing the results in the same variable (i.e merging the result sets).
	f = []helpers.Filter{{"Status=", helpers.ActivityStatusPaused}}
	OrderBy = "-TimeStamp"
	dst2 := helpers.GetActivity(c, f, OrderBy)

	// merge the started and paused activities
	dst = append(dst, dst2...)

	t := template.Must(template.ParseFiles(
		"html/home.html",
		"html/_ActivityList.html",
		"html/_SvgButtons.html",
		"html/_mdl.html",
		"html/_footer.html",
		"html/_header.html",
	))

	//[TODO: need to merge the dst with the activity icon information before passing to template]

	// prepare the final data structure to pass to templates: add the user name to the activities list.
	a := helpers.HomePgData{user.Current(c).String(), len(dst), dst}

	// execute the template while passing the required data to be rendered.
	if err := t.Execute(w, a); err != nil {
		panic(err)
	}
}

//[TODO: need to document this function]
func handleHistory(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	var dst []helpers.ActivityLog

	dst = helpers.GetActivity(c, nil, "-TimeStamp")

	t := template.Must(template.ParseFiles(
		"html/history.html",
		"html/_ActivityList.html",
		"html/_SvgButtons.html",
		"html/_mdl.html",
		"html/_footer.html",
		"html/_header.html",
	))
	a := helpers.HomePgData{user.Current(c).String(), len(dst), dst}

	if err := t.Execute(w, a); err != nil {
		panic(err)
	}

}
func handleActivitySearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	c := appengine.NewContext(r)

	if r.Method == "GET" {

		var dst []helpers.ActivityLog
		dst = helpers.GetRecomm(c)

		t := template.Must(template.ParseFiles(
			"html/SearchActivity.html",
			"html/_ActivityList.html",
			"html/_SvgButtons.html",
			"html/_header.html",
			"html/_mdl.html",
		))
		a := helpers.HomePgData{user.Current(c).String(), len(dst), dst}
		if err := t.Execute(w, a); err != nil {
			panic(err)
		}
	} else if r.Method == "POST" {

		var dst []helpers.ActivityLog
		var f []helpers.Filter // filter slice to store the number filters
		var OrderBy string

		// retrieve the Started activities
		f = []helpers.Filter{{"ActivityName=", r.FormValue("activity")}}

		OrderBy = "-TimeStamp"
		dst = helpers.GetActivity(c, f, OrderBy)

		t := template.Must(template.ParseFiles(
			"html/searchResults.html",
			"html/_ActivityList.html",
			"html/_SvgButtons.html",
			"html/_mdl.html",
			"html/_footer.html",
			"html/_header.html",
		))
		// prepare the final data structure to pass to templates: add the user name to the activities list.
		a := helpers.HomePgData{user.Current(c).String(), len(dst), dst}

		if err := t.Execute(w, a); err != nil {
			panic(err)
		}

	}
}

func handleActivityAddByForm(w http.ResponseWriter, r *http.Request) {

	ActivityName := r.FormValue("activity")

	handleActivityAdd(w, r, ActivityName)

}

//[TODO: need to document this function]
func handleActivityAddByURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	//[TODO: query param approach is not the effecient way to handle, as the parameters values in the url are visible to anyone,a nd it could pose security issues. so, need to explore the way in which we can pass the values as part of the HTTP header/body instead of URL]
	i := strings.Index(r.RequestURI, "?")       // since url.ParseQuery is not able to retrieve the first key (of the query string) correctly, find the position of ? in the url and
	qs := r.RequestURI[i+1 : len(r.RequestURI)] //  substring it and then

	m, _ := url.ParseQuery(qs) // parse it

	ActivityName := m["ActivityName"][0]
	handleActivityAdd(w, r, ActivityName)

}

func handleActivityAdd(w http.ResponseWriter, r *http.Request, ActivityName string) {
	c := appengine.NewContext(r)
	a := helpers.ActivityLog{
		ActivityName: ActivityName,
		TimeStamp:    time.Now(),
		StartTime:    time.Now(),
		Status:       helpers.ActivityStatusStarted,
		UserId:       user.Current(c).String(),
	}

	helpers.InsertActivity(c, a)

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
