package routers

import (
	"github.com/gorilla/mux"
	"github.com/joshbrusa/cad-http/src/middlewares"
)

type UserRouter struct {
	Router *mux.Router
}

func NewUserRouter(router *mux.Router) *UserRouter {
	return &UserRouter{
		Router: router,
	}
}

func (userRouter *UserRouter) UseMiddlewares(
	loggerMiddleware *middlewares.LoggerMiddleware,
	panicMiddleware *middlewares.PanicMiddleware,
) {
	userRouter.Router.Use(loggerMiddleware.Handle)
	userRouter.Router.Use(panicMiddleware.Handle)
}
