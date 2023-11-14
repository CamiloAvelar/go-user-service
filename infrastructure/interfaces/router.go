package infrainterfaces

import (
	"github.com/gorilla/mux"
)

type Router struct {
	Injections *ServerInjections
	Router     *mux.Router
}
