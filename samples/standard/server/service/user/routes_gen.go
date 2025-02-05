// Code generated by server_generator. DO NOT EDIT.
// generated version: devel

package user

import (
	"io"
	"log"
	"net/http"

	"github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/go-generalize/api_gen/samples/standard/server/wrapper"
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
	router.POST("update_user_name", r.PostUpdateUserName(p))
	router.POST("update_user_password", r.PostUpdateUserPassword(p))
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
				log.Printf("failed to JSON binding(/service/user): %+v", err)
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

// PostUpdateUserName ...
func (r *Routes) PostUpdateUserName(p *props.ControllerProps) echo.HandlerFunc {
	i := NewPostUpdateUserNameController(p)

	b, ok := (interface{})(i).(interface{ AutoBind() bool })
	bindable := !ok || b.AutoBind()

	return func(c echo.Context) error {
		var (
			req  *PostUpdateUserNameRequest
			werr *wrapper.APIError
		)

		if bindable {
			req = new(PostUpdateUserNameRequest)
			if err := c.Bind(req); err != nil {
				log.Printf("failed to JSON binding(/service/user/update_user_name): %+v", err)
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
		res, err := i.PostUpdateUserName(c, req)
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

// PostUpdateUserPassword ...
func (r *Routes) PostUpdateUserPassword(p *props.ControllerProps) echo.HandlerFunc {
	i := NewPostUpdateUserPasswordController(p)

	b, ok := (interface{})(i).(interface{ AutoBind() bool })
	bindable := !ok || b.AutoBind()

	return func(c echo.Context) error {
		var (
			req  *PostUpdateUserPasswordRequest
			werr *wrapper.APIError
		)

		if bindable {
			req = new(PostUpdateUserPasswordRequest)
			if err := c.Bind(req); err != nil {
				log.Printf("failed to JSON binding(/service/user/update_user_password): %+v", err)
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
		res, err := i.PostUpdateUserPassword(c, req)
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

// IPostUpdateUserNameController ...
type IPostUpdateUserNameController interface {
	PostUpdateUserName(c echo.Context, req *PostUpdateUserNameRequest) (res *PostUpdateUserNameResponse, err error)
}

// IPostUpdateUserPasswordController ...
type IPostUpdateUserPasswordController interface {
	PostUpdateUserPassword(c echo.Context, req *PostUpdateUserPasswordRequest) (res *PostUpdateUserPasswordResponse, err error)
}
