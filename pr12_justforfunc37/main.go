package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/crs-d/go-github/github"
)

func main() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	var data github.PullRequestEvent
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		logrus.Errorf("could not decode req: %v", err)
		http.Error(w, "could not decode req", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "pull req: %d", *data.PullRequest.ID)
}