package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_name(t *testing.T) {
	handler, err := NewHello("./index.html")
	if err != nil {t.Error(err)}
	srv := httptest.NewServer(handler)
	res, err := http.Get(srv.URL+"/?name=joy")
	if err != nil {t.Error(err)}

	expected, err := ioutil.ReadFile("./teststatics/joy.html")
	if err != nil { t.Error(err)}

	actual := make([]byte, res.ContentLength)
	res.Body.Read(actual)
	if string(actual) != string(expected) {
		t.Errorf("\n%s\n%s", string(expected), string(actual))
	}
}