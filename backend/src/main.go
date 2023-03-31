package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/drull1000/notetaking-app/src/database"
	"github.com/drull1000/notetaking-app/src/server"
	s "github.com/drull1000/notetaking-app/src/structs"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	var c s.Config
	flag.StringVar(&c.Database, "database", os.Getenv("DATABASE_NAME"), "database name")
	flag.Parse()

	db, err := database.InitDB(c.Database)
	if err != nil {
		log.Fatal(err)
	}

	env := &server.Env{DB: db}

	http.HandleFunc("/files", env.FilesShowAll)
	http.HandleFunc("/files/show", env.FilesShow)
	http.HandleFunc("/files/upload", env.FilesUpload)
	http.HandleFunc("/files/delete", env.FilesDelete)
	http.HandleFunc("/fileEditContent", env.FileEditContent)
	http.HandleFunc("/fileEditName", env.FileEditName)

	fmt.Println("https://127.0.0.1:3000/")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
