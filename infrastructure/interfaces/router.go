package infrainterfaces

import (
	"github.com/gorilla/mux"
)

type Router struct {
	Injections HttpServerInjections
	Router     *mux.Router
}
