package app

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/sirupsen/logrus"

	"reflect"
	"testing"
	"time"

	"github.com/thirteenths/test-bmstu23/internal/domain"
	"github.com/thirteenths/test-bmstu23/internal/domain/requests"
	"github.com/thirteenths/test-bmstu23/internal/domain/responses"
	mock_storage "github.com/thirteenths/test-bmstu23/internal/infra/mock"
)

type EventServiceTestSuite struct {
	suite.Suite
	ctrl         *gomock.Controller
	mockStorage  *mock_storage.MockeventServiceStorage
	eventService *EventService
}

func (suite *EventServiceTestSuite) BeforeEach(t provider.T) {
	t.Epic("Demo")
	t.Feature("BeforeAfter")
	t.NewStep("This Step will be before Each")

	suite.ctrl = gomock.NewController(t)
	suite.mockStorage = mock_storage.NewMockeventServiceStorage(suite.ctrl)
	suite.eventService = NewEventService(logrus.New(), suite.mockStorage)
}

func (suite *EventServiceTestSuite) AfterEach(t provider.T) {
	t.NewStep("AfterEach Step")

	suite.ctrl.Finish()
}

func (suite *EventServiceTestSuite) TestGetAllEvent(t provider.T) {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest         string
		request          []domain.Event
		expectedResponse *responses.GetAllEvent
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			expectedResponse: &responses.GetAllEvent{
				Event: []responses.Event{
					{
						ID:          1,
						Name:        "testAll",
						Description: "testAbout",
						Date:        time.Time{},
					},
				},
			},
			request: []domain.Event{
				{
					ID:          1,
					Name:        "testAll",
					Description: "testAbout",
					Date:        time.Time{},
				}},
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: nil,
			request:          []domain.Event{},
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().GetAllEvent(ctx).Return(test.request, test.expectedError)

		// Call the service method
		actualResponse, actualError := suite.eventService.GetAllEvent(ctx)

		// Compare the expected and actual responses
		t.Require().Equal(test.expectedError, actualError)
		t.Require().True(reflect.DeepEqual(test.expectedResponse, actualResponse))
	}
}

func (suite *EventServiceTestSuite) TestGetEvent(t provider.T) {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest         string
		request          domain.Event
		expectedResponse *responses.GetEvent
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			expectedResponse: &responses.GetEvent{
				ID:          1,
				Name:        "testAll",
				Description: "testAbout",
				Date:        time.Time{},
			},
			request: domain.Event{
				ID:          1,
				Name:        "testAll",
				Description: "testAbout",
				Date:        time.Time{},
			},
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: nil,
			request:          domain.Event{},
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().GetEvent(ctx, 1).Return(test.request, test.expectedError)

		// Call the service method
		actualResponse, actualError := suite.eventService.GetEvent(ctx, 1)

		// Compare the expected and actual responses
		t.Require().Equal(test.expectedError, actualError)
		t.Require().True(reflect.DeepEqual(test.expectedResponse, actualResponse))
	}
}

func (suite *EventServiceTestSuite) TestCreateEvent(t provider.T) {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest         string
		request          int
		expectedResponse *requests.CreateEvent
		expect           *domain.Event
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			expectedResponse: &requests.CreateEvent{
				Name:        "testAll",
				Description: "testAbout",
				Date:        time.Time{},
			},
			expect: &domain.Event{
				ID:          0,
				Name:        "testAll",
				Description: "testAbout",
				Date:        time.Time{},
			},
			request:       1,
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: nil,
			expect:           nil,
			request:          0,
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().CreateEvent(ctx, *test.expect).Return(test.request, test.expectedError)

		// Call the service method
		actualResponse, actualError := suite.eventService.CreateEvent(ctx, *test.expectedResponse)

		// Compare the expected and actual responses
		t.Require().Equal(test.expectedError, actualError)
		t.Require().Equal(test.request, actualResponse)
	}
}

func (suite *EventServiceTestSuite) TestUpdateEvent(t provider.T) {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest         string
		request          int
		expectedResponse *requests.UpdateEvent
		expect           *domain.Event
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			expectedResponse: &requests.UpdateEvent{
				Name:        "testAll",
				Description: "testAbout",
				Date:        time.Time{},
			},
			expect: &domain.Event{
				ID:          0,
				Name:        "testAll",
				Description: "testAbout",
				Date:        time.Time{},
			},
			request:       1,
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: nil,
			expect:           nil,
			request:          0,
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().UpdateEvent(ctx, *test.expect, 1).Return(test.expectedError)

		// Call the service method
		actualError := suite.eventService.UpdateEvent(ctx, *test.expectedResponse, 1)

		// Compare the expected and actual responses
		t.Require().Equal(test.expectedError, actualError)
		// t.Require().True(reflect.DeepEqual(test.expectedResponse, actualResponse))
	}
}

func (suite *EventServiceTestSuite) TestDeleteEvent(t provider.T) {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest         string
		request          int
		expectedResponse *requests.CreateEvent
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			expectedResponse: &requests.CreateEvent{
				Name:        "testAll",
				Description: "testAbout",
				Date:        time.Time{},
			},
			request:       1,
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: nil,
			request:          0,
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().DeleteEvent(ctx, 1).Return(test.expectedError)

		// Call the service method
		actualError := suite.eventService.DeleteEvent(ctx, 1)

		// Compare the expected and actual responses
		t.Require().Equal(test.expectedError, actualError)
		// t.Require().True(reflect.DeepEqual(test.expectedResponse, actualResponse))
	}
}

func TestSuiteRunner_EventService(t *testing.T) {
	t.Setenv("ALLURE_OUTPUT_PATH", "../../")
	suite.RunSuite(t, new(EventServiceTestSuite))
}
