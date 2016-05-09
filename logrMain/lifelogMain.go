package logrMain

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
