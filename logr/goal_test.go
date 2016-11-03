package logr

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muly/aeunittest"

	"golang.org/x/net/context"
	"google.golang.org/appengine/aetest"
)

var (
	testserver *httptest.Server
	goalUrl    string
)

func init() {
	testserver = httptest.NewServer(Handlers())
	goalUrl = testserver.URL + "/goal"

}

func TestGoal(t *testing.T) {
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	h := Handlers()

	testGoal(t, c, h)

}

func testGoal(t *testing.T, c context.Context, h http.Handler) {
	tcs := aeunittest.TestCases{}
	tc := aeunittest.TestCase{}

	//reset (tc), input (tc), append (to tcs) the test cases one after the other
	tc = aeunittest.TestCase{} // reset
	//input
	tc.Name = "Goal Post new record test"
	tc.RequestBody = `{"Name":"test1","Notes":"test"}`
	tc.HttpVerb = "POST"
	tc.Uri = goalUrl
	tc.WantStatusCode = http.StatusCreated
	tcs = append(tcs, tc) // append

	tc = aeunittest.TestCase{}
	tc.Name = "Goal Post duplicate record test"
	tc.RequestBody = `{"Name":"test1","Notes":"test"}`
	tc.HttpVerb = "POST"
	tc.Uri = goalUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tc = aeunittest.TestCase{}
	tc.Name = "Goal Get existing record test"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusOK
	tcs = append(tcs, tc)

	for _, tc := range tcs { // run each test case
		// set the common parameters related to webapp and testing.
		tc.Context = c
		tc.Handler = h
		tc.T = t

		tc.Run()
	}

}
