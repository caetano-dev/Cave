package database

import (
	"database/sql"
	"log"
	"strings"

	s "github.com/drull1000/cave/pkgs/structs"
	_ "github.com/mattn/go-sqlite3"
)

type ClientInterface interface {
	Ping() error
	Connect(name string) error
	Insert(file s.File) (int64, error)
	UpdateFilename(id int64, filename string) error
	DeleteByID(id int64) error
	GetByID(uid int64) (s.FileDatabase, error)
	GetAll() ([]*s.FileDatabase, error)
}

type Client struct {
	DB *sql.DB
}

func (c *Client) Ping() error {
	return c.DB.Ping()
}

// InitDB is the function that initiates the database.
func (c *Client) Connect(name string) error {
	var err error

	c.DB, err = sql.Open("sqlite3", name)
	if err != nil {
		return err
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

	_, err = c.DB.Exec(createTableStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, createTableStmt)
		return err
	}

	return nil
}

// Insert filename and tags. This does not upload the file, it just inserts it in the DB.
func (c *Client) Insert(file s.File) (int64, error) {
	stmt, err := c.DB.Prepare("INSERT INTO files(hash, type, filename, filepath, tags) values(?, ?, ?, ?, ?)")
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
func (c *Client) UpdateFilename(id int64, filename string) error {
	stmt, err := c.DB.Prepare("UPDATE files SET filename=? WHERE uid=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(filename, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID is the function that deletes the file in the database.
func (c *Client) DeleteByID(id int64) error {
	stmt, err := c.DB.Prepare("DELETE FROM FILES WHERE uid=?")
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
func (c *Client) GetByID(uid int64) (s.FileDatabase, error) {
	row := c.DB.QueryRow("SELECT * FROM files WHERE uid = $1", uid)

	file := new(s.FileDatabase)
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
func (c *Client) GetAll() ([]*s.FileDatabase, error) {
	rows, err := c.DB.Query("SELECT * FROM files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	files := make([]*s.FileDatabase, 0)
	for rows.Next() {
		file := new(s.FileDatabase)
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
