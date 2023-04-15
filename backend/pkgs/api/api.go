package api

import (
	"github.com/drull1000/cave/pkgs/database"

	"github.com/go-chi/chi/v5"
)

type FilesResource struct{}

var DBClient database.ClientInterface

func SetDBClient(c database.ClientInterface) {
	DBClient = c
}

func (rs FilesResource) Routes(dbClient database.ClientInterface) chi.Router {
	r := chi.NewRouter()

	SetDBClient(dbClient)

	r.Get("/", rs.List)
	r.Post("/", rs.Create)
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)
		r.Put("/", rs.Update)
		r.Delete("/", rs.Delete)
	})

	return r
}

// http.HandleFunc("/files", env.FilesShowAll)
// http.HandleFunc("/files/show", env.FilesShow)
// http.HandleFunc("/files/upload", env.FilesUpload)
// http.HandleFunc("/files/delete", env.FilesDelete)
// http.HandleFunc("/files/create", env.FilesCreate)
// http.HandleFunc("/fileEditContent", env.FileEditContent)
// http.HandleFunc("/fileEditName", env.FileEditName)
