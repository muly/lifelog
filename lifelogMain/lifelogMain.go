package ActivityLoggerMain

import (
	//"appengine"
	//"appengine/datastore"
	//"appengine/user"
	//"fmt"
	//"html/template"
	"net/http"
	//"net/url"
	//"strings"
	//"time"
)

func init() {
	// note: order doesn't look to be important, atleast so far
	http.HandleFunc("/history/", handleHistory)
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/activity/search", handleActivitySearch)
	http.HandleFunc("/activity/add", handleActivityAdd) // the /activity/ should match with what is in HTML form action ?? not really sure
	http.HandleFunc("/activity/update", handleActivityUpdate)
}

/*
notes: lessons learns from errors;
1) text/template or html/template did not matter
2) HandleFunc pattern should match the form's action; the match should be exact, i.e. case sensitive as well as the / should also match. looks like that this is not true always, or may be it is true but the chrome browser cache is handling it somehow.
*/
