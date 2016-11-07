package lifelog

import (
	"github.com/gorilla/mux"

	//"ctrl"

	"errors"
	//"fmt"
	"net/http"
)

var ErrorNoMatch = errors.New("No Matching Record")

var PageSize = 10 // default page size

func init() {

	r := Handlers()

	http.Handle("/", r)
}

func Handlers() *mux.Router {

	r := mux.NewRouter()
	//r.HandleFunc("/", handler).Methods("GET")

	////
	// activity log: add
	r.HandleFunc("/activitylog", HandleActivityLogPost).Methods("POST")
	// activity log: update
	r.HandleFunc("/activitylog/{id}", HandleActivityLogPut).Methods("PUT")
	// activity log: delete
	r.HandleFunc("/activitylog/{id}", HandleActivityLogDelete).Methods("DELETE")
	// activity log: search
	r.HandleFunc("/activitylog", HandleActivityLogsGet).Methods("GET")
	r.HandleFunc("/activitylog/{id}", HandleActivityLogGet).Methods("GET")

	////
	// activity: add
	r.HandleFunc("/activity", HandleActivityPost).Methods("POST")
	r.HandleFunc("/activity/{id}", HandleActivityPut).Methods("PUT")
	// activity: delete
	r.HandleFunc("/activity/{id}", HandleActivityDelete).Methods("DELETE")
	// activity: search
	r.HandleFunc("/activity", HandleActivitiesGet).Methods("GET")
	r.HandleFunc("/activity/{id}", HandleActivityGet).Methods("GET")

	//// goal
	r.HandleFunc("/goal", HandleGoalPost).Methods("POST")
	r.HandleFunc("/goal/{id}", HandleGoalGet).Methods("GET")
	r.HandleFunc("/goal/{id}", HandleGoalPut).Methods("PUT")
	r.HandleFunc("/goal/{id}", HandleGoalDelete).Methods("DELETE")
	r.HandleFunc("/goal", HandleGoalsGet).Methods("GET")

	return r

}

/*func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world! 123")

}*/
