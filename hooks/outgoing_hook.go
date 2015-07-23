package hooks

import (
	"appengine"
	"appengine/datastore"
	"appengine/search"
	"fmt"
	"net/http"
)

func handleSlackOutgoing(rw http.ResponseWriter, req *http.Request) {
	var err error
	c := appengine.NewContext(req)
	msg := Message{}
	err = msg.BindRequest(req)
	if err != nil {
		jsonReplyMap(rw, http.StatusInternalServerError, "error", err.Error())
		return
	}
	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, msg.EntityKind(), nil), &msg)
	if err != nil {
		jsonReplyMap(rw, http.StatusInternalServerError, "error", err.Error())
		return
	}
	messages, err := search.Open("Messages")
	if err != nil {
		jsonReplyMap(rw, http.StatusInternalServerError, "error", err.Error())
		return
	}
	msgSearch := MessageSearchFromMessage(msg)
	_, err = messages.Put(c, fmt.Sprintf("%d", key.IntID()), &msgSearch)
	if err != nil {
		jsonReplyMap(rw, http.StatusInternalServerError, "error", err.Error())
		return
	}
	jsonReplyMap(rw, 200, "msg", msg, "key", key)
}
