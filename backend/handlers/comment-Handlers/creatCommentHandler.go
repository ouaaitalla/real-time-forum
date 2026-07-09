package commenthandlers

import (
	"encoding/json"
	"net/http"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/models"
)

func CreatCommentHandler(w http.ResponseWriter, r *http.Request) {
	var req models.ReqCreatComment
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.Res(w, 400, "Failed to decode request", "error", "", 0)
		return
	}
	req

}
