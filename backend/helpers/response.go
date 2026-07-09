package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real-time-forum/backend/models"
)

func Res(w http.ResponseWriter, statcode int, text string, tp string, user string, id int64) {
	var res models.Res
	res.Type = tp
	res.Username = user
	res.Text = text
	res.ID = id
	w.WriteHeader(statcode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}

func ResData(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		fmt.Println(err)
		Res(w, 500, "internal server error", "error", "", 0)
		return
	}
}
