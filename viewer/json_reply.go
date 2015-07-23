package viewer

import (
	"encoding/json"
	"errors"
	"net/http"
)

func jsonReply(rw http.ResponseWriter, responseCode int, msg interface{}) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(responseCode)
	return json.NewEncoder(rw).Encode(msg)
}

func jsonReplyMap(rw http.ResponseWriter, responseCode int, args ...interface{}) error {
	m := make(map[string]interface{})
	if len(args) == 0 {
		return jsonReply(rw, responseCode, m)
	} else if len(args)%2 != 0 {
		return errors.New("args count should be even")
	}
	for i := 0; i < len(args); i += 2 {
		name, ok := args[i].(string)
		if !ok {
			return errors.New("current arg should be with string type")
		}
		m[name] = args[i+1]
	}
	return jsonReply(rw, responseCode, m)
}
