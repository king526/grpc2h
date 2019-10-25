package pb

import (
	"net/http"

	"github.com/king526/grpc2h"
)

func GetHTTPHandler(srv interface{}) http.Handler {
	var b grpc2h.HandlerBuilder
	b.SetGRPCServer(srv)
	// that is why put builder method here: so we use unexported server stream implement.
	b.RegisterSeverStream(itemServiceListServer{})
	h, _ := b.Build()
	return h
}
