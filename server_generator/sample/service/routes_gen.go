// Code generated by server_generator. DO NOT EDIT.
// generated version: devel

package service

import (
	"io"
	"log"
	"net/http"

	"github.com/go-generalize/api_gen/server_generator/sample/props"
	"github.com/go-generalize/api_gen/server_generator/sample/wrapper"
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
	router.GET("article", r.GetArticle(p))
	return r
}

// GetArticle ...
func (r *Routes) GetArticle(p *props.ControllerProps) echo.HandlerFunc {
	i := NewGetArticleController(p)

	b, ok := (interface{})(i).(interface{ AutoBind() bool })
	bindable := !ok || b.AutoBind()

	return func(c echo.Context) error {
		var (
			req  *GetArticleRequest
			werr *wrapper.APIError
		)

		if bindable {
			req = new(GetArticleRequest)
			if err := c.Bind(req); err != nil {
				log.Printf("failed to JSON binding(/service/article): %+v", err)
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
		res, err := i.GetArticle(c, req)
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

// IGetArticleController ...
type IGetArticleController interface {
	GetArticle(c echo.Context, req *GetArticleRequest) (res *GetArticleResponse, err error)
}
