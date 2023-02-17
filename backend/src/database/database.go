package database

import (
	"database/sql"
	"hash"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: Implement sha256 hash to compare files
// File is the struct for the file the user is going to upload.
type File struct {
	Hash     hash.Hash
	Filename string
	Tags     []string
}

type FileDatabase struct {
	ID        int64
	Hash      string
	Filename  string
	Tags      string
	CreatedAt string
}

var DB *sql.DB

func init() {
	var err error

	if err = os.Remove("database.db"); err != nil {
		log.Fatal(err)
	}

	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS files (
    	uid INTEGER PRIMARY KEY AUTOINCREMENT,
    	hash BLOB NOT NULL UNIQUE,
    	filename VARCHAR(30) NOT NULL,
    	tags VARCHAR(64) NULL,
    	created_at TEXT DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

// insert filename and tags. This is not upload.
func Insert(db *sql.DB, file File) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO files(hash, filename, tags) values(?, ?, ?)")
	if err != nil {
		return 0, err
	}

	tags := file.Tags
	tagsString := strings.Join(tags, ";")

	res, err := stmt.Exec(file.Hash, file.Filename, tagsString)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// func update(db *sql.DB, id int64, file File) error {
// 	stmt, err := db.Prepare("UPDATE files SET tags=? WHERE uid=?")
// 	if err != nil {
// 		return err
// 	}

// 	tags := file.tags
// 	tagsString := strings.Join(tags, ";")

// 	_, err = stmt.Exec(tagsString, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func delete(db *sql.DB, id int64) error {
// 	stmt, err := db.Prepare("DELETE from FILES WHERE uid=?")
// 	if err != nil {
// 		return err
// 	}

// 	_, err = stmt.Exec(id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func GetByID(db *sql.DB, uid int64) (FileDatabase, error) {
	row := db.QueryRow("SELECT * FROM files WHERE uid = $1", uid)

	file := new(FileDatabase)
	err := row.Scan(&file.ID, &file.Hash, &file.Filename, &file.Tags, &file.CreatedAt)
	if err == sql.ErrNoRows {
		return *file, nil
	} else if err != nil {
		return *file, err
	}

	return *file, nil
}

func GetAll(db *sql.DB) ([]*FileDatabase, error) {
	rows, err := db.Query("SELECT * FROM files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	files := make([]*FileDatabase, 0)
	for rows.Next() {
		file := new(FileDatabase)
		err := rows.Scan(&file.ID, &file.Hash, &file.Filename, &file.Tags, &file.CreatedAt)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}
