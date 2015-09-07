package modl

import (
	"appengine"
	"appengine/search"
	"net/http"
	"strings"
)

//[TODO: need to document this function]
func FullTextSearchActivity(c appengine.Context, src string, w http.ResponseWriter) ([]ActivityNameOnly, error) { //TODO: w parameter only for debugging
	i, _ := search.Open("ActivityLog")

	t := i.Search(c, `ActivityName:`+src, nil)
	dst := make([]ActivityNameOnly, 10) //TODO: hardcoded to 10 for now. need to figureout alternative to dynamically get the size of the search results
	for i := 0; ; i++ {                 //TODO: is there a better way to handle the loop for search output? need to research
		_, err := t.Next(&dst[i])
		if err == search.Done {
			break
		} else if err != nil {
			return nil, err
		}
	}
	//fmt.Fprintln(w, "FullTextSearchActivity", dst)
	return dst, nil

}

// AddActivityToIndex function will add activity to full text search index. the way it works is the index would have an ID (using the simillar string as activity name to maintain avoiding duplicates entries. and the index fields include just the activity name, as we need full text search on just the activity name)
// When retrieving, the full text search on this index returns all the matching activity names, which inturn are used in the datasource query filter to get the complete activity records if required.
func ActivityAddToIndex(c appengine.Context, rec ActivityLog, w http.ResponseWriter) { //TODO: w is only for debug purpose, will be removed eventually
	i, err := search.Open("ActivityLog")
	if err != nil {
		panic(err.Error())
	}
	id := strings.Replace(rec.ActivityName, " ", "_", -1) // convert spaces to underscore as appengine/search.Index dont allow space in the index ID
	a := ActivityNameOnly{rec.ActivityName}
	_, err = i.Put(c, id, &a)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Fprintln(w, "Index Index")
	//fmt.Fprintln(w, i, s)

}
