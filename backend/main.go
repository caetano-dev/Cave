package main

import (
    "database/sql"
    "fmt"
    "strings"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

//File is the struct for the file the user is going to upload.
type File struct{
  filename string
  path string
  tags []string
}

func main() {
    db, err := sql.Open("sqlite3", "./database.db")
    checkErr(err)

    defer db.Close()

  sqlStmt := `
  CREATE TABLE IF NOT EXISTS files (
    uid INTEGER PRIMARY KEY AUTOINCREMENT,
    filename VARCHAR(30) NOT NULL,
    tags VARCHAR(64) NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);
	`
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

  id := insert(db, f)
  update(db, id, g)
  delete(db, id)
}
    
func update(db *sql.DB, id int64, file File) {
    stmt, err := db.Prepare("UPDATE files SET filename=? WHERE uid=?")
    checkErr(err)

    res, err := stmt.Exec("astaxieupdate", id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    // query
    rows, err := db.Query("SELECT * FROM files")
    checkErr(err)
    var uid int
    var filename string
    var path string
    var tags string
    var created string

    for rows.Next() {
        err = rows.Scan(&uid, &filename, &tags, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(filename)
        fmt.Println(path)
        fmt.Println(tags)
        fmt.Println(created)
    }

    rows.Close() //good habit to close

}

func delete(db *sql.DB, id int64){
    stmt, err := db.Prepare("DELETE from FILES WHERE uid=?")
    checkErr(err)

    res, err := stmt.Exec(id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()
}

//insert filename and tags. This is not upload.
func insert(db *sql.DB, file File) int64 {
  stmt, err := db.Prepare("INSERT INTO files(filename, tags) values(?, ?)")
  checkErr(err)

  tags := file.tags

  res, err := stmt.Exec(file.filename, strings.Join(tags, ";"))
  checkErr(err)

  id, err := res.LastInsertId()
  checkErr(err)

  return id
}

func get_all(db *sql.DB, file File){
  stmt, err := db.Prepare("SELECT * FROM files")
  checkErr(err)

  res, err := stmt.Exec()
  checkErr(err)

  fmt.Println(res)
}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}