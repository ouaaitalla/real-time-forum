package middleware
import (
	"net/http"
)
func MethodeMiddleware(method string, next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if (r.Method !=  method){
			//err
		}
		next(w, r)
	}
}