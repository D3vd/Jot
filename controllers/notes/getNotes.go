package notes

import (
	"encoding/json"
	"github.com/R3l3ntl3ss/Jot/models"
	"net/http"
)

func (n Controller) GetAllNotes(w http.ResponseWriter, r *http.Request) {

	notes, err := n.M.GetAllNotes()

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

	res := models.AllNotesResponse{
		Count: len(notes),
		Notes: notes,
	}

	json.NewEncoder(w).Encode(res)
}
