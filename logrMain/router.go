package logrMain

import (
	"fmt"
	"net/http"

	"ctrl"

	"github.com/gorilla/mux"
)

func init() {

	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")

	////
	// activity log: add
	r.HandleFunc("/activitylog", ctrl.HandleActivityLogPost).Methods("POST")
	// activity log: update
	r.HandleFunc("/activitylog/{activitylogid}", ctrl.HandleActivityLogPut).Methods("PUT")
	// activity log: delete
	r.HandleFunc("/activitylog/{activitylogid}", ctrl.HandleActivityLogDelete).Methods("DELETE")

	////
	// activity: add
	r.HandleFunc("/activity", ctrl.HandleActivityPost).Methods("POST")
	r.HandleFunc("/activity/{activityid}", ctrl.HandleActivityPut).Methods("PUT")
	// activity: delete
	r.HandleFunc("/activity", ctrl.HandleActivityDelete).Methods("DELETE")
	// activity: search
	r.HandleFunc("/activity", ctrl.HandleActivityGet).Methods("GET")

	//// goal
	r.HandleFunc("/goal", ctrl.HandleGoalPost).Methods("POST")
	r.HandleFunc("/goal/{goal}", ctrl.HandleGoalGet).Methods("GET")
	r.HandleFunc("/goal/{goal}", ctrl.HandleGoalPut).Methods("PUT")
	r.HandleFunc("/goal/{goal}", ctrl.HandleGoalDelete).Methods("DELETE")
	r.HandleFunc("/goal", ctrl.HandleGoalsGet).Methods("GET")
	// TODO: goal search. need to

	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world! 123")

}
