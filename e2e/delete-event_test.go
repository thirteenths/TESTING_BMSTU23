package e2e

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"testing"
)

func Test_DeleteEvent(t *testing.T) {
	// declare the request url and body
	url := "http://localhost:5000/bmstu-stud-web/api/events/2"
	body := strings.NewReader("This is the request body.")

	// we can set a custom method here, like http.MethodPut
	// or http.MethodDelete, http.MethodPatch, etc.
	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print the status code
	fmt.Println("Status:", resp.Status)
}
