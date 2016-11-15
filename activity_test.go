package lifelog

/*
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


func testActivity2(t *testing.T, c context.Context, h http.Handler) {
	tcs := aeunittest.TestCases{}
	tc := aeunittest.TestCase{}

	//reset (tc), input (tc), append (to tcs) the test cases one after the other
	tc = aeunittest.TestCase{} // reset
	//input
	tc.Name = "Activity Post new record test"
	tc.RequestBody = `{"Name":"test1","GoalID":"test"}`
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusCreated
	tcs = append(tcs, tc) // append

	tc = aeunittest.TestCase{}
	tc.Name = "Activity Get existing record test"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusOK

	tcs = append(tcs, tc)
	tc.Name = "Activity Post new record with duplicate key "
	tc.RequestBody = `{"Name":"test1","GoalID":"test"}`
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Post Field validation-String field"
	tc.RequestBody = `{"Name":123,"GoalID":"Test123"}`
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Post Field validation-String field"
	tc.RequestBody = `{"Name":"123","GoalID":123}`
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Post Field validation-Passing blank in mandatory fields"
	tc.RequestBody = `{"Name":""."GoalID":"test4"}`
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Post with typo in field name in json/Sending Extra field"
	tc.RequestBody = `{"Name":"test2","ID":"test2"}`
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Post Invalid Json"
	tc.RequestBody = `{"Name":"test3"GoalID":"testnotes3"}`
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusNotFound
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Post without body"
	tc.RequestBody = ``
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Post with wrong key name"
	tc.RequestBody = `{"id":"test4","GoalID":"testNotes4"}`
	tc.HttpVerb = "POST"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tc = aeunittest.TestCase{}
	tc.Name = "Activity Get with key"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusOK
	tcs = append(tcs, tc)

	tc = aeunittest.TestCase{}
	tc.Name = "Activity Get time field in response should be of RFC 3339 format "
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusOK
	tcs = append(tcs, tc)

	tc = aeunittest.TestCase{}
	tc.Name = "Activity Get with key"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = activityUrl + "/xyz"
	tc.WantStatusCode = http.StatusNotFound
	tcs = append(tcs, tc)

	tc = aeunittest.TestCase{}
	tc.Name = "Activity Get with blank key"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusNotFound
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Put Field validation-String field"
	tc.RequestBody = `{"Name":"test1","GoalID":123}`
	tc.HttpVerb = "PUT"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Put Field validation-Passing blank in mandatory fields"
	tc.RequestBody = `{"Name":"","GoalID":"test4"}`
	tc.HttpVerb = "PUT"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Put Editing a record"
	tc.RequestBody = `{"Name":"test1",'"GoalID":"TestNotes New"}`
	tc.HttpVerb = "PUT"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusOK
	tcs = append(tcs, tc)

	tc = aeunittest.TestCase{}
	tc.Name = "Activity Put Successful saving of the record to database"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusOK
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Put new key should NOT be allowed"
	tc.RequestBody = `{"Name":"xyz","GoalID":"Notes123"}`
	tc.HttpVerb = "PUT"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Put with typo in field name in json/Sending Extra field"
	tc.RequestBody = `{"Name":"test1","Id":"test2"}`
	tc.HttpVerb = "PUT"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Put Invalid Json"
	tc.RequestBody = `{"Name":"test3"GoalID":"testnotes3"}`
	tc.HttpVerb = "PUT"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Put key in URL not same as key in Body"
	tc.RequestBody = `{"Name":"test10","GoalID":"Test123"}`
	tc.HttpVerb = "PUT"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Put without parameter in URI"
	tc.RequestBody = `{"Name":"test1","GoalID":"test1new"}`
	tc.HttpVerb = "PUT"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusForbidden
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Delete without key"
	tc.RequestBody = ``
	tc.HttpVerb = "DELETE"
	tc.Uri = activityUrl
	tc.WantStatusCode = http.StatusForbidden
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Delete remove record"
	tc.RequestBody = ``
	tc.HttpVerb = "DELETE"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusOK
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Delete Get same record"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = activityUrl + "/test1"
	tc.WantStatusCode = http.StatusNotFound
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Activity Delete non Existing record"
	tc.RequestBody = ``
	tc.HttpVerb = "DELETE"
	tc.Uri = activityUrl + "/xyz"
	tc.WantStatusCode = http.StatusNotFound
	tcs = append(tcs, tc)

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

		//t.Log("Running Test case:" + tc.Name)
		tc.Run()

	}

	t.Log("Activity test cases execution completed")
}
*/
