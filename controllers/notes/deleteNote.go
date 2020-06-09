package notes

import (
	"encoding/json"
	"github.com/R3l3ntl3ss/Jot/models"
	"net/http"
)

func (n Controller) DeleteNote(w http.ResponseWriter, r *http.Request) {

	var req models.DeleteNote

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.ID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		res := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Ensure that the request body is formatted right",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if err = n.M.DeleteNote(req.ID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		res := models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while deleting note from DB",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := models.SimpleResponse{
		Code:    http.StatusOK,
		Message: "Successfully Deleted Note from DB",
	}

	json.NewEncoder(w).Encode(res)

}
