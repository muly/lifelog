package ctrl

import (
	"fmt"
	"net/http"
)

func HandleActivityPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Inserts Activity")

}
