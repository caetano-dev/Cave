package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/drull1000/notetaking-app/src/database"
	"github.com/drull1000/notetaking-app/src/server"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Database string
}

func main() {
	var c Config
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
	http.HandleFunc("/healthcheck", env.HealthCheck)

	fmt.Println("Hello")
	http.ListenAndServe("127.0.0.1:3000", nil)
	fmt.Println("World")
}
