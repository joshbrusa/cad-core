package router

import (
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
)

type UserRouter struct {
	Env    *core.Env
	Logger *core.Logger
	Mux    *http.ServeMux
}

func NewUserRouter(
	env *core.Env,
	logger *core.Logger,
	mux *http.ServeMux,
) *UserRouter {
	return &UserRouter{
		Env:    env,
		Logger: logger,
		Mux:    mux,
	}
}

func (userRouter *UserRouter) Route() {
	userRouter.Mux.HandleFunc("/", userRouter)
}
