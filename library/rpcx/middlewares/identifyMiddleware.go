package middlewares

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/homelight/json"
	"github.com/me-cs/dianping-gozero/library/auth"
	"github.com/me-cs/dianping-gozero/library/metadatax"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/token"
	"google.golang.org/grpc/metadata"
)

var (
	Unauthorized = errors.New("user unauthorized")
)

type IdentifyMiddleware struct {
	jwtSecret string
}

func NewIdentifyMiddleware(jwtSecret string) *IdentifyMiddleware {
	return &IdentifyMiddleware{jwtSecret: jwtSecret}
}

func (m *IdentifyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	parser := token.NewTokenParser()
	return func(w http.ResponseWriter, r *http.Request) {
		userToken, err := parser.ParseToken(r, m.jwtSecret, "")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if !userToken.Valid {
			httpx.Error(w, Unauthorized)
			return
		}
		user, err := userFromRequest(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		d, err := json.Marshal(user)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		md := metadata.Pairs(metadatax.UserInner, base64.StdEncoding.EncodeToString(d))
		newCtx := metadata.NewIncomingContext(r.Context(), md)
		r.WithContext(newCtx)
		next(w, r)
	}
}

func userFromRequest(r *http.Request) (auth.User, error) {
	userInfo, ok := r.Context().Value(metadatax.User).(string)
	if !ok {
		return nil, Unauthorized
	}
	user, err := auth.DecodeUser(userInfo)
	if err != nil {
		return nil, err
	}
	return user, nil
}
