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

/*
func TestActivity(t *testing.T) {
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	//testActivity(t, c, h)
}

func TestActivityLog(t *testing.T) {
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	testActivityLog(t, c, h)
}
*/
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

		tc.Run()
	}

	//t.Log("Activity test cases execution completed")
}

func testGoal(t *testing.T, c context.Context, h http.Handler) {
	tcs := aeunittest.TestCases{}

	if err := tcs.Load(`testcases\Goal Small.csv`, ',', true); err != nil {
		t.Fatal(err)
	}

	for _, tc := range tcs { // run each test case
		// set the common parameters related to webapp and testing.
		tc.Context = c
		tc.Handler = h
		tc.T = t

		//tc.Log("Want:", tc.WantResponseBody)

		//tc.Run()
		g1 := tc.Run1()

		got := Goal2{}

		if err := json.Unmarshal(g1, &got); err != nil {
			tc.Error("Want Response Body invalid format: ", err.Error())
			return
		}

		want := Goal2{}

		if err := json.Unmarshal([]byte(tc.WantResponseBody), &want); err != nil {
			tc.Error("Want Response Body invalid format: ", err.Error())
			return
		}

		//		tc.Log("Want:", want)
		//		tc.Log("Got: ", got)
		//tc.Log("DeepEqual:", reflect.DeepEqual(got, want))

		if !reflect.DeepEqual(got, want) {
			tc.Error(tc.Name, ": Response Body : wanted ", want, " but got ", got)
		}

		//tc.Log("Got ", string(GotResponseBody))
	}

	//t.Log("Goal test cases execution completed")
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

		tc.Run()
	}

	//t.Log("ActivityLog test cases execution completed")
}
