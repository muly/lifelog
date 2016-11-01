package logrMain

import (
	"github.com/gorilla/mux"

	"ctrl"
	"fmt"
	"net/http"
)

func init() {

	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")

	////
	// activity log: add
	r.HandleFunc("/activitylog", ctrl.HandleActivityLogPost).Methods("POST")
	// activity log: update
	r.HandleFunc("/activitylog/{id}", ctrl.HandleActivityLogPut).Methods("PUT")
	// activity log: delete
	r.HandleFunc("/activitylog/{id}", ctrl.HandleActivityLogDelete).Methods("DELETE")
	// activity log: search
	r.HandleFunc("/activitylog", ctrl.HandleActivityLogsGet).Methods("GET")
	r.HandleFunc("/activitylog/{id}", ctrl.HandleActivityLogGet).Methods("GET")

	////
	// activity: add
	r.HandleFunc("/activity", ctrl.HandleActivityPost).Methods("POST")
	r.HandleFunc("/activity/{id}", ctrl.HandleActivityPut).Methods("PUT")
	// activity: delete
	r.HandleFunc("/activity/{id}", ctrl.HandleActivityDelete).Methods("DELETE")
	// activity: search
	r.HandleFunc("/activity", ctrl.HandleActivitiesGet).Methods("GET")
	r.HandleFunc("/activity/{id}", ctrl.HandleActivityGet).Methods("GET")

	//// goal
	r.HandleFunc("/goal", ctrl.HandleGoalPost).Methods("POST")
	r.HandleFunc("/goal/{id}", ctrl.HandleGoalGet).Methods("GET")
	r.HandleFunc("/goal/{id}", ctrl.HandleGoalPut).Methods("PUT")
	r.HandleFunc("/goal/{id}", ctrl.HandleGoalDelete).Methods("DELETE")
	r.HandleFunc("/goal", ctrl.HandleGoalsGet).Methods("GET")
	// TODO: goal search. need to

	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world! 123")

}
