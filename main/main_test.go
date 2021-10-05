package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/BAndonovski/short-url/api"
	"github.com/gorilla/mux"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	api.Router(mux.NewRouter()).ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func init() {
	os.Setenv("FINN_PREFIX", "localhost:8080")
}

var short string

func TestEncodeFailEmpty(t *testing.T) {
	var reqJson = []byte(`{"url":""}`)
	req, _ := http.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer((reqJson)))
	req.Header.Set("Content-Type", "application/json")

	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)

	var p api.ShortLinkProtocol
	json.Unmarshal(resp.Body.Bytes(), &p)
	short = p.Url
}

func TestEncodeOK(t *testing.T) {
	var reqJson = []byte(`{"url":"http://finn.auto"}`)
	req, _ := http.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer((reqJson)))
	req.Header.Set("Content-Type", "application/json")

	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	var p api.ShortLinkProtocol
	json.Unmarshal(resp.Body.Bytes(), &p)
	short = p.Url
}

func TestDecodeFail(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet,
		"/decode?url=FAIL",
		nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, resp.Code)
}

func TestDecodeOK(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet,
		fmt.Sprintf("/decode?url=%s", short),
		nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
	var p api.ShortLinkProtocol
	json.Unmarshal(resp.Body.Bytes(), &p)
	if p.Url != "http://finn.auto" {
		t.Errorf("Incorrect url returned, expected http://finn.auto, got %s", p.Url)
	}
}
