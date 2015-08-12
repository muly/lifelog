package ActivityLoggerMain

import (
	"appengine"
	"appengine/datastore"
	//"appengine/user"
	//"errors"
	"fmt"
	"helpers"
	//"html/template"
	"net/http"
	//"net/url"
	//"strings"
	//"time"
)

type ActivityNameOnly struct {
	ActivityName string
}

const UsefulMaterialIcons = `
<html>
  <head>
    <!-- Material Design Lite -->
    <script src="https://storage.googleapis.com/code.getmdl.io/1.0.2/material.min.js"></script>
    <link rel="stylesheet" href="https://storage.googleapis.com/code.getmdl.io/1.0.2/material.indigo-pink.min.css">
    <!-- Material Design icon font -->
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
  </head>
  <body>
    <table>
      <thead>
        <tr>
          <th>Icon</th>
          <th>Label</th>
        </tr>
      </thead>
      <tbody>
      <!--Groups:::::::::-->
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon mdl-button--colored"> <i class="material-icons">laptop</i> </button></td>
          <td>work related group</td>
        </tr>
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon mdl-button--colored"> <i class="material-icons">home</i> </button></td>
          <td>Home activities related group</td>
        </tr>
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon mdl-button--colored"> <i class="material-icons">favorite</i> </button></td>
          <td>favorite activities group</td>
        </tr>
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon mdl-button--colored"> <i class="material-icons">school</i> </button></td>
          <td>learning related group</td>
        </tr>
    
        <!--activities:::::::::::-->
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon"> <i class="material-icons">brightness_3</i> </button></td>
          <td>sleeping</td>
        </tr>
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon"> <i class="material-icons">directions_bike</i> </button></td>
          <td>biking</td>
        </tr>
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon"> <i class="material-icons">directions_run</i> </button></td>
          <td>running</td>
        </tr>
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon"> <i class="material-icons">directions_walk</i> </button></td>
          <td>walking</td>
        </tr>
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon"> <i class="material-icons">restaurant_menu</i> </button></td>
          <td>eating</td>
        </tr>
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon"> <i class="material-icons">local_grocery_store</i> </button></td>
          <td>shopping</td>
        </tr>
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon"> <i class="material-icons">local_play</i> </button></td>
          <td>watching movie</td>
        </tr>
        <tr>
          <td><button class="mdl-button mdl-js-button mdl-button--icon"> <i class="material-icons">whatshot</i> </button></td>
          <td>cooking</td>
        </tr>


      </tbody>
    </table>
  </body>
</html>
`

func handleIconLab(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, UsefulMaterialIcons)
}

func handleTestLab(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	var recSet []ActivityNameOnly
	parentKey := helpers.GetActivityTableKeyByUser(c)

	activeRecs := datastore.NewQuery("activityRecord").Ancestor(parentKey)

	//activeRecs = activeRecs.Filter("Status=", helpers.ActivityStatusStarted)
	activeRecs = activeRecs.Project("ActivityName") //.Distinct() // pulling distinct activity names only

	activeRecs = activeRecs.Filter("ActivityName =", "search me")

	//	activeRecs = activeRecs.Order(orderBy)
	t := activeRecs.Run(c)

	for {
		//var recSet activityRecord
		_, err := t.Next(&recSet)
		if err == datastore.Done {
			break
		}
		if err != nil {
			c.Errorf("Running query: %v", err)
			break
		}
		fmt.Fprintln(w, "..")
		fmt.Fprintln(w, recSet)
	}

	//	fmt.Fprintln(w, recSet)
	// recSet

}

/*
//Note: had to create a new function to retrieve distinct activity names based on search criteria.
//		this did not fit in the GetActivity function because of .Project() returns only single column result set
//		where as the non projection query returns all fields, hence mismatch and resulting in issue. hence seperated them.
//Note: this code is NOT working, not returning records and so I need to revisit.
func GetActivityNames(c appengine.Context, filters []Filter, orderBy string) []string { //[TODO: need to return error]
	parentKey := GetActivityTableKeyByUser(c)
	recSet := []string{}
	fmt.Println("GetActivityNames")
	activeRecs := datastore.NewQuery("activityRecord").Ancestor(parentKey)

	for _, f := range filters {
		activeRecs = activeRecs.Filter(f.Left, f.Right)
	}

	activeRecs = activeRecs.Project("ActivityName").Distinct() // pulling distinct activity names only

	activeRecs = activeRecs.Order(orderBy)
	activeRecs.GetAll(c, &recSet)

	return recSet

}*/
