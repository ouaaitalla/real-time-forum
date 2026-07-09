package authHandlers

import (
	"net/http"

	"real-time-forum/backend/helpers"
)

func CheckSessionHandler(w http.ResponseWriter, r *http.Request) {
	if helpers.LoggedIn(r) {
		helpers.Res(w, 200, "you're logedin", "success", "", 0)
		return
	}
	helpers.Res(w, 401, "Session not found.", "error", "", 0)
	return
}
