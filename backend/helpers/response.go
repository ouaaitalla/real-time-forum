package helpers

import (
	"net/http"
	"encoding/json"
	"real-time-forum/backend/models"
)
func Res(w http.ResponseWriter, statcode int, text string, tp string, user string){
	var res models.Res
	res.Type = tp;
	res.Username = user
	res.Text = text
	w.WriteHeader(statcode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}