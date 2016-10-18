package ctrl

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"

	"model"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"types"
	"util"
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

	act.CreatedOn = time.Now() //SetDefaults()

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

func HandleActivitiesGet(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	vars, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	als := model.Activities{}
	alFilter := model.Activity{}

	if val, exists := vars["name"]; exists {
		alFilter.Name = val[0]

	}
	if val, exists := vars["GoalID"]; exists {
		alFilter.GoalID = val[0]
	}

	var limit, offset int

	if val, exists := vars["pagesize"]; exists {
		if limit, err = strconv.Atoi(val[0]); err != nil {
			http.Error(w, "pagesize should be a number. "+err.Error(), http.StatusBadRequest)
		}
	} else {
		limit = types.PageSize
	}

	if val, exists := vars["page"]; exists {
		if offset, err = strconv.Atoi(val[0]); err != nil {
			http.Error(w, "page should be a number. "+err.Error(), http.StatusBadRequest)
		}
		offset = (offset - 1) * limit
	}

	if err := als.Get(c, alFilter, offset, limit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(als); err != nil {
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
