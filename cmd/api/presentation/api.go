package api

type API struct {
}

type Handler interface {
	RegisterRoutes()
}
