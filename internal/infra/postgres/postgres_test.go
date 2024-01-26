package postgres

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/sirupsen/logrus"
	"testing"
	"time"

	"github.com/thirteenths/test-bmstu23/internal/domain"
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

func (suite *EventRepositoryTestSuite) BeforeAll(t provider.T) {
	t.Setenv("ALLURE_OUTPUT_PATH", "../../")
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

func (suite *EventRepositoryTestSuite) TestGetAllEvent(t provider.T) {
	testCase := map[int]struct {
		nameTest         string
		column           []string
		expectedResponse domain.Event
		expectedError    error
		Query            string
		actualResponse   []domain.Event
		actualError      error
	}{
		1: {
			nameTest:         "Test Ok",
			column:           []string{"id", "name", "description", "date"},
			expectedResponse: domain.Event{ID: 1, Name: "name", Description: "description"},
			expectedError:    nil,
			Query:            getAllEventQuery,
			actualResponse:   []domain.Event{{ID: 1, Name: "name", Description: "description", Date: time.Time{}}},
			actualError:      nil,
		},
		2: {
			nameTest:         "Test Ok empty response",
			column:           []string{"id", "name", "description", "date"},
			expectedResponse: domain.Event{},
			expectedError:    nil,
			Query:            getAllEventQuery,
			actualResponse:   []domain.Event{{}},
			actualError:      nil,
		},
		3: {
			nameTest:         "Test Error",
			column:           []string{"id", "name", "description", "date"},
			expectedResponse: domain.Event{ID: 1, Name: "name", Description: "description"},
			expectedError:    errors.New("error postgres"),
			Query:            getAllEventQuery,
			actualResponse:   []domain.Event{},
			actualError:      errors.New("error postgres"),
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectQuery(test.Query).
			WillReturnRows(sqlmock.NewRows(test.column).
				AddRow(
					test.expectedResponse.ID,
					test.expectedResponse.Name,
					test.expectedResponse.Description,
					test.expectedResponse.Date,
				))

		resp, err := suite.repository.GetAllEvent()

		t.Require().Nil(err)
		t.Require().NotNil(resp)
	}
}

func (suite *EventRepositoryTestSuite) TestGetEvent(t provider.T) {
	testCase := map[int]struct {
		nameTest         string
		column           []string
		expectedResponse domain.Event
		expectedError    error
		Query            string
		actualResponse   domain.Event
		actualError      error
	}{
		1: {
			nameTest:         "Test Ok",
			column:           []string{"id", "name", "description", "date"},
			expectedResponse: domain.Event{ID: 1, Name: "name", Description: "description"},
			expectedError:    nil,
			Query:            "SELECT ID, NAME, DESCRIPTION, DATE FROM EVENTS",
			actualResponse:   domain.Event{ID: 1, Name: "name", Description: "description", Date: time.Time{}},
			actualError:      nil,
		},
		2: {
			nameTest:         "Test Ok empty response",
			column:           []string{"id", "name", "description", "date"},
			expectedResponse: domain.Event{},
			expectedError:    nil,
			Query:            "SELECT ID, NAME, DESCRIPTION, DATE FROM EVENTS",
			actualResponse:   domain.Event{},
			actualError:      nil,
		},
		3: {
			nameTest:         "Test Error",
			column:           []string{"id", "name", "description", "date"},
			expectedResponse: domain.Event{ID: 1, Name: "name", Description: "description"},
			expectedError:    errors.New("error postgres"),
			Query:            "SELECT ID, NAME, DESCRIPTION, DATE FROM EVENTS",
			actualResponse:   domain.Event{},
			actualError:      errors.New("error postgres"),
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectQuery(test.Query).
			WillReturnRows(sqlmock.NewRows(test.column).
				AddRow(
					test.expectedResponse.ID,
					test.expectedResponse.Name,
					test.expectedResponse.Description,
					test.expectedResponse.Date,
				))

		resp, err := suite.repository.GetEvent(1)

		t.Require().Nil(err)
		t.Require().NotNil(resp)
	}
}

func (suite *EventRepositoryTestSuite) TestCreateEvent(t provider.T) {
	testCase := map[int]struct {
		nameTest         string
		column           []string
		expectedResponse domain.Event
		expectedError    error
		Query            string
		actualResponse   domain.Event
		actualError      error
	}{
		1: {
			nameTest:         "Test Ok",
			column:           []string{"id"},
			expectedResponse: domain.Event{ID: 1, Name: "name", Description: "description"},
			expectedError:    nil,
			Query:            "INSERT INTO EVENTS",
			actualResponse:   domain.Event{ID: 1, Name: "name", Description: "description", Date: time.Time{}},
			actualError:      nil,
		},
		2: {
			nameTest:         "Test Error",
			column:           []string{"id"},
			expectedResponse: domain.Event{ID: 1, Name: "name", Description: "description"},
			expectedError:    errors.New("error postgres"),
			Query:            "INSERT INTO EVENTS",
			actualResponse:   domain.Event{},
			actualError:      errors.New("error postgres"),
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectQuery(test.Query).WithArgs(
			test.expectedResponse.Name,
			test.expectedResponse.Description,
			test.expectedResponse.Date,
		).
			WillReturnRows(sqlmock.NewRows(test.column).
				AddRow(
					test.expectedResponse.ID,
				))

		resp, err := suite.repository.CreateEvent(test.expectedResponse)

		t.Require().Nil(err)
		t.Require().NotNil(resp)
	}
}

func (suite *EventRepositoryTestSuite) TestUpdateEvent(t provider.T) {
	testCase := map[int]struct {
		nameTest         string
		column           []string
		expectedResponse domain.Event
		expectedError    error
		Query            string
		actualResponse   domain.Event
		actualError      error
	}{
		1: {
			nameTest:         "Test Ok",
			expectedResponse: domain.Event{ID: 1, Name: "name", Description: "description"},
			expectedError:    nil,
			Query:            "UPDATE EVENTS",
			actualResponse:   domain.Event{ID: 1, Name: "name", Description: "description", Date: time.Time{}},
			actualError:      nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: domain.Event{ID: 1, Name: "name", Description: "description"},
			expectedError:    errors.New("error postgres"),
			Query:            "UPDATE EVENTS",
			actualResponse:   domain.Event{},
			actualError:      errors.New("error postgres"),
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectExec(test.Query).WithArgs(
			test.expectedResponse.Name,
			test.expectedResponse.Description,
			test.expectedResponse.Date,

			test.expectedResponse.ID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		err := suite.repository.UpdateEvent(test.expectedResponse, 1)

		t.Require().Nil(err)
	}
}

func (suite *EventRepositoryTestSuite) TestDeleteEvent(t provider.T) {
	testCase := map[int]struct {
		nameTest         string
		column           []string
		expectedResponse domain.Event
		expectedError    error
		Query            string
		actualResponse   domain.Event
		actualError      error
	}{
		1: {
			nameTest:         "Test Ok",
			expectedResponse: domain.Event{ID: 1, Name: "name", Description: "description"},
			expectedError:    nil,
			Query:            "DELETE FROM EVENTS",
			actualResponse:   domain.Event{ID: 1, Name: "name", Description: "description", Date: time.Time{}},
			actualError:      nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: domain.Event{ID: 1, Name: "name", Description: "description"},
			expectedError:    errors.New("error postgres"),
			Query:            "DELETE FROM EVENTS",
			actualResponse:   domain.Event{},
			actualError:      errors.New("error postgres"),
		},
	}

	for _, test := range testCase {
		suite.mock.ExpectExec(test.Query).WithArgs(
			test.expectedResponse.ID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		err := suite.repository.DeleteEvent(1)

		t.Require().Nil(err)
	}
}

func TestSuiteRunner_Event(t *testing.T) {
	t.Setenv("ALLURE_OUTPUT_PATH", "../../../")
	suite.RunSuite(t, new(EventRepositoryTestSuite))
}
