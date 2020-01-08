package middleware

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"../model"
	"fmt"
	"golang.org/x/net/context"
)

type Key int
const MyKey Key = 0

func Validate(page http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		// if no Auth cookie is set then return a 404 not found
		cookie, err := req.Cookie("Auth")
		if err != nil {
			http.NotFound(res,req)
			return
		}

		// return a token using the cookie
		token, err := jwt.ParseWithClaims(cookie.Value, &model.Claims{}, func(token *jwt.Token)(interface{}, error) {
			// Make sure token signature wasn't change
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte("secret"), nil
		})

		if err != nil {
			http.NotFound(res, req)
			return
		}

		// Grab the token claims and pass them to the original request
		 if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			 ctx := context.WithValue(req.Context(), MyKey, *claims)
			 page(res, req.WithContext(ctx))
		 } else {
			 http.NotFound(res, req)
			 return
		 }
	})
}
