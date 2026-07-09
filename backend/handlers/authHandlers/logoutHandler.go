package authHandlers

import (
	"fmt"
	"net/http"

	"real-time-forum/backend/helpers"
	"real-time-forum/database"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	user_id, err := helpers.GetUserFromRequest(r)
	if err != nil {
		helpers.Res(w, 404, "not found user", "error", "", 0)
		return
	}
	_, err = database.DB.Exec("DELETE FROM sessions WHERE user_id = ?", user_id)
	if err != nil {
		fmt.Println(err, "hello")
		helpers.Res(w, 404, "Failed Delet Session", "error", "", 0)
		return
	}
	helpers.Res(w, 200, "Session deleted successfully", "success", "", 0)
	return
}
