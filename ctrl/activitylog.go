package ctrl

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"net/url"
	//"types"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	//"google.golang.org/appengine/log"
	"google.golang.org/appengine/datastore"
)

func HandleActivityLogGet(w http.ResponseWriter, r *http.Request) {
	q, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	c := appengine.NewContext(r)
	al := StructActivityLog{}
	//p := mux.Vars(r)
	al.Name = q["Name"][0]
	//fmt.Printf(al.Name)

	// generate key
	key := datastore.NewKey(c, "ActivityLog", al.Name, 0, nil)
	if err := datastore.Get(c, key, &al); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Name = %q\nNotes=%q\n", al.Name, al.Notes)

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

type StructActivityLog struct {
	Name      string
	Notes     string `json:"Notes,omitempty"`
	StartTime string `json:"StartTime,omitempty"`
	//EndTime   time.Time
}
