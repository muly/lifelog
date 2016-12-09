package lifelog

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	gorillacontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	//"google.golang.org/appengine/log"
)

func HandleGoalPost(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	//c := appengine.NewContext(r)
	var c context.Context
	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}

	goal := Goal{}

	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		WriteResponse(w, http.StatusBadRequest, "application/json", ErrorResponse{err.Error()})
		return
	}

	//  if record already exists with the same goal name, then return
	goalSrc := Goal{}
	goalSrc.Name = goal.Name
	if err := goalSrc.Get(c); err == ErrorNoMatch {
		// do nothing
	} else if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		//http.Error(w, ErrorResponse{err.Error()}, http.StatusInternalServerError)
		WriteResponse(w, http.StatusInternalServerError, "application/json", ErrorResponse{err.Error()})
		return
	} else {
		//http.Error(w, ErrorResponse{"record already exists"}, http.StatusBadRequest)
		WriteResponse(w, http.StatusBadRequest, "application/json", ErrorResponse{"record already exists"})
		return
	}

	goal.CreatedOn = time.Now()

	if err := goal.Put(c); err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		WriteResponse(w, http.StatusInternalServerError, "application/json", ErrorResponse{err.Error()})
		return
	}

	//w.WriteHeader(http.StatusCreated)
	//if err := json.NewEncoder(w).Encode(goal); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	WriteResponse(w, http.StatusCreated, "application/json", goal)

}

// HandleGoalPut handles the PUT operation on the Goal entity type.
// Pass the goal string key in the URI
// And pass the json body with all the fields of goal struct.
// Pass all the fields. if a field is not changed, pass the unchanged value. Any missing fields will result in updating the database with the respective zero value, so Make sure you pass all the fields, even though the value is not changed.
func HandleGoalPut(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	var c context.Context
	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}

	goal := Goal{}

	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)

	// if the goal name (string key) provided in the URI doesn't exist in database, then return
	goalSrc := Goal{}
	goalSrc.Name = params["id"]
	if err := goalSrc.Get(c); err == ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if goal name from body has a value other than the actual goal name in db; i.e if goal name is being changed, dont allow
	if goal.Name != "" && StringKey(goal.Name) != params["id"] { // TODO: Bug: changing the goal name to its equivalent string key is permitted. need to troubleshoot and fix so that any change is not allowed.
		http.Error(w, "cannot update key column - Goal Name", http.StatusBadRequest)
		return
	}

	//
	goal.ModifiedOn = time.Now()

	// update
	if err := goal.Put(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(goal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleGoalGet(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	var c context.Context
	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}

	params := mux.Vars(r)

	goalName, exists := params["id"]
	if !exists {
		http.Error(w, "Goal parameter is missing in URI", http.StatusBadRequest)
		return
	}

	goal := Goal{}
	goal.Name = goalName

	// if given goal is not found, return appropriate error
	if err := goal.Get(c); err == ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(goal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//w.WriteHeader(http.StatusOK) // removed because of "multiple response.WriteHeader calls" error

}
func HandleGoalDelete(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	var c context.Context
	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}

	params := mux.Vars(r)

	goalName, exists := params["id"]
	if !exists {
		w.WriteHeader(http.StatusInternalServerError)
		// add error notesmessage that "goal parameter is missing"
		return
	}

	goal := Goal{}
	goal.Name = goalName

	err := goal.Delete(c)
	if err == ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))

}

func HandleGoalsGet(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	var c context.Context
	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}

	vars, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	goals := Goals{}
	goalFilter := Goal{}

	if val, exists := vars["name"]; exists {
		goalFilter.Name = val[0]

	}
	if val, exists := vars["notes"]; exists {
		goalFilter.Notes = val[0]
	}

	var limit, offset int

	if val, exists := vars["pagesize"]; exists {
		if limit, err = strconv.Atoi(val[0]); err != nil {
			http.Error(w, "pagesize should be a number. "+err.Error(), http.StatusBadRequest)
		}
	} else {
		limit = PageSize
	}

	if val, exists := vars["page"]; exists {
		if offset, err = strconv.Atoi(val[0]); err != nil {
			http.Error(w, "page should be a number. "+err.Error(), http.StatusBadRequest)
		}
		offset = (offset - 1) * limit
	}

	if err := goals.Get(c, goalFilter, offset, limit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(goals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
