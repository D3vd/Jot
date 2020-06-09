package models

type SaveNote struct {
	Content string `json:"content"`
}

type DeleteNote struct {
	ID string `json:"id"`
}

type UpdateNote struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
