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
)

func HandleActivityLogPost(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	var c context.Context
	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}

	al := ActivityLog{}

	if err := json.NewDecoder(r.Body).Decode(&al); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//  if record already exists with the same Activity name, then return
	alSrc := ActivityLog{}
	alSrc.Name = al.Name
	if err := alSrc.Get(c); err == ErrorNoMatch {
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(al); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleActivityLogPut(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	var c context.Context
	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}

	al := ActivityLog{}

	if err := json.NewDecoder(r.Body).Decode(&al); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	if al.Name != "" && id != al.Name {
		http.Error(w, "key in the URI and key in Request body are not matching", http.StatusBadRequest)
		return
	}
	al.Name = id

	// if the goal name (string key) provided in the URI doesn't exist in database, then return
	alsrc := ActivityLog{}
	alsrc.Name = id
	if err := alsrc.Get(c); err == ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if goal name from body has a value other than the actual goal name in db; i.e if goal name is being changed, dont allow
	if alsrc.Name != "" && StringKey(alsrc.Name) != params["id"] { // TODO: Bug: changing the goal name to its equivalent string key is permitted. need to troubleshoot and fix so that any change is not allowed.
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(al); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleActivityLogGet(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	var c context.Context
	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}

	params := mux.Vars(r)

	al := ActivityLog{}
	al.Name = params["id"]

	// if given goal is not found, return appropriate error
	if err := al.Get(c); err == ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(al); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//w.WriteHeader(http.StatusOK) // removed because of "multiple response.WriteHeader calls" error

}

func HandleActivityLogsGet(w http.ResponseWriter, r *http.Request) {
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

	als := ActivityLogs{}
	alFilter := ActivityLog{}

	if val, exists := vars["name"]; exists {
		alFilter.Name = val[0]

	}
	if val, exists := vars["notes"]; exists {
		alFilter.Notes = val[0]
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

	if err := als.Get(c, alFilter, offset, limit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(als); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func HandleActivityLogDelete(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	var c context.Context
	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}

	params := mux.Vars(r)
	al := ActivityLog{}

	al.Name = params["id"]

	err := al.Delete(c)
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
