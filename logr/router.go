package logr

import (
	"fmt"
	"net/http"

	//"ctrl"

	"github.com/gorilla/mux"
)

func init() {

	r := Handlers()

	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world! 123")

}

func Handlers() *mux.Router {

	r := mux.NewRouter()
	//r.HandleFunc("/", handler).Methods("GET")

	////
	// activity log: add
	r.HandleFunc("/activitylog", HandleActivityLogPost).Methods("POST")
	// activity log: update
	r.HandleFunc("/activitylog/{activitylogid}", HandleActivityLogPut).Methods("PUT")
	// activity log: delete
	r.HandleFunc("/activitylog/{activitylogid}", HandleActivityLogDelete).Methods("DELETE")
	// activity log: search
	r.HandleFunc("/activitylog", HandleActivityLogsGet).Methods("GET")
	r.HandleFunc("/activitylog/{activitylogid}", HandleActivityLogGet).Methods("GET")

	////
	// activity: add
	r.HandleFunc("/activity", HandleActivityPost).Methods("POST")
	r.HandleFunc("/activity/{activityid}", HandleActivityPut).Methods("PUT")
	// activity: delete
	r.HandleFunc("/activity/{activityid}", HandleActivityDelete).Methods("DELETE")
	// activity: search
	r.HandleFunc("/activity", HandleActivitiesGet).Methods("GET")
	r.HandleFunc("/activity/{activityid}", HandleActivityGet).Methods("GET")

	//// goal
	r.HandleFunc("/goal", HandleGoalPost).Methods("POST")
	r.HandleFunc("/goal/{goal}", HandleGoalGet).Methods("GET")
	r.HandleFunc("/goal/{goal}", HandleGoalPut).Methods("PUT")
	r.HandleFunc("/goal/{goal}", HandleGoalDelete).Methods("DELETE")
	r.HandleFunc("/goal", HandleGoalsGet).Methods("GET")

	return r

}
