package e2e_test

import (
	"bufio"
	"fmt"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"net/http"
)

type GetEventEnd2EndTestSuite struct {
	suite.Suite
}

func (suite *GetEventEnd2EndTestSuite) BeforeAll(t provider.T) {
	err := goose.UpTo(db, "migrations", 20240117165259)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}
}

func (suite *GetEventEnd2EndTestSuite) AfterAll(t provider.T) {
	/*err := goose.DownTo(db, "migrations", 20240117165259)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}*/
}

func TestEndToEnd_GetEvent(t *testing.T) {
	suite.RunSuite(t, new(GetEventEnd2EndTestSuite))
}

func (suite *GetEventEnd2EndTestSuite) TestGetEvent(t provider.T) {
	resp, err := http.Get("http://localhost:5000/bmstu-stud-web/api/events/1")
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
