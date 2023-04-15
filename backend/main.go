package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/drull1000/cave/pkgs/api"
	"github.com/drull1000/cave/pkgs/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbClient := &database.Client{}
	if err := dbClient.Connect("database.db"); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Mount("/files", api.FilesResource{}.Routes(dbClient))

	fmt.Println("http://127.0.0.1:3000/")

	http.ListenAndServe("127.0.0.1:3000", r)
}
