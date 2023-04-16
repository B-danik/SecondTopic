package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/B-danik/SecondTopic/pkg/service"
	mock_service "github.com/B-danik/SecondTopic/pkg/service/mocks"
	"github.com/B-danik/SecondTopic/todo"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"

	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user todo.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           todo.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"Email": "xasxx@mail.ru",
				"Password": "123",
				"Name": "Daniyar",
				"Lastname": "Bay"
			}`,
			inputUser: todo.User{
				Name:     "Daniyar",
				LastName: "Bay",
				Email:    "xasxx@mail.ru",
				Password: "123",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user todo.User) {
				r.EXPECT().Create(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: "{\"id\":1}\n",
		},
		{
			name:      "FAILED",
			inputBody: ``,
			inputUser: todo.User{
				Name:     "Daniyar",
				LastName: "Bay",
				Email:    "xxx@mail.ru",
				Password: "123",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user todo.User) {
				r.EXPECT().Create(user).Return(0, nil)
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "FAILED",
			inputBody: `{
				"Email": "xasxx@mail.ru",
				"Password": "123",
				"Name": "Daniyar",
				"Lastname": "Bay"
			}`,
			inputUser: todo.User{},
			mockBehavior: func(r *mock_service.MockAuthorization, user todo.User) {
				r.EXPECT().Create(user).Return(0, nil)
			},
			expectedStatusCode: 404,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Manager{services}

			// Init Endpoint
			r := echo.New()
			r.POST("/sign-up", handler.SignUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedRequestBody)
		})
	}
}
