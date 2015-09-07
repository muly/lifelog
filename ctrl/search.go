package ctrl

import (
	"appengine"
	"appengine/user"
	"fmt"
	"html/template"
	"modl"
	"net/http"
)

func HandleActivitySearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	c := appengine.NewContext(r)

	var t *template.Template
	var actiLog []modl.ActivityLog //TODO: need to rename actiLog to actiLogs as it can hold more than one

	if r.Method == "GET" {

		actiLog = modl.GetRecomm(c)

		t = template.Must(template.ParseFiles(
			"html/SearchActivity.html",
			"html/_ActivityList.html",
			"html/_SvgButtons.html",
			"html/_header.html",
			"html/_mdl.html",
			"html/_activityForms.html",
		))

	} else if r.Method == "POST" {

		a, _ := modl.FullTextSearchActivity(c, r.FormValue("activity"), w)
		actiLog = make([]modl.ActivityLog, len(a))
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
			"html/_activityForms.html",
		))

		//fmt.Fprintln(w, "Search: ", a)

	}

	actiLogIcn := modl.AddIconsToActivityLog(actiLog)
	a := modl.HomePgData{user.Current(c).String(), len(actiLog), actiLogIcn}
	if err := t.Execute(w, a); err != nil {
		panic(err)
	}

}
