package api

import (
	"github.com/JamesTiberiusKirk/go_web_template/api/route"
	"github.com/JamesTiberiusKirk/go_web_template/server"
	"github.com/JamesTiberiusKirk/go_web_template/session"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

// Api api struct
type Api struct {
	rootApiPath    string
	publicRoutes   []*route.Route
	authedRoutes   []*route.Route
	echoGroup      *echo.Group
	db             *gorm.DB
	sessionManager *session.Manager
	routes         server.RoutesMap
}

// NewApi new api instance
func NewApi(group *echo.Group, rootApiPath string, db *gorm.DB,
	sesessionManager *session.Manager) *Api {
	return &Api{
		rootApiPath: rootApiPath,
		publicRoutes: []*route.Route{
			route.NewHelloWorld(),
		},
		authedRoutes: []*route.Route{
			route.NewUsersRoute(db, route.NewUserRoute(db)),
		},
		echoGroup:      group,
		db:             db,
		sessionManager: sesessionManager,
		routes:         server.RoutesMap{},
	}
}

// Serve api
func (a *Api) Serve() {
	a.echoGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	a.mapRoutes(&a.publicRoutes)
	a.mapRoutes(&a.authedRoutes, session.SessionAuthMiddleware(a.sessionManager))
}

// GetRoutes returns available routes from this server
func (a *Api) GetRoutes() server.RoutesMap {
	return a.routes
}

func (a *Api) SetRoutes(t string, r server.RoutesMap) {
	// NOOP
}

func (a *Api) mapRoutes(routes *[]*route.Route, middlewares ...echo.MiddlewareFunc) {
	for _, r := range *routes {
		routes := r.Init("", a.echoGroup, middlewares...)

		for k, v := range routes {
			a.routes[k] = a.rootApiPath + v
		}
	}
}
