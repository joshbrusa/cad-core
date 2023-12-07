package router

import (
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/handler"
)

type RootRouter struct {
	Mux         *http.ServeMux
	RootHandler *handler.RootHandler
	UserRouter  *UserRouter
}

func NewRootRouter(
	env *core.Env,
	logger *core.Logger,
	mux *http.ServeMux,
) *RootRouter {
	rootHandler := handler.NewRootHandler(env, logger)
	userRouter := NewUserRouter(env, logger, mux)

	return &RootRouter{
		Mux:         mux,
		RootHandler: rootHandler,
		UserRouter:  userRouter,
	}
}

func (rootRouter *RootRouter) Route() {
	rootRouter.Mux.HandleFunc("/", rootRouter.RootHandler.HandleRoot())
}
