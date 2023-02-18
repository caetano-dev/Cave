package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: Implement sha256 hash to compare files
// File is the struct for the file the user is going to upload.
type File struct {
	Hash     string
	Filename string
	Tags     []string
}

// File database is the struct with the database colunms
//type FileDatabase struct {
//	ID        int64
//	Hash      string
//	Filename  string
//	Tags      string
//	CreatedAt string
//}

type FileDatabase struct {
	ID        int64    `json:"id"`
	Hash      string   `json:"hash"`
	Filename  string   `json:"filename"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
}

// InitDB is the function that initiates the database.
func InitDB(name string) (*sql.DB, error) {
	var err error

	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS files (
    	uid INTEGER PRIMARY KEY AUTOINCREMENT,
    	hash BLOB NOT NULL,
    	filename VARCHAR(30) NOT NULL,
    	tags VARCHAR(64) NULL,
    	created_at TEXT DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}

	return db, nil
}

// Insert filename and tags. This does not upload the file, it just inserts it in the DB.
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

// UpdateFile is the function that updated the file in the database
// func UpdateFile(db *sql.DB, id int64, file File) error {
// 	stmt, err := db.Prepare("UPDATE files SET hash=?, filename=?, tags=? WHERE uid=?")
// 	if err != nil {
// 		return err
// 	}

// 	tags := file.tags
// 	tagsString := strings.Join(tags, ";")

// 	_, err = stmt.Exec(file.Hash, file.Filename, tagsString, id)
// 	if err != nil {
// 		return err
// 	}

//	return nil
// }

// DeleteByID is the function that deletes the file in the database.
func DeleteByID(db *sql.DB, id int64) error {
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

// GetByID is the function that gets a file by ID from the database.
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

// GetAll is the function that gets all the files from the database.
func GetAll(db *sql.DB) ([]*FileDatabase, error) {
	rows, err := db.Query("SELECT * FROM files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	files := make([]*FileDatabase, 0)
	for rows.Next() {
		file := new(FileDatabase)
		var tags string // temporary variable to hold string value of "tags" column
		err := rows.Scan(&file.ID, &file.Hash, &file.Filename, &tags, &file.CreatedAt)
		if err != nil {
			return nil, err
		}
		fmt.Println(file.Tags)

		file.Tags = strings.Split(tags, ";")
		files = append(files, file)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}
