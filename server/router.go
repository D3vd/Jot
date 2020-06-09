package server

import (
	"github.com/R3l3ntl3ss/Jot/controllers/notes"
	"github.com/R3l3ntl3ss/Jot/mongo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := &mongo.Mongo{}
	db.Init()

	r.Route("/notes", func(r chi.Router) {

		n := notes.Controller{
			M: db,
		}

		r.Get("/", n.GetAllNotes)
		r.Put("/", n.SaveNote)
		r.Post("/", n.UpdateNote)
		r.Delete("/", n.DeleteNote)
	})

	return r
}
