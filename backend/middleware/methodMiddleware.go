package middleware

import (
	"net/http"
	"real-time-forum/backend/helpers"
)
func MethodeMiddleware(method string, next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if (r.Method !=  method){
			helpers.Res(w, 405, "Method not allowed", "error", "", 0)
			return
		}
		next(w, r)
	}
}