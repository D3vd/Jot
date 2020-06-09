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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		res := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Ensure that the request body is formatted right",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err = n.M.SaveNote(req.Content)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		res := models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Error while inserting notes into DB",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := models.SimpleResponse{
		Code:    http.StatusOK,
		Message: "Successfully Saved Note",
	}

	json.NewEncoder(w).Encode(res)
}
