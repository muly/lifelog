package lifelog

import (
	"fmt"
	"net/http"

	"ctrl"

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
	r.HandleFunc("/activitylog", ctrl.HandleActivityLogPost).Methods("POST")
	// activity log: update
	r.HandleFunc("/activitylog/{activitylogid}", ctrl.HandleActivityLogPut).Methods("PUT")
	// activity log: delete
	r.HandleFunc("/activitylog/{activitylogid}", ctrl.HandleActivityLogDelete).Methods("DELETE")
	// activity log: search
	r.HandleFunc("/activitylog", ctrl.HandleActivityLogsGet).Methods("GET")
	r.HandleFunc("/activitylog/{activitylogid}", ctrl.HandleActivityLogGet).Methods("GET")

	////
	// activity: add
	r.HandleFunc("/activity", ctrl.HandleActivityPost).Methods("POST")
	r.HandleFunc("/activity/{activityid}", ctrl.HandleActivityPut).Methods("PUT")
	// activity: delete
	r.HandleFunc("/activity/{activityid}", ctrl.HandleActivityDelete).Methods("DELETE")
	// activity: search
	r.HandleFunc("/activity", ctrl.HandleActivitiesGet).Methods("GET")
	r.HandleFunc("/activity/{activityid}", ctrl.HandleActivityGet).Methods("GET")

	//// goal
	r.HandleFunc("/goal", ctrl.HandleGoalPost).Methods("POST")
	r.HandleFunc("/goal/{goal}", ctrl.HandleGoalGet).Methods("GET")
	r.HandleFunc("/goal/{goal}", ctrl.HandleGoalPut).Methods("PUT")
	r.HandleFunc("/goal/{goal}", ctrl.HandleGoalDelete).Methods("DELETE")
	r.HandleFunc("/goal", ctrl.HandleGoalsGet).Methods("GET")

	return r

}
