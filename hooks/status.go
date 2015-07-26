package hooks

import (
	"appengine"
	"appengine/datastore"
	"appengine/search"
	"net/http"
	// "time"
)

type StatusResponse struct {
	MessagesDatastore int64 `json:"messages_datastore"`
	MessageSearch     int64 `json:"messages_search"`
}

func handleStatus(rw http.ResponseWriter, req *http.Request) {
	// ts := time.Now()

	c := appengine.NewContext(req)

	// Count messages in Google Datastore
	q := datastore.NewQuery(Message{}.EntityKind())
	t := q.Run(c)
	var status StatusResponse
	for {
		var msg Message
		_, err := t.Next(&msg)
		if err == datastore.Done {
			break
		} else if err != nil {
			c.Errorf("fetching new Message: %v", err)
			break
		}
		status.MessagesDatastore++
	}

	// Count messages in Google AppEngine Search Index
	index, err := search.Open("Messages")
	if err != nil {
		jsonReplyMap(rw, http.StatusInternalServerError, "error", err.Error())
		return
	}
	for t := index.Search(c, "", nil); ; {
		_, err := t.Next(nil)
		if err == search.Done {
			break
		} else if err != nil {
			jsonReplyMap(rw, http.StatusInternalServerError, "error", err.Error())
			return
		}
		status.MessageSearch++
	}

	jsonReplyMap(rw, http.StatusOK, "message", status)
}
