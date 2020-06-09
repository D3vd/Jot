package notes

import (
	"encoding/json"
	"github.com/R3l3ntl3ss/Jot/models"
	"net/http"
)

func (n Controller) UpdateNote(w http.ResponseWriter, r *http.Request) {

	var req models.UpdateNote

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.ID == "" || req.Content == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		res := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Ensure that the request body is formatted right",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err = n.M.UpdateNote(req.ID, req.Content)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		res := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Note with Id does not exist / there is no change in content",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := models.SimpleResponse{
		Code:    http.StatusOK,
		Message: "Successfully updated Note",
	}

	json.NewEncoder(w).Encode(res)
}
