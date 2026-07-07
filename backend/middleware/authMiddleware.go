package middleware

import (
	"net/http"
	"real-time-forum/backend/helpers"
	"real-time-forum/helpers"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		if (!helpers.LoggedIn(r)){
			helpers.Res(w, 404, "user not ")
		}
	}
}