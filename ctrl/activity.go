package ctrl

import (
	"encoding/json"
	//"fmt"
	"net/http"

	"model"
	"types"

	//"github.com/gorilla/mux"
	"google.golang.org/appengine"
	//"google.golang.org/appengine/log"
)

func HandleActivityLogPost(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// read the json into struct //TODO: need to complete this
	activityLog := types.ActivityLog{}

	if err := json.NewDecoder(r.Body).Decode(&activityLog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// persist it in db
	if err := model.ActivityLogPost(c, &activityLog); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// send response back to w: http status & json of the record that is created
	w.WriteHeader(http.StatusCreated)

}

func HandleActivityLogPut(w http.ResponseWriter, r *http.Request) {
}

func HandleActivityPost(w http.ResponseWriter, r *http.Request) {
}
func HandleActivityGet(w http.ResponseWriter, r *http.Request) {
}
func HandleActivityDelete(w http.ResponseWriter, r *http.Request) {
}
