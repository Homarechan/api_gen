// Code generated by server_generator. DO NOT EDIT.
// generated version: devel

package sample

import (
	"io"
	"log"
	"net/http"

	"github.com/go-generalize/api_gen/server_generator/sample/apigen/props"
	"github.com/go-generalize/api_gen/server_generator/sample/apigen/wrapper"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
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
	router.GET("", r.Get(p))
	router.POST("create_table", r.PostCreateTable(p))
	router.POST("create_user", r.PostCreateUser(p))
	return r
}

// Get ...
func (r *Routes) Get(p *props.ControllerProps) echo.HandlerFunc {
	i := NewGetController(p)

	b, ok := (interface{})(i).(interface{ AutoBind() bool })
	bindable := !ok || b.AutoBind()

	return func(c echo.Context) error {
		var (
			req  *GetRequest
			werr *wrapper.APIError
		)

		if bindable {
			req = new(GetRequest)
			if err := c.Bind(req); err != nil {
				log.Printf("failed to JSON binding(/): %+v", err)
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return c.JSON(werr.Status, werr.Body)
				}
				return err
			}
		}
		res, err := i.Get(c, req)
		if err != nil {
			if xerrors.As(err, &werr) {
				log.Printf("%+v", werr)
				return c.JSON(werr.Status, werr.Body)
			}
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// PostCreateTable ...
func (r *Routes) PostCreateTable(p *props.ControllerProps) echo.HandlerFunc {
	i := NewPostCreateTableController(p)

	b, ok := (interface{})(i).(interface{ AutoBind() bool })
	bindable := !ok || b.AutoBind()

	return func(c echo.Context) error {
		var (
			req  *PostCreateTableRequest
			werr *wrapper.APIError
		)

		if bindable {
			req = new(PostCreateTableRequest)
			if err := c.Bind(req); err != nil {
				log.Printf("failed to JSON binding(/create_table): %+v", err)
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return c.JSON(werr.Status, werr.Body)
				}
				return err
			}
		}
		res, err := i.PostCreateTable(c, req)
		if err != nil {
			if xerrors.As(err, &werr) {
				log.Printf("%+v", werr)
				return c.JSON(werr.Status, werr.Body)
			}
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

	b, ok := (interface{})(i).(interface{ AutoBind() bool })
	bindable := !ok || b.AutoBind()

	return func(c echo.Context) error {
		var (
			req  *PostCreateUserRequest
			werr *wrapper.APIError
		)

		if bindable {
			req = new(PostCreateUserRequest)
			if err := c.Bind(req); err != nil {
				log.Printf("failed to JSON binding(/create_user): %+v", err)
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return c.JSON(werr.Status, werr.Body)
				}
				return err
			}
		}
		res, err := i.PostCreateUser(c, req)
		if err != nil {
			if xerrors.As(err, &werr) {
				log.Printf("%+v", werr)
				return c.JSON(werr.Status, werr.Body)
			}
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// IGetController ...
type IGetController interface {
	Get(c echo.Context, req *GetRequest) (res *GetResponse, err error)
}

// IPostCreateTableController ...
type IPostCreateTableController interface {
	PostCreateTable(c echo.Context, req *PostCreateTableRequest) (res *PostCreateTableResponse, err error)
}

// IPostCreateUserController ...
type IPostCreateUserController interface {
	PostCreateUser(c echo.Context, req *PostCreateUserRequest) (res *PostCreateUserResponse, err error)
}
