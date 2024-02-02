package e2e_test

import (
	"bufio"
	"fmt"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestGetEvent(t *testing.T) {
	err := goose.UpTo(db, "migrations", 20240117165259)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}

	resp, err := http.Get("http://localhost:5000/bmstu-stud-web/api/events/")
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

	err = goose.DownTo(db, "migrations", 20240117165259)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}
}
