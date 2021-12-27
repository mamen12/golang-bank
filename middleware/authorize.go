package middleware

import (
	"bank/dto"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
)

func NewAuthorize() Authorize {
	return Authorize{}
}

type Authorize struct{}

func (auth Authorize) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			http.Error(w, "Dibutuhkan autentikasi. Silahkan login.", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		credential := &dto.UserCredential{}

		token, err := jwt.ParseWithClaims(tokenString, credential, func(token *jwt.Token) (interface{}, error) {
			return []byte("supersecret"), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token invalid. Dibutuhkan autentikasi. Silahkan login.", http.StatusUnauthorized) // Token expired/key tidak cocok(invalid)
			return
		}

		context.Set(r, "user", credential)
		// fmt.Printf("%+v", claims)
		next.ServeHTTP(w, r)
	})
}
