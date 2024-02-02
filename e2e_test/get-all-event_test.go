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

type GetAllEventEnd2EndTestSuite struct {
	suite.Suite
}

func (suite *GetAllEventEnd2EndTestSuite) BeforeAll(t provider.T) {
	err := goose.UpTo(db, "migrations", 20240117165259)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}
}

func (suite *GetAllEventEnd2EndTestSuite) AfterAll(t provider.T) {
	err := goose.DownTo(db, "migrations", 20240117165259)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}
}

func TestEndToEnd_GetAllEvent(t *testing.T) {
	suite.RunSuite(t, new(GetAllEventEnd2EndTestSuite))
}

func (suite *GetAllEventEnd2EndTestSuite) TestGetAllEvent(t provider.T) {
	resp, err := http.Get("http://localhost:5000/bmstu-stud-web/api/events/")
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
