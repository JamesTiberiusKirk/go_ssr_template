package server

type RoutesMap map[string]string

type Server interface {
	Serve()
}
