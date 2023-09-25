package service_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/counter/mock_counter"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/model"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/repository/mock_repository"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/service"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

func Test_ShortenLink(t *testing.T) {
	// Initialize the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create test cases
	testCases := []struct {
		name           string
		longLink       string
		expectedResult string
		expectedError  error
		doMockLinkRepo func(mocks *mock_repository.MockLinker)
		doMockCounter  func(mocks *mock_counter.MockCounter)
	}{
		{
			name:           "Success. Found the existing link in DB",
			longLink:       "https://google.com/this-is-a-very-long-link",
			expectedResult: "{\"short_link\":\"http://localhost:8080/1\"}\n",
			expectedError:  nil,
			doMockLinkRepo: func(mocks *mock_repository.MockLinker) {
				mocks.EXPECT().GetByLongLink("https://google.com/this-is-a-very-long-link").Return(&model.Link{ShortLink: "http://localhost:8080/1"}, nil).Times(1)
			},
			doMockCounter: func(mocks *mock_counter.MockCounter) {},
		},
		{
			name:           "Success. Not found the link in DB",
			longLink:       "https://google.com/this-is-a-very-long-link",
			expectedResult: "{\"short_link\":\"example.com/1\"}\n",
			expectedError:  nil,
			doMockLinkRepo: func(mocks *mock_repository.MockLinker) {
				mocks.EXPECT().GetByLongLink("https://google.com/this-is-a-very-long-link").Return(nil, nil).Times(1)
				mocks.EXPECT().Create(gomock.Any()).Return(nil, nil).Times(1)
			},
			doMockCounter: func(mocks *mock_counter.MockCounter) {
				mocks.EXPECT().GetOrderNumber().Return(1, nil).Times(1)
			},
		},
	}

	// Run test cases
	for _, tc := range testCases {
		// Create a mock for the repository
		mockLinkRepo := mock_repository.NewMockLinker(ctrl)
		tc.doMockLinkRepo(mockLinkRepo)

		// Create a mock for the counter
		mockCounter := mock_counter.NewMockCounter(ctrl)
		tc.doMockCounter(mockCounter)

		svc := service.New(mockLinkRepo, mockCounter, nil)
		res, c := setupRouter(svc)

		err := svc.ShortenLink(c)

		if err != tc.expectedError {
			t.Errorf("Test_ShortenLink: expected error is %v but got %v", tc.expectedError, err)
		}
		if res.Body.String() != tc.expectedResult {
			t.Errorf("Test_ShortenLink: expected result is %s but got %s", tc.expectedResult, res.Body.String())
		}

	}

}

func setupRouter(svc service.Service) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	e.POST("/shorten-link", svc.ShortenLink)

	payload := `{
		"long_link": "https://google.com/this-is-a-very-long-link"
	}`
	req := httptest.NewRequest(http.MethodPost, "/shorten-link", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, "application/json")

	rec := httptest.NewRecorder()
	return rec, e.NewContext(req, rec)
}
