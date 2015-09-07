package ctrl

import (
	"appengine"
	"appengine/user"
	"html/template"
	"modl"
	"net/http"
)

//[TODO: need to document this function]
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	var f []modl.Filter // filter slice to store the number filters
	var OrderBy string

	// retrieve the Started activities
	f = []modl.Filter{{"Status=", modl.ActivityStatusStarted}}
	OrderBy = "-TimeStamp"
	actiLogS := modl.ActivityGet(c, f, OrderBy)
	//............................................

	// retrieve the Paused activities
	// since the datastore query doesn't support filter on multiple values (like an IN operator or OR operator in SQL), we are doing it in 2 passes and storing the results in the same variable (i.e merging the result sets).
	f = []modl.Filter{{"Status=", modl.ActivityStatusPaused}}
	OrderBy = "-TimeStamp"
	actiLogP := modl.ActivityGet(c, f, OrderBy)

	// merge the started and paused activities
	actiLog := append(actiLogS, actiLogP...)

	actiLogIcn := modl.AddIconsToActivityLog(actiLog)

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
	a := modl.HomePgData{user.Current(c).String(), len(actiLog), actiLogIcn}

	// execute the template while passing the required data to be rendered.
	if err := t.Execute(w, a); err != nil {
		panic(err)
	}
}
