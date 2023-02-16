package main

import (
	"database/sql"
	"fmt"
	"log"
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
