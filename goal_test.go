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
	//testserver *httptest.Server
	goalUrl string
)

func init() {
	//testserver = httptest.NewServer(Handlers())
	goalUrl = testserver.URL + "/goal"

}

/*func TestGoal(t *testing.T) {
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	h := Handlers()

	//testGoal(t, c, h)

}*/

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
	tc.Name = "Goal Get existing record test"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode =  http.StatusOK

	tcs = append(tcs, tc)
	tc.Name = "Goal Post new record with duplicate key "
	tc.RequestBody = `{"Name":"test1","Notes":"test"}`
	tc.HttpVerb = "POST"
	tc.Uri = goalUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc) 

	tcs = append(tcs, tc)
	tc.Name = "Goal Post Field validation-String field"
	tc.RequestBody = `{"Name":123,"Notes":"Test123"}`
	tc.HttpVerb = "POST"
	tc.Uri = goalUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Post Field validation-String field"
	tc.RequestBody = `{"Name":"123","Notes":123}`
	tc.HttpVerb = "POST"
	tc.Uri = goalUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Post Field validation-Passing blank in mandatory fields"
	tc.RequestBody = `{"Name":""."Notes":"test4"}`
	tc.HttpVerb = "POST"
	tc.Uri = goalUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Post with typo in field name in json 
/Sending Extra field"
	tc.RequestBody = `{"Name":"test2","Note":"test2"}`
	tc.HttpVerb = "POST"
	tc.Uri = goalUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Post Invalid Json"
	tc.RequestBody = `{"Name":"test3"Notes":"testnotes3"}`
	tc.HttpVerb = "POST"
	tc.Uri = goalUrl
	tc.WantStatusCode = http.StatusNotFound
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Post without body"
	tc.RequestBody = ``
	tc.HttpVerb = "POST"
	tc.Uri = goalUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc) 

	tcs = append(tcs, tc)
	tc.Name = "Goal Post with wrong key name"
	tc.RequestBody = `{"id":"test4","Notes":"testNotes4"}`
	tc.HttpVerb = "POST"
	tc.Uri = goalUrl
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc) 




	tc = aeunittest.TestCase{}
	tc.Name = "Goal Get with key"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusOK
	tcs = append(tcs, tc)

	tc = aeunittest.TestCase{}
	tc.Name = "Goal Get time field in response should be of RFC 3339 format "
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusOK
	tcs = append(tcs, tc)

	tc = aeunittest.TestCase{}
	tc.Name = "Goal Get with key"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = goalUrl + "/xyz"
	tc.WantStatusCode = http.StatusNotFound
	tcs = append(tcs, tc)

	tc = aeunittest.TestCase{}
	tc.Name = "Goal Get with blank key"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = goalUrl 
	tc.WantStatusCode = http.StatusNotFound
	tcs = append(tcs, tc)



	tcs = append(tcs, tc)
	tc.Name = "Goal Put Field validation-String field"
	tc.RequestBody = `{"Name":"test1","Notes":123}`
	tc.HttpVerb = "PUT"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc) 

	tcs = append(tcs, tc)
	tc.Name = "Goal Put Field validation-Passing blank in mandatory fields"
	tc.RequestBody = `{"Name":"","Notes":"test4"}`
	tc.HttpVerb = "PUT"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc) 

	tcs = append(tcs, tc)
	tc.Name = "Goal Put Editing a record"
	tc.RequestBody = `{"Name":"test1",'"Notes":"TestNotes New"}`
	tc.HttpVerb = "PUT"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.Statusok
	tcs = append(tcs, tc) 

	tc = aeunittest.TestCase{}
	tc.Name = "Goal Put Successful saving of the record to database"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusOK
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Put new key should NOT be allowed"
	tc.RequestBody = `{"Name":"xyz","Notes":"Notes123"}`
	tc.HttpVerb = "PUT"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc) 

	tcs = append(tcs, tc)
	tc.Name = "Goal Put with typo in field name in json 
/Sending Extra field"
	tc.RequestBody = `{"Name":"test1","Note":"test2"}`
	tc.HttpVerb = "PUT"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Put Invalid Json"
	tc.RequestBody = `{"Name":"test3"Notes":"testnotes3"}`
	tc.HttpVerb = "PUT"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Put key in URL not same as key in Body"
	tc.RequestBody = `{"Name":"test10","Notes":"Test123"}`
	tc.HttpVerb = "PUT"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusBadRequest
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Put without parameter in URI"
	tc.RequestBody = `{"Name":"test1","Notes":"test1new"}`
	tc.HttpVerb = "PUT"
	tc.Uri = goalUrl 
	tc.WantStatusCode = http.StatusForbidden
	tcs = append(tcs, tc)





	tcs = append(tcs, tc)
	tc.Name = "Goal Delete without key"
	tc.RequestBody = ``
	tc.HttpVerb = "DELETE"
	tc.Uri = goalUrl 
	tc.WantStatusCode = http.StatusForbidden
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Delete remove record"
	tc.RequestBody = ``
	tc.HttpVerb = "DELETE"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.statusok
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Delete Get same record"
	tc.RequestBody = ``
	tc.HttpVerb = "GET"
	tc.Uri = goalUrl + "/test1"
	tc.WantStatusCode = http.StatusNotFound
	tcs = append(tcs, tc)

	tcs = append(tcs, tc)
	tc.Name = "Goal Delete non Existing record"
	tc.RequestBody = ``
	tc.HttpVerb = "DELETE"
	tc.Uri = goalUrl + "/xyz"
	tc.WantStatusCode = http.StatusNotFound
	tcs = append(tcs, tc)

	for _, tc := range tcs { // run each test case
		// set the common parameters related to webapp and testing.
		tc.Context = c
		tc.Handler = h
		tc.T = t

		tc.Run()
	}

	t.Log("Goal test cases execution completed")

}
