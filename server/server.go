package server

type RoutesMap map[string]string

type Server interface {
	Serve(existingRoutes RoutesMap)
	GetRoutes() RoutesMap
	GetRoutesType() string
}
