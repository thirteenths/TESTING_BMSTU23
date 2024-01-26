package e2e

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func Test_CreateEvent(t *testing.T) {
	// we will run an HTTP server locally to test the POST request
	url := "http://localhost:5000/bmstu-stud-web/api/events/"

	// create post body
	data := []byte(`{"name":"Женщины-комики","description":"Женщины-комики — шоу, которое покажет силу женского юмора. В нём участвуют только девушки и только с лучшим своим материалом. Три опытных комикессы, а также ведущая, расскажут качественные шутки. Как про мужчин, феминизм и психотерапию, так и про глобализацию, энтропию и многое другое. Берите подругу, друга, да всех берите и приходите. Мы докажем вам, что женщины умеют шутить обо всём. 18+","date":"2024-01-25T16:46:57.727879Z"}`)
	body := bytes.NewReader(data)

	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print the status code
	fmt.Println("Status:", resp.Status)
}
