package lifelog

import (
	//"github.com/muly/aeunittest"
	"net/http"
	//"net/http/httptest"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/muly/aeunittest"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine/aetest"
)

var (
	h *mux.Router = Handlers()
)

func TestGoal(t *testing.T) {
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	testGoal(t, c, h)
}

func TestActivity(t *testing.T) {
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	testActivity(t, c, h)
}

func TestActivityLog(t *testing.T) {
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	testActivityLog(t, c, h)
}

//Note: had to write three separate Test* functions and call the individual test* functions, as wrapping all the 3 test* functions into single Test function is causing performance issues

func testActivity(t *testing.T, c context.Context, h http.Handler) {
	tcs := aeunittest.TestCases{}

	if err := tcs.Load(`testcases\lifelog test cases - Activity.csv`, ',', true); err != nil {
		t.Fatal(err)
	}

	for _, tc := range tcs { // run each test case
		// set the common parameters related to webapp and testing.
		tc.Context = c
		tc.Handler = h
		tc.T = t

		tc.RunCheckStatusCode()

	}
}

func testGoal(t *testing.T, c context.Context, h http.Handler) {

	// Load the test case data
	tcs := aeunittest.TestCases{}
	if err := tcs.Load(`testcases\lifelog test cases - Goal.csv`, ',', true); err != nil {
		t.Fatal(err)
	}

	for _, tc := range tcs { // run each test case
		// set the common parameters related to webapp and testing.
		tc.Context = c
		tc.Handler = h
		tc.T = t

		// execute test case to check the status code and capture the response body
		gotResponseBody := tc.RunCheckStatusCode()

		if tc.SkipFlag { //skip the Response Body test if skip flag is true
			continue
		}

		if tc.WantStatusCode/100 != 2 { // skip Response Body Test for non-success cases
			continue
		}

		// modify the 'got' to remove the system fields
		got := GoalSimple{}
		if err := json.Unmarshal(gotResponseBody, &got); err != nil {
			tc.Error(tc.Name, ": Got Response Body invalid format: \n", string(gotResponseBody), "\n", err.Error())
			continue
		}

		// modify the 'want' to remove the system fields, if any
		want := GoalSimple{}
		if err := json.Unmarshal([]byte(tc.WantResponseBody), &want); err != nil {
			tc.Error(tc.Name, ": Want Response Body invalid format: \n", tc.WantResponseBody, "\n", err.Error())
			continue
		}

		// compare the 'got' with 'want', and report if not matching
		if !reflect.DeepEqual(got, want) {
			tc.Error(tc.Name, ": Response Body : wanted ", want, " but got ", got)
			continue
		}

	}

}

func testActivityLog(t *testing.T, c context.Context, h http.Handler) {
	tcs := aeunittest.TestCases{}

	if err := tcs.Load(`testcases\lifelog test cases - ActivityLog.csv`, ',', true); err != nil {
		t.Fatal(err)
	}

	for _, tc := range tcs { // run each test case
		// set the common parameters related to webapp and testing.
		tc.Context = c
		tc.Handler = h
		tc.T = t

		tc.RunCheckStatusCode()
	}

	//t.Log("ActivityLog test cases execution completed")
}
