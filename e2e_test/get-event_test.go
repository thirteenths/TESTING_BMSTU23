package e2e_test

import (
	"bufio"
	"fmt"
	"net/http"
	"testing"
)

func TestGetEvent(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/rush-stand-up-club/api/events/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
