package ctrl

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"

	"model"
	"net/http"
	"time"

	"types"
	"util"
)

func HandleActivityLogPost(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	al := model.ActivityLog{}

	if err := json.NewDecoder(r.Body).Decode(&al); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//  if record already exists with the same Activity name, then return
	alSrc := model.ActivityLog{}
	alSrc.Name = al.Name
	if err := alSrc.Get(c); err == types.ErrorNoMatch {
		// do nothing
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		http.Error(w, "record already exists", http.StatusBadRequest)
		return
	}

	al.CreatedOn = time.Now()

	if err := al.Put(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(al); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func HandleActivityLogPut(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	al := model.ActivityLog{}

	if err := json.NewDecoder(r.Body).Decode(&al); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)

	// if the goal name (string key) provided in the URI doesn't exist in database, then return
	alsrc := model.ActivityLog{}
	alsrc.Name = params["activitylogid"]
	if err := alsrc.Get(c); err == types.ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if goal name from body has a value other than the actual goal name in db; i.e if goal name is being changed, dont allow
	if alsrc.Name != "" && util.StringKey(alsrc.Name) != params["activitylogid"] { // TODO: Bug: changing the goal name to its equivalent string key is permitted. need to troubleshoot and fix so that any change is not allowed.
		http.Error(w, "cannot update key column - ActivityLog Name", http.StatusBadRequest)
		return
	}

	//
	al.ModifiedOn = time.Now()

	// update
	if err := al.Put(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//w.WriteHeader(http.StatusInternalServerError)
		//w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(al); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func HandleActivityLogGet(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	params := mux.Vars(r)

	al := model.ActivityLog{}
	al.Name = params["activitylogid"]

	// if given goal is not found, return appropriate error
	if err := al.Get(c); err == types.ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(al); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//w.WriteHeader(http.StatusOK) // removed because of "multiple response.WriteHeader calls" error

}

func HandleActivityLogsGet(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	als := model.ActivityLogs{}
	if err := als.Get(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(als); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func HandleActivityLogDelete(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	params := mux.Vars(r)
	al := model.ActivityLog{}

	al.Name = params["activitylogid"]

	err := al.Delete(c)
	if err == types.ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusOK)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
