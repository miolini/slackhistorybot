package viewer

import (
	"net/http"
)

func handleIndex(rw http.ResponseWriter, req *http.Request) {
	jsonReply(rw, 200, "Hello World")
}
