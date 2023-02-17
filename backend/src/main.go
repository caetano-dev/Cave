package main

import (
	"net/http"

	"github.com/drull1000/notetaking-app/src/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/files", server.FilesIndex)
	http.HandleFunc("/files/show", server.FilesShow)
	http.HandleFunc("/files/upload", server.FilesUpload)

	http.ListenAndServe(":3000", nil)
}
