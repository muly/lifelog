package ctrl

import (
	"encoding/json"
	//"fmt"
	"model"
	"net/http"
	//"net/url"
	"time"
	"types"
	"util"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	//"google.golang.org/appengine/log"
	//"google.golang.org/appengine/datastore"
)

func HandleActivityPost(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	act := model.Activity{}

	if err := json.NewDecoder(r.Body).Decode(&act); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//  if record already exists with the same Activity name, then return
	actSrc := model.Activity{}
	actSrc.Name = act.Name
	if err := actSrc.Get(c); err == types.ErrorNoMatch {
		// do nothing
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		http.Error(w, "record already exists", http.StatusBadRequest)
		return
	}

	act.CreatedDate = time.Now() //SetDefaults()

	if err := act.Put(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(act); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func HandleActivityPut(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	act := model.Activity{}

	if err := json.NewDecoder(r.Body).Decode(&act); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)

	// if the goal name (string key) provided in the URI doesn't exist in database, then return
	actsrc := model.Activity{}
	actsrc.Name = params["activityid"]
	if err := actsrc.Get(c); err == types.ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if goal name from body has a value other than the actual goal name in db; i.e if goal name is being changed, dont allow
	if actsrc.Name != "" && util.StringKey(actsrc.Name) != params["activityid"] { // TODO: Bug: changing the goal name to its equivalent string key is permitted. need to troubleshoot and fix so that any change is not allowed.
		http.Error(w, "cannot update key column - Activity Name", http.StatusBadRequest)
		return
	}

	//
	act.ModifiedOn = time.Now() //SetDefaults()

	// update
	if err := act.Put(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//w.WriteHeader(http.StatusInternalServerError)
		//w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(act); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func HandleActivityGet(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	params := mux.Vars(r)

	act := model.Activity{}
	act.Name = params["activityid"]

	// if given goal is not found, return appropriate error
	if err := act.Get(c); err == types.ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(act); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//w.WriteHeader(http.StatusOK) // removed because of "multiple response.WriteHeader calls" error

}

func HandleActivitysGet(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	acts := model.Activitys{}
	if err := acts.Get(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(acts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func HandleActivityDelete(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	params := mux.Vars(r)
	act := model.Activity{}

	act.Name = params["activityid"]

	err := act.Delete(c)
	if err == types.ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusOK)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
