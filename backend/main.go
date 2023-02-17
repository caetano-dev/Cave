package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: Implement sha256 hash to compare files
// File is the struct for the file the user is going to upload.
type File struct {
	hash     string
	filename string
	tags     []string
}

type FileDatabase struct {
	uid       int64
	hash      string
	filename  string
	tags      []string
	createdAt string
}

func main() {
	r := mux.NewRouter()

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		getAll(db)
		fmt.Println("Database opened and all items fetched")
	})

  //TODO: download the file in the user's device. For now, this only fetches the file.
	r.HandleFunc("/download/{filename}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		filename, ok := vars["filename"]
		if !ok {
			fmt.Println("filename is missing in paramenters")
		}
		file, err := getFile(db, filename)
		if err != nil {
			fmt.Println(file)
		}
	})

	http.ListenAndServe(":3000", r)
}

// code for testing the database and performing all operations.
func createDB() {
	if err := os.Remove("database.db"); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS files (
    	uid INTEGER PRIMARY KEY AUTOINCREMENT,
    	filename VARCHAR(30) NOT NULL,
    	tags VARCHAR(64) NULL,
    	created_at TEXT DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

// insert filename and tags. This is not upload.
func insert(db *sql.DB, file File) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO files(filename, tags) values(?, ?)")
	if err != nil {
		return 0, err
	}

	tags := file.tags
	tagsString := strings.Join(tags, ";")

	res, err := stmt.Exec(file.filename, tagsString)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func update(db *sql.DB, id int64, file File) error {
	stmt, err := db.Prepare("UPDATE files SET tags=? WHERE uid=?")
	if err != nil {
		return err
	}

	tags := file.tags
	tagsString := strings.Join(tags, ";")

	_, err = stmt.Exec(tagsString, id)
	if err != nil {
		return err
	}

	return nil
}

func delete(db *sql.DB, id int64) error {
	stmt, err := db.Prepare("DELETE from FILES WHERE uid=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func getFile(db *sql.DB, filename string) (FileDatabase, error) {
	var file FileDatabase

	err := db.QueryRow("SELECT * FROM files WHERE filename=?", filename).Scan(&file.uid, &file.filename, &file.tags, &file.createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return file, fmt.Errorf("file %d not found", filename)
		}
		return file, err
	}
	return file, nil
}

func getAll(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM files")
	if err != nil {
		return err
	}

	var tempFile FileDatabase

	for rows.Next() {
		err = rows.Scan(&tempFile.uid, &tempFile.filename, &tempFile.tags, &tempFile.createdAt)
		if err != nil {
			return err
		}

		fmt.Println(tempFile.uid, tempFile.filename, tempFile.tags, tempFile.createdAt)
	}

	rows.Close()

	return nil
}

// add the db parameters
func addFileToDatabase(filename string, tags string, filePath string) error {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}
	defer db.Close()

	sha256, err := getFileHash(filePath)
	if err != nil {
		return err
	}
	createdAt := time.Now().Format("2006-01-02 15:04:05")

	_, err = db.Exec("INSERT INTO files (filename, tags, sha256, createdAt) VALUES (?, ?, ?, ?)", filename, tags, sha256, createdAt)
	if err != nil {
		return err
	}

	return nil
}
