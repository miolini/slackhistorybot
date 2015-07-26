package hooks

import (
	"appengine"
	"appengine/datastore"
	"appengine/search"
	"fmt"
	"net/http"
	"strings"
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

	// Disabled
	if false && msg.UserName != "slackbot" && strings.Contains(strings.ToLower(msg.Text), "пщ") {
		jsonReplyMap(rw, 200, "text", fmt.Sprintf("@%s: не ПЩ, а Go!\nhttps://pbs.twimg.com/media/B3-o2B4CMAANNH3.png:large", msg.UserName))
	} else {
		jsonReplyMap(rw, 200, "msg", msg, "key", key)
	}

}
