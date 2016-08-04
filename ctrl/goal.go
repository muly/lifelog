package ctrl

import (
	"encoding/json"
	//"fmt"
	"net/http"

	"model"
	"types"

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

	goal.SetDefaults()

	if err := goal.Post(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		w.WriteHeader(http.StatusInternalServerError)
		// add error notesmessage that "goal parameter is missing"
		return
	}

	goal := model.Goal{}
	goal.Name = goalName

	err := goal.Get(c)
	if err == types.ErrorNoMatch {
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func HandleActivityLogDelete(w http.ResponseWriter, r *http.Request) {
}
func HandleActivityPut(w http.ResponseWriter, r *http.Request) {
}
func HandleGoalPut(w http.ResponseWriter, r *http.Request) {
}
func HandleGoalsGet(w http.ResponseWriter, r *http.Request) {
}
