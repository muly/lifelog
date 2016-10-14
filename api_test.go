package lifelog

import (
	"fmt"

	//"ctrl"
	//"net/http"
	"net/http/httptest"
	//"testing"
)

func init() {
	server := httptest.NewServer(Handlers())

	usersUrl := fmt.Sprintf("%s/users", server.URL)

	//var t *testing.T
	fmt.Println("Hello.................", usersUrl)
	//t.Log(usersUrl, "$$$$$$$$$$$$$$$$$$$$$$$$$$$")
}
