package handlers

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/B-danik/SecondTopic/pkg/service"
	mock_service "github.com/B-danik/SecondTopic/pkg/service/mocks"
	"github.com/B-danik/SecondTopic/todo"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	type mockBehavior func(s *mock_service.MockBook, book todo.Book)

	testTable := []struct {
		name                string
		inputBody           string
		inputBook           todo.Book
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		// {
		// 	name: "OK",
		// 	inputBody: `{
		// 		"Name": "Test"
		// 	}`,
		// 	inputBook: todo.Book{
		// 		Name: "Test",
		// 	},
		// 	mockBehavior: func(s *mock_service.MockBook, book todo.Book) {
		// 		s.EXPECT().Create(book.Name).Return(1, nil)
		// 	},
		// 	expectedStatusCode:  200,
		// 	expectedRequestBody: `"id": 1`,
		// },
		{
			name:                "FAILED",
			inputBody:           ``,
			inputBook:           todo.Book{},
			mockBehavior:        func(s *mock_service.MockBook, book todo.Book) {},
			expectedStatusCode:  404,
			expectedRequestBody: "",
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockBook(c)
			test.mockBehavior(repo, test.inputBook)

			services := &service.Service{Book: repo}
			handler := Manager{services}

			r := echo.New()
			r.POST("/create", handler.CreateBook)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedRequestBody)
		})
	}

}
