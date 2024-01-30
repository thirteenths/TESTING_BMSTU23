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
	"github.com/thirteenths/test-bmstu23/internal/infra/builder"
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
		mockResponse     []domain.Event
		expectedResponse *responses.GetAllEvent
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			expectedResponse: &responses.GetAllEvent{
				Event: []responses.Event{
					{
						ID:   1,
						Name: "Big Stand Up",
						Description: "Big Stand Up — шоу с самым большим процентом смеющихся людей. Здесь только опытные комики и шутки, проверенные не одной сотней избирательных зрителей." +
							" Приходите убедиться в пятницу, субботу и воскресенье, если вам больше 18 лет и вы свободны в пятницу, субботу и воскресенье.",
						Date: time.Time{},
					},
					{
						ID:   2,
						Name: "Жёсткий стендап",
						Description: "Жёсткий стендап — это шоу, где комики могут шутить обо всем, о чем хотят, не боясь, что их сочтут сумасшедшими. " +
							"А зрители могут смеяться над всем, над чем хотят, не боясь, что это неуместно. Точно будут шутки про ХХX, не*******ю и к****c. " +
							"В шоу участвуют 4 комика и ведущий. Состав обновляется раз в месяц — успеете попасть на любимого стендапера. Приходите на шоу, чтобы посмеяться без стыда и, " +
							"возможно, расширить границы дозволенного. 18+",
						Date: time.Time{},
					},
				},
			},
			mockResponse: []domain.Event{
				*builder.EventMother{}.Obj1(),
				*builder.EventMother{}.Obj2(),
			},
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: nil,
			mockResponse:     nil,
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().GetAllEvent(ctx).Return(test.mockResponse, test.expectedError)

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
		mockResponse     domain.Event
		expectedResponse *responses.GetEvent
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			expectedResponse: &responses.GetEvent{
				ID:   1,
				Name: "Big Stand Up",
				Description: "Big Stand Up — шоу с самым большим процентом смеющихся людей. Здесь только опытные комики и шутки, проверенные не одной сотней избирательных зрителей." +
					" Приходите убедиться в пятницу, субботу и воскресенье, если вам больше 18 лет и вы свободны в пятницу, субботу и воскресенье.",
				Date: time.Time{},
			},
			mockResponse:  *builder.EventMother{}.Obj1(),
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: nil,
			mockResponse:     domain.Event{},
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().GetEvent(ctx, 1).Return(test.mockResponse, test.expectedError)

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
		nameTest      string
		request       int
		mockRequest   *requests.CreateEvent
		mockResponse  *domain.Event
		expectedError error
	}{
		1: {
			nameTest: "Test Ok",
			mockRequest: &requests.CreateEvent{
				Name: "Big Stand Up",
				Description: "Big Stand Up — шоу с самым большим процентом смеющихся людей. Здесь только опытные комики и шутки, проверенные не одной сотней избирательных зрителей." +
					" Приходите убедиться в пятницу, субботу и воскресенье, если вам больше 18 лет и вы свободны в пятницу, субботу и воскресенье.",
				Date: time.Time{},
			},
			mockResponse:  builder.EventMother{}.Obj1(),
			request:       1,
			expectedError: nil,
		},
		2: {
			nameTest:      "Test Error",
			mockRequest:   nil,
			mockResponse:  nil,
			request:       0,
			expectedError: errors.New("storage error")},
	}
	testCase[1].mockResponse.ID = 0
	for _, test := range testCase {

		// Mock the storage method call
		suite.mockStorage.EXPECT().CreateEvent(ctx, *test.mockResponse).Return(test.request, test.expectedError)

		// Call the service method
		actualResponse, actualError := suite.eventService.CreateEvent(ctx, *test.mockRequest)

		// Compare the expected and actual responses
		t.Require().Equal(test.expectedError, actualError)
		t.Require().Equal(test.request, actualResponse)
	}
}

func (suite *EventServiceTestSuite) TestUpdateEvent(t provider.T) {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest      string
		mockRequest   *requests.UpdateEvent
		mockResponse  *domain.Event
		expectedError error
	}{
		1: {
			nameTest: "Test Ok",
			mockRequest: &requests.UpdateEvent{
				Name: "Big Stand Up",
				Description: "Big Stand Up — шоу с самым большим процентом смеющихся людей. Здесь только опытные комики и шутки, проверенные не одной сотней избирательных зрителей." +
					" Приходите убедиться в пятницу, субботу и воскресенье, если вам больше 18 лет и вы свободны в пятницу, субботу и воскресенье.",
				Date: time.Time{},
			},
			mockResponse:  builder.EventMother{}.Obj1(),
			expectedError: nil,
		},
		2: {
			nameTest:      "Test Error",
			mockRequest:   nil,
			mockResponse:  nil,
			expectedError: errors.New("storage error")},
	}
	testCase[1].mockResponse.ID = 0
	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().UpdateEvent(ctx, *test.mockResponse, 1).Return(test.expectedError)

		// Call the service method
		actualError := suite.eventService.UpdateEvent(ctx, *test.mockRequest, 1)

		// Compare the expected and actual responses
		t.Require().Equal(test.expectedError, actualError)
	}
}

func (suite *EventServiceTestSuite) TestDeleteEvent(t provider.T) {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest      string
		expectedError error
	}{
		1: {
			nameTest:      "Test Ok",
			expectedError: nil,
		},
		2: {
			nameTest:      "Test Error",
			expectedError: errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().DeleteEvent(ctx, 1).Return(test.expectedError)

		// Call the service method
		actualError := suite.eventService.DeleteEvent(ctx, 1)

		// Compare the expected and actual responses
		t.Require().Equal(test.expectedError, actualError)
	}
}

func TestSuiteRunner_EventService(t *testing.T) {
	t.Setenv("ALLURE_OUTPUT_PATH", "../../")
	suite.RunSuite(t, new(EventServiceTestSuite))
}
