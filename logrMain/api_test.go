package lifelog

import (
	"fmt"

	//"ctrl"
	"net/http"
	//"github.com/arschles/testsrv"
	"net/http/httptest"
	"testing"
)

func Test123(t *testing.T) {
	h := Handlers()

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/goal", nil)
	if err != nil {
		t.Fatalf("error constructing test HTTP request [%s]", err)
	}

	h.ServeHTTP(w, r)

	url := r.RequestURI

	/*
	   	srv := testsrv.StartServer(r)
	   	// always close the server at the end of each test
	   	defer srv.Close()

	   url:= srv.URLStr()
	*/
	//var t *testing.T
	fmt.Println("Hello.................", url)
	t.Log(url, "$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	fmt.Println("Hello.................", w.Code)
	t.Log(w.Code, "$$$$$$$$$$$$$$$$$$$$$$$$$$$")

}
