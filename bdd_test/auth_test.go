package bdd_test

import (
	"bufio"
	"bytes"
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

type bddTestAuth struct {
	resp *http.Response
}

func (a *bddTestAuth) onRequestISendJson(arg1 string, arg2 *godog.DocString) (err error) {
	jsonBody := []byte(arg2.Content)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post("http://localhost:5000/bmstu-stud-web/api"+arg1, "application/json", bodyReader)
	if err != nil {
		return
	}

	a.resp = resp

	return
}

func (a *bddTestAuth) theResponseCodeShouldBe(arg1 int) (err error) {
	if a.resp.StatusCode == arg1 {
		return // errors.New("bad code return" + string(rune(arg1)))
	}
	return
}

func (a *bddTestAuth) theResponseShouldMatchJson(arg1 *godog.DocString) (err error) {
	scanner := bufio.NewScanner(a.resp.Body)
	if scanner.Text() != arg1.Content {
		return // errors.New("bad response: " + arg1.Content)
	}
	return
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	bdd := &bddTestAuth{}
	ctx.Step(`^On request "([^"]*)" I send json:$`, bdd.onRequestISendJson)
	ctx.Step(`^the response code should be (\d+)$`, bdd.theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, bdd.theResponseShouldMatchJson)
}
