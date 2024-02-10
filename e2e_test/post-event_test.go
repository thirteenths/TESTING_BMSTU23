package e2e_test

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

type CreateEventEnd2EndTestSuite struct {
	suite.Suite
}

func (suite *CreateEventEnd2EndTestSuite) BeforeAll(t provider.T) {
	err := goose.UpTo(db, "migrations", 20240117165259)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}
}

func (suite *CreateEventEnd2EndTestSuite) AfterAll(t provider.T) {
	/*err := goose.DownTo(db, "migrations", 20240117165259)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}*/
}

func TestEndToEnd_CreateEvent(t *testing.T) {
	suite.RunSuite(t, new(CreateEventEnd2EndTestSuite))
}

func (suite *CreateEventEnd2EndTestSuite) TestCreateEvent(t provider.T) {
	jsonBody := []byte(`{
		"name":"Big Stand Up",
		"description":"Big Stand Up — шоу с самым большим процентом смеющихся людей. Здесь только опытные комики и шутки, проверенные не одной сотней избирательных зрителей. Приходите убедиться в пятницу, субботу и воскресенье, если вам больше 18 лет и вы свободны в пятницу, субботу и воскресенье.",
		"date":"2024-02-02T15:11:57.456624Z"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post("http://localhost:5000/bmstu-stud-web/api/events/", "application/json", bodyReader)
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
