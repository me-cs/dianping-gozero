package middlewares

import (
	"net/http"

	"github.com/me-cs/dianping-gozero/library/metadatax"
	"github.com/thinkeridea/go-extend/exnet"
	"google.golang.org/grpc/metadata"
)

type NetworkMiddleware struct {
}

func NewNetworkMiddleware() *NetworkMiddleware {
	return &NetworkMiddleware{}
}

func (m *NetworkMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteIp := exnet.ClientPublicIP(r)
		md, ok := metadata.FromIncomingContext(r.Context())
		if !ok {
			md = metadata.MD{}
		}
		md.Set(metadatax.RemoteIP, remoteIp)
		newCtx := metadata.NewIncomingContext(r.Context(), md)
		r = r.WithContext(newCtx)
		next(w, r)
	}
}
