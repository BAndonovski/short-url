package api

import (
	"encoding/json"
	"net/http"
)

type ShortLinkProtocol struct {
	Url string `json:"url"`
}

func protocolFromRequest(r *http.Request) (ShortLinkProtocol, error) {
	var p ShortLinkProtocol
	err := json.NewDecoder(r.Body).Decode(&p)
	return p, err
}
