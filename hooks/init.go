package hooks

import (
	"net/http"
)

func init() {
	http.HandleFunc("/slack_outgoing", handleSlackOutgoing)
	http.HandleFunc("/search_delete", handleSearchDelete)
}
