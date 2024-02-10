package postgres

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/sirupsen/logrus"
	"testing"

	"github.com/thirteenths/test-bmstu23/internal/domain"
	"github.com/thirteenths/test-bmstu23/internal/infra/builder"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logrus.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

type EventRepositoryTestSuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository *Postgres
}

func (suite *EventRepositoryTestSuite) BeforeEach(t provider.T) {
	t.Epic("Data Access Layer")
	t.Feature("BeforeEach")
	t.NewStep("This Step will be before Each")

	db, mock := NewMock()
	suite.mock = mock
	suite.repository = NewMockPostgres(db)
}

func (suite *EventRepositoryTestSuite) AfterEach(t provider.T) {
	t.NewStep("AfterEach Step")
}

func (suite *EventRepositoryTestSuite) TestGetAllEvent_OK(t provider.T) {
	column := []string{"id", "name", "description", "date"}
	query := "SELECT ID, NAME, DESCRIPTION, DATE FROM EVENTS"
	testCase := map[int]struct {
		nameTest       string
		mockResponse   domain.Event
		mockError      error
		expectResponse []domain.Event
		expectError    error
	}{
		1: {
			nameTest:       "Test Ok",
			mockResponse:   *builder.EventMother{}.Obj1(),
			mockError:      nil,
			expectResponse: []domain.Event{*builder.EventMother{}.Obj1()},
			expectError:    nil,
		},
		2: {
			nameTest:       "Test Ok empty response",
			mockResponse:   *builder.EventMother{}.Obj0(),
			mockError:      nil,
			expectResponse: []domain.Event{*builder.EventMother{}.Obj0()},
			expectError:    nil,
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectQuery(query).
			WillReturnRows(sqlmock.NewRows(column).
				AddRow(
					test.mockResponse.ID,
					test.mockResponse.Name,
					test.mockResponse.Description,
					test.mockResponse.Date,
				))

		resp, err := suite.repository.GetAllEvent()

		t.Require().Nil(err)
		t.Require().NotNil(resp)

		t.Require().Equal(test.expectResponse, resp)
	}
}

func (suite *EventRepositoryTestSuite) TestGetAllEvent_Error(t provider.T) {
	query := "SELECT ID, NAME, DESCRIPTION, DATE FROM EVENTS"
	testCase := map[int]struct {
		nameTest       string
		mockResponse   domain.Event
		mockError      error
		expectResponse []domain.Event
		expectError    error
	}{
		3: {
			nameTest:       "Test Error",
			mockResponse:   domain.Event{ID: 1, Name: "name", Description: "description"},
			mockError:      errors.New("error postgres"),
			expectResponse: nil,
			expectError:    errors.New("error postgres"),
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectQuery(query).WillReturnError(test.expectError)

		resp, err := suite.repository.GetAllEvent()

		t.Require().Nil(resp)
		t.Require().NotNil(err)
	}
}

func (suite *EventRepositoryTestSuite) TestGetEvent_OK(t provider.T) {
	column := []string{"id", "name", "description", "date"}
	query := "SELECT ID, NAME, DESCRIPTION, DATE FROM EVENTS"
	testCase := map[int]struct {
		nameTest       string
		mockResponse   domain.Event
		mockError      error
		expectResponse domain.Event
		expectError    error
	}{
		1: {
			nameTest:       "Test Ok",
			mockResponse:   *builder.EventMother{}.Obj1(),
			mockError:      nil,
			expectResponse: *builder.EventMother{}.Obj1(),
			expectError:    nil,
		},
		2: {
			nameTest:       "Test Ok empty response",
			mockResponse:   *builder.EventMother{}.Obj0(),
			mockError:      nil,
			expectResponse: *builder.EventMother{}.Obj0(),
			expectError:    nil,
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectQuery(query).
			WillReturnRows(sqlmock.NewRows(column).
				AddRow(
					test.mockResponse.ID,
					test.mockResponse.Name,
					test.mockResponse.Description,
					test.mockResponse.Date,
				))

		resp, err := suite.repository.GetEvent(1)

		t.Require().Nil(err)
		t.Require().NotNil(resp)

		t.Require().Equal(test.expectResponse, resp)
	}
}

func (suite *EventRepositoryTestSuite) TestGetEvent_Error(t provider.T) {
	query := "SELECT ID, NAME, DESCRIPTION, DATE FROM EVENTS"
	testCase := map[int]struct {
		nameTest       string
		mockResponse   domain.Event
		mockError      error
		expectResponse domain.Event
		expectError    error
	}{
		3: {
			nameTest:       "Test Error",
			mockResponse:   *builder.EventMother{}.Obj0(),
			mockError:      errors.New("error postgres"),
			expectResponse: *builder.EventMother{}.Obj0(),
			expectError:    errors.New("error postgres"),
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectQuery(query).
			WillReturnError(test.mockError)

		resp, err := suite.repository.GetEvent(1)

		t.Require().NotNil(err)

		t.Require().Equal(test.expectResponse, resp)
	}
}

func (suite *EventRepositoryTestSuite) TestCreateEvent_OK(t provider.T) {
	column := []string{"id"}
	query := "INSERT INTO EVENTS"
	testCase := map[int]struct {
		nameTest       string
		mockResponse   domain.Event
		mockError      error
		expectResponse int
		expectError    error
	}{
		1: {
			nameTest:       "Test Ok",
			mockResponse:   domain.Event{ID: 1, Name: "name", Description: "description"},
			mockError:      nil,
			expectResponse: 1,
			expectError:    nil,
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectQuery(query).WithArgs(
			test.mockResponse.Name,
			test.mockResponse.Description,
			test.mockResponse.Date,
		).
			WillReturnRows(sqlmock.NewRows(column).
				AddRow(
					test.mockResponse.ID,
				))

		resp, err := suite.repository.CreateEvent(test.mockResponse)

		t.Require().Nil(err)
		t.Require().NotNil(resp)

		t.Require().Equal(test.expectResponse, resp)
	}
}

func (suite *EventRepositoryTestSuite) TestCreateEvent_Error(t provider.T) {
	query := "INSERT INTO EVENTS"
	testCase := map[int]struct {
		nameTest       string
		mockResponse   domain.Event
		mockError      error
		expectResponse int
		expectError    error
	}{
		2: {
			nameTest:       "Test Error",
			mockResponse:   domain.Event{ID: 1, Name: "name", Description: "description"},
			mockError:      errors.New("error postgres"),
			expectResponse: 0,
			expectError:    errors.New("error postgres"),
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectQuery(query).WithArgs(
			test.mockResponse.Name,
			test.mockResponse.Description,
			test.mockResponse.Date,
		).
			WillReturnError(test.expectError)

		resp, err := suite.repository.CreateEvent(test.mockResponse)

		t.Require().NotNil(err)
		t.Require().Equal(test.expectResponse, resp)
	}
}

func (suite *EventRepositoryTestSuite) TestUpdateEvent_OK(t provider.T) {
	query := "UPDATE EVENTS"
	testCase := map[int]struct {
		nameTest    string
		mockArgs    domain.Event
		mockError   error
		expectError error
	}{
		1: {
			nameTest:    "Test Ok",
			mockArgs:    domain.Event{ID: 1, Name: "name", Description: "description"},
			mockError:   nil,
			expectError: nil,
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectExec(query).WithArgs(
			test.mockArgs.Name,
			test.mockArgs.Description,
			test.mockArgs.Date,

			test.mockArgs.ID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		err := suite.repository.UpdateEvent(test.mockArgs, 1)

		t.Require().Nil(err)
	}
}

func (suite *EventRepositoryTestSuite) TestUpdateEvent_Error(t provider.T) {
	query := "UPDATE EVENTS"
	testCase := map[int]struct {
		nameTest    string
		mockArgs    domain.Event
		mockError   error
		expectError error
	}{
		2: {
			nameTest:    "Test Error",
			mockArgs:    domain.Event{ID: 1, Name: "name", Description: "description"},
			mockError:   errors.New("error postgres"),
			expectError: errors.New("error postgres"),
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectExec(query).WithArgs(
			test.mockArgs.Name,
			test.mockArgs.Description,
			test.mockArgs.Date,

			test.mockArgs.ID,
		).WillReturnError(test.expectError)

		err := suite.repository.UpdateEvent(test.mockArgs, 1)

		t.Require().NotNil(err)
	}
}

func (suite *EventRepositoryTestSuite) TestDeleteEvent_OK(t provider.T) {
	query := "DELETE FROM EVENTS"
	testCase := map[int]struct {
		nameTest    string
		mockArgs    int
		mockError   error
		expectError error
	}{
		1: {
			nameTest:    "Test Ok",
			mockArgs:    1,
			mockError:   nil,
			expectError: nil,
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectExec(query).WithArgs(
			test.mockArgs,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		err := suite.repository.DeleteEvent(1)

		t.Require().Nil(err)
	}
}

func (suite *EventRepositoryTestSuite) TestDeleteEvent_Error(t provider.T) {
	query := "DELETE FROM EVENTS"
	testCase := map[int]struct {
		nameTest    string
		mockArgs    int
		mockError   error
		expectError error
	}{
		2: {
			nameTest:    "Test Error",
			mockArgs:    1,
			mockError:   errors.New("error postgres"),
			expectError: errors.New("error postgres"),
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectExec(query).WithArgs(
			test.mockArgs,
		).WillReturnError(test.expectError)

		err := suite.repository.DeleteEvent(1)

		t.Require().NotNil(err)
	}
}

func TestSuiteRunner_Event(t *testing.T) {
	t.Setenv("ALLURE_OUTPUT_PATH", "../../../")
	suite.RunSuite(t, new(EventRepositoryTestSuite))
}
