package helpers

import (
	//"appengine/aetest"
	//"appengine/datastore"
	//"appengine/user"
	//"testing"
	//"time"
)


/*
func TestInsertActivity_Experimenting(t *testing.T) {

	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	testActivityName := "unit testing: TestInsertActivity: " + time.Now().String()

	a := ActivityLog{
		ActivityName: testActivityName,
	}

	//var UserId *user.User
	//user.Current(c).String()
	parentKey := datastore.NewKey(c, "activity", "", 0, nil)

	//childKey := datastore.NewIncompleteKey(c, "activityRecord", parentKey)

	_, err = datastore.Put(c, parentKey, &a)

	if err != nil {
		t.Error(testActivityName + err.Error())
	}
}


func TestInsertActivity_thatPASS(t *testing.T) {

	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	testActivityName := "unit testing: TestInsertActivity: " + time.Now().String()

	a := ActivityLog{
		ActivityName: testActivityName,
	}

	UserId := "" //user.Current(c).String()
	parentKey := datastore.NewKey(c, "activity", UserId, 0, nil)

	//childKey := datastore.NewIncompleteKey(c, "activityRecord", parentKey)

	_, err = datastore.Put(c, parentKey, &a)

	if err != nil {
		t.Error(testActivityName + err.Error())
	}
}


//func TestGetActivity(t *testing.T) { // pick a record that is already in datastore and see if you can retrieve it. not sure if that is possible because of the fake context that we are using here

func TestInsertActivity_JustInsert(t *testing.T) {

	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	testActivityName := "unit testing: TestInsertActivity: " + time.Now().String()

	a := ActivityLog{
		ActivityName: testActivityName,
	}

	if err := InsertActivity(c, a); err != nil {
		t.Error(testActivityName + err.Error())
	}
}

/*
func TestInsertActivity_InsertAndRetrieve(t *testing.T) {

	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	testActivityName := "unit testing: TestInsertActivity: " + time.Now().String()

	a := ActivityLog{
		ActivityName: testActivityName,
	}

	err = InsertActivity(c, a)

	//--------------------------------
	var f []Filter // filter slice to store the number filters
	var OrderBy string

	// retrieve the Started activities
	f = []Filter{{"ActuvutyName=", testActivityName}}
	OrderBy = "-TimeStamp"
	dst := GetActivity(c, f, OrderBy)

	if len(dst) != 1 {
		t.Error("Inserted 1 record, but retrieved ", len(dst), " records")
	}
}
*/
