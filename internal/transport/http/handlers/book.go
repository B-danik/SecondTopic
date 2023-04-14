package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/B-danik/SecondTopic/todo"
	"github.com/labstack/echo"
)

func (m *Manager) CreateBook(c echo.Context) error {
	check, err := m.UserIdentity(c)
	if err != nil && check != true {
		return err
	}

	var book todo.Book
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	err = json.Unmarshal(b, &book)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	m.srv.Book.CreateBook(book.Name)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	i := c.Get("userId").(int)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": i,
	})
}
