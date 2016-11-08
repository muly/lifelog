package lifelog

import (
	//"github.com/muly/aeunittest"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
)

var (
	testserver *httptest.Server = httptest.NewServer(Handlers())
)

func init() {
	//testserver =
}

func TestAPI(t *testing.T) {
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	h := Handlers()

	testGoal(t, c, h)
	testActivity(t, c, h)
	testActivityLog(t, c, h)

	testserver.Close()
}
