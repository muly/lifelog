package lifelog

import (
	"net/http"
	//"net/http/httptest"
	"testing"

	"github.com/muly/aeunittest"

	"golang.org/x/net/context"
	//"google.golang.org/appengine/aetest"
)

var (
	activityUrl string
)

func init() {
	activityUrl = testserver.URL + "/activity"
}

/*func TestActivity(t *testing.T) {
}*/

func testActivity(t *testing.T, c context.Context, h http.Handler) {
	tcs := aeunittest.TestCases{}
	tc := aeunittest.TestCase{}

	//reset (tc), input (tc), append (to tcs) the test cases one after the other
	tc = aeunittest.TestCase{} // reset
	//input
	tc.Name = "Activity Post new record test"
	tc.RequestBody = `{"Name":"test1"}`
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusCreated
	tcs = append(tcs, tc) // append

	tc = aeunittest.TestCase{}
	tc.Name = "Activity Post duplicate record test"
	tc.RequestBody = `{"Name":"test1"}`
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tc = aeunittest.TestCase{}
	tc.Name = "Activity Get existing record test"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusOK
	tcs = append(tcs, tc)

	for _, tc := range tcs { // run each test case
		// set the common parameters related to webapp and testing.
		tc.Context = c
		tc.Handler = h
		tc.T = t

		tc.Run()
	}

	t.Log("Activity test cases execution completed")
}
