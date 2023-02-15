package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: Implement md5 hash to compare files
// File is the struct for the file the user is going to upload.
type File struct {
	// hash string
	filename string
	tags     []string
}

func main() {
	if err := os.Remove("database.db"); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "./database.db")
	checkErr(err)
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

	f := File{
		filename: "example.txt",
		tags:     []string{"example", "text", "file"},
	}

	g := File{
		filename: "example.txt",
		tags:     []string{"text", "file"},
	}

	h := File{
		filename: "another_file.txt",
		tags:     []string{"text", "file"},
	}

	i := File{
		filename: "example.txt",
		tags:     []string{"text"},
	}

	f_id := insert(db, f)
	g_id := insert(db, g)
	insert(db, i)

	update(db, g_id, h)
	delete(db, f_id)

	getAll(db)
}

// insert filename and tags. This is not upload.
func insert(db *sql.DB, file File) int64 {
	stmt, err := db.Prepare("INSERT INTO files(filename, tags) values(?, ?)")
	checkErr(err)

	tags := file.tags
	tagsString := strings.Join(tags, ";")

	res, err := stmt.Exec(file.filename, tagsString)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	return id
}

func update(db *sql.DB, id int64, file File) {
	stmt, err := db.Prepare("UPDATE files SET tags=? WHERE uid=?")
	checkErr(err)

	tags := file.tags
	tagsString := strings.Join(tags, ";")

	_, err = stmt.Exec(tagsString, id)
	checkErr(err)
}

func delete(db *sql.DB, id int64) {
	stmt, err := db.Prepare("DELETE from FILES WHERE uid=?")
	checkErr(err)

	_, err = stmt.Exec(id)
	checkErr(err)
}

func getAll(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM files")
	checkErr(err)

	var uid int
	var filename string
	var tags string
	var created string

	for rows.Next() {
		err = rows.Scan(&uid, &filename, &tags, &created)
		checkErr(err)
		fmt.Println(uid, filename, tags, created)
	}

	rows.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
