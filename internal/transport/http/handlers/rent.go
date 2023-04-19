package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/B-danik/SecondTopic/todo"
	"github.com/labstack/echo"
)

func (m *Manager) GetRent(c echo.Context) error {
	userid := c.Get("userId").(int)
	list, err := m.srv.Rent.GetRent(userid)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"rent": list,
	})
}

func (m *Manager) CreateRent(c echo.Context) error {
	var rent todo.Rent
	userid := c.Get("userId").(int)
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	err = json.Unmarshal(b, &rent)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	m.srv.Rent.CraeteRent(userid, rent.Book_id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusOK)
}
