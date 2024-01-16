package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
)

func TestPageWithCounter_ServeHTTP(t *testing.T) {
	counter := PageWithCounter{heading: "title", content: "some content"}
	r		:= httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	w		:= httptest.NewRecorder()

	counter.ServeHTTP(w,r)
	if counter.counter != 1 {
		t.Errorf("We expected 1 view but we received %d", counter.counter)
	}
	msg := fmt.Sprintf("<h1>%s</h1>\n<p>%s</p>\n<p>Views: %d</p>", counter.heading, counter.content, 1)
	if w.Body.String() != msg {
		t.Errorf("1.msg: Expected %s but received %s", msg, w.Body.String())
	}

	w = httptest.NewRecorder()
	counter.ServeHTTP(w,r)
	if counter.counter != 2 {
		t.Errorf("We expected 1 view but we received %d", counter.counter)
	}

	msg = fmt.Sprintf("<h1>%s</h1>\n<p>%s</p>\n<p>Views: %d</p>", counter.heading, counter.content, 2)
	if w.Body.String() != msg {
		t.Errorf("2.msg: Expected  '%s' but received %s", msg, w.Body.String())
	}
}