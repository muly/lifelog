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

/*



func HandleActivityLogGet(w http.ResponseWriter, r *http.Request) {

	p := mux.Vars(r)
	c := appengine.NewContext(r)
	al := StructActivityLog{}
	al.Name = p["activitylogid"]

	// generate key
	key := datastore.NewKey(c, "ActivityLog", al.Name, 0, nil)
	if err := datastore.Get(c, key, &al); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Name = %q\nNotes=%q\n", al.Name, al.Notes)

}

func HandleActivityLogsGet(ww http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	AvtivityLogs := []StructActivityLog{}

}

func HandleActivityLogPut(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	al := StructActivityLog{}
	// read the data from req //{CategoryName:"grocery",description:"new descrption"}
	if err := json.NewDecoder(r.Body).Decode(&al); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Printf(al.Name)
	// generate key
	key := datastore.NewKey(c, "ActivityLog", al.Name, 0, nil)
	_, err := datastore.Put(c, key, &al)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}
func HandleActivityLogDelete(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	c := appengine.NewContext(r)
	al := StructActivityLog{}
	al.Name = p["activitylogid"]
	fmt.Printf(al.Name)

	// generate key
	key := datastore.NewKey(c, "ActivityLog", al.Name, 0, nil)
	if err := datastore.Delete(c, key); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}
func HandleActivityLogPost(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// read the json into struct //TODO: need to complete this
	activityLog := model.ActivityLog{}

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
*/
