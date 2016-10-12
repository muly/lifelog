package ctrl

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"time"

	"model"
	"types"
	"util"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	//"google.golang.org/appengine/log"
)

func HandleGoalPost(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	goal := model.Goal{}

	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//  if record already exists with the same goal name, then return
	goalSrc := model.Goal{}
	goalSrc.Name = goal.Name
	if err := goalSrc.Get(c); err == types.ErrorNoMatch {
		// do nothing
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		http.Error(w, "record already exists", http.StatusBadRequest)
		return
	}

	goal.CreatedOn = time.Now()

	if err := goal.Put(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(goal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// HandleGoalPut handles the PUT operation on the Goal entity type.
// Pass the goal string key in the URI
// And pass the json body with all the fields of goal struct.
// Pass all the fields. if a field is not changed, pass the unchanged value. Any missing fields will result in updating the database with the respective zero value, so Make sure you pass all the fields, even though the value is not changed.
func HandleGoalPut(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	goal := model.Goal{}

	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)

	// if the goal name (string key) provided in the URI doesn't exist in database, then return
	goalSrc := model.Goal{}
	goalSrc.Name = params["goal"]
	if err := goalSrc.Get(c); err == types.ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if goal name from body has a value other than the actual goal name in db; i.e if goal name is being changed, dont allow
	if goal.Name != "" && util.StringKey(goal.Name) != params["goal"] { // TODO: Bug: changing the goal name to its equivalent string key is permitted. need to troubleshoot and fix so that any change is not allowed.
		http.Error(w, "cannot update key column - Goal Name", http.StatusBadRequest)
		return
	}

	//
	goal.ModifiedOn = time.Now()

	// update
	if err := goal.Put(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//w.WriteHeader(http.StatusInternalServerError)
		//w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(goal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func HandleGoalGet(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	params := mux.Vars(r)

	goalName, exists := params["goal"]
	if !exists {
		http.Error(w, "Goal parameter is missing in URI", http.StatusBadRequest)
		return
	}

	goal := model.Goal{}
	goal.Name = goalName

	// if given goal is not found, return appropriate error
	if err := goal.Get(c); err == types.ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(goal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//w.WriteHeader(http.StatusOK) // removed because of "multiple response.WriteHeader calls" error

}
func HandleGoalDelete(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	params := mux.Vars(r)

	goalName, exists := params["goal"]
	if !exists {
		w.WriteHeader(http.StatusInternalServerError)
		// add error notesmessage that "goal parameter is missing"
		return
	}

	goal := model.Goal{}
	goal.Name = goalName

	err := goal.Delete(c)
	if err == types.ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusOK)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func HandleGoalsGet(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	goals := model.Goals{}
	if err := goals.Get(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(goals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
