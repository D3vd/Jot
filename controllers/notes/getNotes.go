package notes

import "net/http"

func (n Controller) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("notes"))
}
