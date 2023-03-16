package database

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: Implement sha256 hash to compare files
// File is the struct for the file the user is going to upload.
type File struct {
	Hash     string
	Type     string
	Filename string
	Filepath string
	Tags     []string
}

// FileDatabase is the structure of the file in the database
type FileDatabase struct {
	ID        int64    `json:"id"`
	Hash      string   `json:"hash"`
	Type      string   `json:"type"`
	Filename  string   `json:"filename"`
	Filepath  string   `json:"filepath"`
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

	const createTableStmt = `
	CREATE TABLE IF NOT EXISTS files (
    	uid INTEGER PRIMARY KEY AUTOINCREMENT,
    	hash BLOB NOT NULL,
    	type VARCHAR(5) NULL,
    	filename VARCHAR(30) NOT NULL,
    	filepath VARCHAR(50) NOT NULL,
    	tags VARCHAR(64) NULL,
    	created_at TEXT DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, createTableStmt)
		return nil, err
	}

	return db, nil
}

// Insert filename and tags. This does not upload the file, it just inserts it in the DB.
func Insert(db *sql.DB, file File) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO files(hash, type, filename, filepath, tags) values(?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	tags := file.Tags
	tagsString := strings.Join(tags, ";")

	res, err := stmt.Exec(file.Hash, "file", file.Filename, file.Filepath, tagsString) //todo: remove this hardcode value
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateFile is the function that updated the file in the database. It receives the file ID and the file structure
// func UpdateFileContent(db *sql.DB, id int64, file File) error {
// 	stmt, err := db.Prepare("UPDATE files SET content=? WHERE uid=?")
// 	if err != nil {
// 		return err
// 	}

// 	_, err = stmt.Exec(File.content, id)
// 	if err != nil {
// 		return err
// 	}

//	return nil
//}

// DeleteByID is the function that deletes the file in the database.
func DeleteByID(db *sql.DB, id int64) error {
	stmt, err := db.Prepare("DELETE FROM FILES WHERE uid=?")
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
	var tags string // temporary variable to hold string value of "tags" column
	err := row.Scan(&file.ID, &file.Hash, &file.Type, &file.Filename, &file.Filepath, &tags, &file.CreatedAt)
	file.Tags = strings.Split(tags, ";")
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
		err := rows.Scan(&file.ID, &file.Hash, &file.Type, &file.Filename, &file.Filepath, &tags, &file.CreatedAt)
		if err != nil {
			return nil, err
		}

		file.Tags = strings.Split(tags, ";")
		files = append(files, file)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}
