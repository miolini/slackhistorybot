package hooks

import (
	"appengine"
	"appengine/search"
	"net/http"
)

func handleSearchDelete(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	getQuery := req.URL.Query()
	indexName := getQuery.Get("index")
	id := getQuery.Get("id")
	index, err := search.Open(indexName)
	if err != nil {
		jsonReplyMap(rw, http.StatusInternalServerError, "error", err.Error())
		return
	}
	err = index.Delete(c, id)
	if err != nil {
		jsonReplyMap(rw, http.StatusInternalServerError, "error", err.Error())
		return
	}
	jsonReplyMap(rw, http.StatusOK, "index", indexName, "id", id)
}
