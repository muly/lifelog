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
	http.HandleFunc("/activity/addbyurl", handleActivityAddByURL) // the /activity/ should match with what is in HTML form action ?? not really sure
	http.HandleFunc("/activity/addbyform", handleActivityAddByForm)
	http.HandleFunc("/activity/update", handleActivityUpdate)
	http.HandleFunc("/labindex", handleActivityIndexLab)
	http.HandleFunc("/labtest", handleLabTest) // not real code. just for practicing new concepts and experimenting. will be removed eventually
	http.HandleFunc("/labicon", handleLabIcon) // not real code. just for practicing new concepts and experimenting. will be removed eventually
}
