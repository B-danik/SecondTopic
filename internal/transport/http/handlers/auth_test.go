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
			mockBehavior: func(s *mock_service.MockAuthorization, user todo.User) {
				s.EXPECT().Create(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: "{\"id\":1}\n",
		},
		{
			name:                "FAILED",
			inputBody:           ``,
			inputUser:           todo.User{},
			mockBehavior:        func(s *mock_service.MockAuthorization, user todo.User) {},
			expectedStatusCode:  http.StatusNotFound,
			expectedRequestBody: "",
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Manager{services}

			r := echo.New()
			r.POST("/sign-up", handler.SignUp)

			http := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(http, req)

			assert.Equal(t, http.Code, test.expectedStatusCode)
			assert.Equal(t, http.Body.String(), test.expectedRequestBody)
		})
	}
}
