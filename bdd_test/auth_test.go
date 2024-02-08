package bdd_test

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"net/http"
	"testing"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func StepDefinitioninition1(ctx context.Context) error {
	jsonBody := []byte(`{
  		"email": "rachelle.huel@ethereal.email",
  		"password": "C6s2S9qe6WrTMB7z3u"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post("http://localhost:5000/bmstu-stud-web/api/users/login", "application/json", bodyReader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	return nil
}

func StepDefinitioninition2(ctx context.Context) error {
	jsonBody := []byte(`{
  		"email": "rachelle.huel@ethereal.email",
  		"code": "23456789"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post("http://localhost:5000/bmstu-stud-web/api/users/verify", "application/json", bodyReader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	return nil
}

func iSendPOSTRequestTo(arg1 string) error {
	return godog.ErrPending
}

func theResponseCodeShouldBe(arg1 int) error {
	return godog.ErrPending
}

func theResponseShouldMatchJson(arg1 *godog.DocString) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I send POST request to "([^"]*)"$`, iSendPOSTRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, theResponseShouldMatchJson)
}
