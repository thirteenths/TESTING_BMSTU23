package integration_test

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
	"testing"

	"github.com/thirteenths/test-bmstu23/internal/app"
	"github.com/thirteenths/test-bmstu23/internal/domain/storage"
	"github.com/thirteenths/test-bmstu23/internal/infra/postgres"
)

type GetEventIntegrationTestSuite struct {
	suite.Suite
	eventService *app.EventService
	ctx          context.Context
	testCase     map[int]struct {
		nameTest      string
		expectedError error
		version       int64
	}
}

func (suite *GetEventIntegrationTestSuite) BeforeAll(t provider.T) {
	// Storage
	storage := storage.NewStorage(*postgres.NewMockPostgres(db))

	logger := log.New()
	suite.eventService = app.NewEventService(logger, storage)

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	suite.ctx = context.Background()
	suite.testCase = map[int]struct {
		nameTest      string
		expectedError error
		version       int64
	}{
		1: {
			nameTest: "Test Error",
			version:  20231206192143,
		},
		2: {
			nameTest: "Test Ok",
			version:  20240117165259,
		},
	}

}

func (suite *GetEventIntegrationTestSuite) AfterEach(t provider.T) {
	// err := goose.DownTo(db, "migrations", 20231206192143)
	// if err != nil {
	//	log.Warnf("Error rollback migration: %s", err)
	// }
}

func (suite *GetEventIntegrationTestSuite) AfterAll(t provider.T) {}

func TestIntegration_GetEvent(t *testing.T) {
	suite.RunSuite(t, &GetEventIntegrationTestSuite{})
}

func (suite *GetEventIntegrationTestSuite) TestGetEvent_Ok(t provider.T) {
	// Migration
	if err := goose.UpTo(db, "migrations", suite.testCase[2].version); err != nil {
		log.Warnf("Error migration: %s", err)
	}
	// Call the service method
	actualResponse, actualError := suite.eventService.GetEvent(suite.ctx, 2)

	// Compare the expected and actual responses
	t.Require().Nil(actualError)
	t.Require().NotNil(actualResponse)
}

/*func (suite *GetEventIntegrationTestSuite) TestGetEvent_Error(t provider.T) {
	// Migration
	if err := goose.DownTo(db, "migrations", suite.testCase[1].version); err != nil {
		log.Warnf("Error migration: %s", err)
	}
	// Call the service method
	actualResponse, actualError := suite.eventService.GetEvent(suite.ctx, 1)
	// Compare the expected and actual responses
	t.Require().NotNil(actualError)
	t.Require().Nil(actualResponse)
}
*/
