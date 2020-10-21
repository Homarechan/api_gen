// Code generated by server_generator. DO NOT EDIT.
// generated version: unknown

package _JobID

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
	router.PUT("job", r.PutJob(p))
	return r
}

// PutJob ...
func (r *Routes) PutJob(p *props.ControllerProps) echo.HandlerFunc {
	i := NewPutJobController(p)
	return func(c echo.Context) error {
		req := new(PutJobRequest)
		if err := c.Bind(req); err != nil {
			log.Printf("failed to JSON binding(/service/user2/{userID}/{JobID}/job): %+v", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PutJob(c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// IPutJobController ...
type IPutJobController interface {
	PutJob(c echo.Context, req *PutJobRequest) (res *PutJobResponse, err error)
}
