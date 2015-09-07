package logrMain

import (
	//"appengine"
	//"appengine/datastore"
	//"appengine/user"
	//"fmt"
	//"html/template"
	"ctrl"
	"net/http"
	//"net/url"
	//"strings"
	//"time"
)

func init() {
	// note: order doesn't look to be important, atleast so far
	http.HandleFunc("/history/", ctrl.HandleHistory)
	http.HandleFunc("/", ctrl.HandleRoot)
	http.HandleFunc("/activity/search", ctrl.HandleActivitySearch)

	http.HandleFunc("/activity/addbyurl", ctrl.HandleActivityAddByURL) // the /activity/ should match with what is in HTML form action ?? not really sure
	http.HandleFunc("/activity/addbyform", ctrl.HandleActivityAddByForm)
	http.HandleFunc("/activity/update", ctrl.HandleActivityUpdate)

	http.HandleFunc("/labindex", handleActivityIndexLab)
	http.HandleFunc("/labtest", handleLabTest) // not real code. just for practicing new concepts and experimenting. will be removed eventually
	http.HandleFunc("/labicon", handleLabIcon) // not real code. just for practicing new concepts and experimenting. will be removed eventually
}
