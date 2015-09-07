package ctrl

import (
	"appengine"
	"appengine/user"
	"html/template"
	"modl"
	"net/http"
)

//[TODO: need to document this function]
func HandleHistory(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	actiLog := modl.ActivityGet(c, nil, "-TimeStamp")

	t := template.Must(template.ParseFiles(
		"html/history.html",
		"html/_ActivityList.html",
		"html/_SvgButtons.html",
		"html/_mdl.html",
		"html/_footer.html",
		"html/_header.html",
	))

	actiLogIcn := modl.AddIconsToActivityLog(actiLog)
	a := modl.HomePgData{user.Current(c).String(), len(actiLog), actiLogIcn}

	if err := t.Execute(w, a); err != nil {
		panic(err)
	}
}
