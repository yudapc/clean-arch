package UserController

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"clean-arch/model/user"
)


var rootResponse = map[string]string {
	"app_name": "clean-arch",
	"status": "ok",
}
var	users = map[int]*User.User{}
var	seq   = 1


func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, rootResponse)
}

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func Create(c echo.Context) error {
	u := &User.User{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func Update(c echo.Context) error {
	u := new(User.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}