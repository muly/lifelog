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
	r.HandleFunc("/activity", ctrl.HandleActivityPost).Methods("POST")
	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")

}
