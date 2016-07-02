package ctrl

import (
	"fmt"
	"net/http"

	"types"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func HandleActivityPost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Inserts Activity")
	c := appengine.NewContext(req)
	key := datastore.NewIncompleteKey(c, "ActivityLog", nil)
	activitylog := types.ActivityLog{}

	datastore.Put(c, key, &activitylog)

	log.Infof(c, "Inserted Activity")

}
