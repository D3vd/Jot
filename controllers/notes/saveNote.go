package notes

import (
	"encoding/json"
	"github.com/R3l3ntl3ss/Jot/models"
	"net/http"
)

func (n Controller) SaveNote(w http.ResponseWriter, r *http.Request) {
	var req models.SaveNote

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}

	w.Write([]byte(req.Content))
}
