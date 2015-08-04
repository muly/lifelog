package helpers

import (
	//"appengine"
	"testing"
)

func TestAdd(t *testing.T) {
	x := 1
	y := 2
	want := 4

	if got := add(x, y); got != want {
		t.Error("wanted: ", want, ", but got: ", got)
	}
}

func TestInsertActivity(t *testing.T) {
	//appengine.Context
	//InsertActivity(c appengine.Context, rec ActivityLog)
	t.Error("Testing InsertActivity().... WIP")
}
