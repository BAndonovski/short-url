package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/BAndonovski/short-url/codec"
	"github.com/BAndonovski/short-url/data"
)

func returnUrl(w http.ResponseWriter, slp ShortLinkProtocol) {
	json, err := json.Marshal(slp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func encode(w http.ResponseWriter, r *http.Request) {
	p, err := protocolFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(p.Url) == 0 {
		http.Error(w, "Empty url", http.StatusBadRequest)
		return
	}

	short := fmt.Sprintf("%s/%s", os.Getenv("FINN_PREFIX"), codec.GenerateShort())
	err = data.Set(p.Url, short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	returnUrl(w, ShortLinkProtocol{
		Url: short,
	})
}

func decode(w http.ResponseWriter, r *http.Request) {
	short := r.URL.Query().Get("url")
	sl, err := data.Get(short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	returnUrl(w, ShortLinkProtocol{
		Url: sl.Original,
	})
}
