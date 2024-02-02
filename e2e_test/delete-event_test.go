package e2e_test

import (
	"bufio"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

type DeleteEventEnd2EndTestSuite struct {
	suite.Suite
}

func (suite *DeleteEventEnd2EndTestSuite) BeforeAll(t provider.T) {
	err := goose.UpTo(db, "migrations", 20240117165259)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}
}

func (suite *DeleteEventEnd2EndTestSuite) AfterAll(t provider.T) {
	err := goose.DownTo(db, "migrations", 20240117165259)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}
}

func TestEndToEnd_DeleteEvent(t *testing.T) {
	suite.RunSuite(t, new(DeleteEventEnd2EndTestSuite))
}

func (suite *DeleteEventEnd2EndTestSuite) TestDeleteEvent(t provider.T) {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "http://localhost:5000/bmstu-stud-web/api/events/1", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	t.Require().NotNil(resp.Body)
	t.Require().Equal(resp.StatusCode, 200)
}
