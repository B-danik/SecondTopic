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

	m.srv.Book.Create(book.Name)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	id := c.Get("userId").(int)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (m *Manager) GetBook(c echo.Context) error {
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

	list, err := m.srv.Book.Get(book.Id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"list": list,
	})
}

func (m *Manager) GetAll(c echo.Context) error {
	check, err := m.UserIdentity(c)
	if err != nil && check != true {
		return err
	}

	list, err := m.srv.Book.GetAll()
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"list": list,
	})
}

func (m *Manager) Delete(c echo.Context) error {
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

	result := m.srv.Book.Delete(book.Id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": result,
	})
}
