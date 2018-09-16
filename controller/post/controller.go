package PostController

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"clean-arch/model/post"
)

var	posts = map[int]*Post.Post{}
var	seq   = 1

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, posts)
}

func Create(c echo.Context) error {
	data := &Post.Post{
		ID: seq,
	}
	if err := c.Bind(data); err != nil {
		return err
	}
	posts[data.ID] = data
	seq++
	return c.JSON(http.StatusCreated, data)
}

func Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, posts[id])
}

func Update(c echo.Context) error {
	data := new(Post.Post)
	if err := c.Bind(data); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	posts[id].Title = data.Title
	posts[id].Content = data.Content
	return c.JSON(http.StatusOK, posts[id])
}

func Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(posts, id)
	return c.NoContent(http.StatusNoContent)
}