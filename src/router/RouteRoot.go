package router

func (router *Router) RouteRoute() {
	router.Mux.HandleFunc("/", router.Handler.HandleRoot())
}
