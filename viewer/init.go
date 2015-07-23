package viewer

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
}
