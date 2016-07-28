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
	r.HandleFunc("/activitylog", ctrl.HandleActivityLogPut).Methods("PUT")

	// activity: search
	r.HandleFunc("/activity", ctrl.HandleActivityGet).Methods("GET")

	////
	// activity: add
	r.HandleFunc("/activity", ctrl.HandleActivityPost).Methods("POST")

	////
	// activity: delete
	r.HandleFunc("/activity", ctrl.HandleActivityDelete).Methods("DELETE")

	//
	// goal: add
	r.HandleFunc("/goal", ctrl.HandleGoalPost).Methods("POST")
	// goal: search
	r.HandleFunc("/goal/{goal}", ctrl.HandleGoalGet).Methods("GET")
	// goal: delete
	r.HandleFunc("/goal/{goal}", ctrl.HandleGoalDelete).Methods("DELETE")

	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world! 123")

}
