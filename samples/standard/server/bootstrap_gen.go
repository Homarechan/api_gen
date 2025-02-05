// Code generated by server_generator. DO NOT EDIT.
// generated version: devel

package sample

import (
	"io"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/go-generalize/api_gen/samples/standard/server/props"
	serverService "github.com/go-generalize/api_gen/samples/standard/server/service"
	serverServiceStaticPage "github.com/go-generalize/api_gen/samples/standard/server/service/static_page"
	serverServiceUser "github.com/go-generalize/api_gen/samples/standard/server/service/user"
	serverServiceUser2 "github.com/go-generalize/api_gen/samples/standard/server/service/user2"
	serverServiceUser2UserID "github.com/go-generalize/api_gen/samples/standard/server/service/user2/_userID"
	serverServiceUser2UserIDJobID "github.com/go-generalize/api_gen/samples/standard/server/service/user2/_userID/_JobID"
	"github.com/labstack/echo/v4"
)

// MiddlewareList ...
type MiddlewareList []*MiddlewareSet

// MiddlewareMap ...
type MiddlewareMap map[string][]echo.MiddlewareFunc

// MiddlewareSet ...
type MiddlewareSet struct {
	Path           string
	MiddlewareFunc []echo.MiddlewareFunc
}

// ToMap ...
func (m MiddlewareList) ToMap() MiddlewareMap {
	mf := make(map[string][]echo.MiddlewareFunc)
	for _, middleware := range m {
		mf[middleware.Path] = middleware.MiddlewareFunc
	}
	return mf
}

// Bootstrap ...
func Bootstrap(p *props.ControllerProps, e *echo.Echo, middlewareList MiddlewareList, opts ...io.Writer) {
	if len(opts) > 0 {
		if w := opts[0]; w != nil {
			log.SetOutput(w)
		}
	}

	middleware := middlewareList.ToMap()

	// error handling
	e.Use(func(before echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				recoverErr := recover()
				if recoverErr == nil {
					return
				}

				debug.PrintStack()

				if httpErr, ok := recoverErr.(*echo.HTTPError); ok {
					err = c.JSON(httpErr.Code, httpErr.Message)
				}

				err = c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "internal server error.",
				})
			}()

			return before(c)
		}
	})

	rootGroup := e.Group("/")
	setMiddleware(rootGroup, "/", middleware)
	NewRoutes(p, rootGroup, opts...)

	serverServiceGroup := rootGroup.Group("service/")
	setMiddleware(serverServiceGroup, "/service/", middleware)
	serverService.NewRoutes(p, serverServiceGroup, opts...)

	serverServiceStaticPageGroup := serverServiceGroup.Group("static_page/")
	setMiddleware(serverServiceStaticPageGroup, "/service/static_page/", middleware)
	serverServiceStaticPage.NewRoutes(p, serverServiceStaticPageGroup, opts...)

	serverServiceUserGroup := serverServiceGroup.Group("user/")
	setMiddleware(serverServiceUserGroup, "/service/user/", middleware)
	serverServiceUser.NewRoutes(p, serverServiceUserGroup, opts...)

	serverServiceUser2Group := serverServiceGroup.Group("user2/")
	setMiddleware(serverServiceUser2Group, "/service/user2/", middleware)
	serverServiceUser2.NewRoutes(p, serverServiceUser2Group, opts...)

	serverServiceUser2UserIDGroup := serverServiceUser2Group.Group(":userID/")
	setMiddleware(serverServiceUser2UserIDGroup, "/service/user2/:userID/", middleware)
	serverServiceUser2UserID.NewRoutes(p, serverServiceUser2UserIDGroup, opts...)

	serverServiceUser2UserIDJobIDGroup := serverServiceUser2UserIDGroup.Group(":JobID/")
	setMiddleware(serverServiceUser2UserIDJobIDGroup, "/service/user2/:userID/:JobID/", middleware)
	serverServiceUser2UserIDJobID.NewRoutes(p, serverServiceUser2UserIDJobIDGroup, opts...)
}

func setMiddleware(group *echo.Group, path string, list MiddlewareMap) {
	if ms, ok := list[path]; ok {
		for i := range ms {
			group.Use(ms[i])
		}
	}
}
