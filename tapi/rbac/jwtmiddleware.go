package rbac

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type JWTMiddleware struct {
	jwtMW *jwtmiddleware.JWTMiddleware
}

func NewJWTMiddlware(authnValidationKey string) *JWTMiddleware {
	return &JWTMiddleware{
		jwtMW: jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(authnValidationKey), nil
			},
			SigningMethod:       jwt.SigningMethodHS256,
			CredentialsOptional: false,
		}),
	}
}

func (j JWTMiddleware) AuthenticateRouter(router *mux.Router) *negroni.Negroni {
	return negroni.New(negroni.HandlerFunc(j.jwtMW.HandlerWithNext), negroni.Wrap(router))
}
