// Code generated by server_generator. DO NOT EDIT.
// generated version: unknown

package sample

import (
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	props "github.com/go-generalize/api_gen/server_generator/sample/props"
)

// Routes ...
type Routes struct {
	router *echo.Group
}

// NewRoutes ...
func NewRoutes(p *props.ControllerProps, router *echo.Group, opts ...io.Writer) *Routes {
	if len(opts) > 0 {
		if w := opts[0]; w != nil {
			log.SetOutput(w)
		}
	}
	r := &Routes{
		router: router,
	}
	router.POST("create_table", r.PostCreateTable(p))
	router.POST("create_user", r.PostCreateUser(p))
	return r
}

// PostCreateTable ...
func (r *Routes) PostCreateTable(p *props.ControllerProps) echo.HandlerFunc {
	i := NewPostCreateTableController(p)
	return func(c echo.Context) error {
		req := new(PostCreateTableRequest)
		if err := c.Bind(req); err != nil {
			log.Printf("failed to JSON binding(/create_table): %+v", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PostCreateTable(c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// PostCreateUser ...
func (r *Routes) PostCreateUser(p *props.ControllerProps) echo.HandlerFunc {
	i := NewPostCreateUserController(p)
	return func(c echo.Context) error {
		req := new(PostCreateUserRequest)
		if err := c.Bind(req); err != nil {
			log.Printf("failed to JSON binding(/create_user): %+v", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PostCreateUser(c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// IPostCreateTableController ...
type IPostCreateTableController interface {
	PostCreateTable(c echo.Context, req *PostCreateTableRequest) (res *PostCreateTableResponse, err error)
}

// IPostCreateUserController ...
type IPostCreateUserController interface {
	PostCreateUser(c echo.Context, req *PostCreateUserRequest) (res *PostCreateUserResponse, err error)
}
