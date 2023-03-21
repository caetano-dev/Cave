package database

import (
	"database/sql"
	"fmt"
	"reflect"

	"testing"

	s "github.com/drull1000/notetaking-app/src/structs"
	_ "github.com/mattn/go-sqlite3"
)

func TestInitDB(t *testing.T) {
	// initialize an in-memory database for testing
	db, err := InitDB(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// check if the "files" table was created
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='files'")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		t.Error("InitDB() did not create the \"files\" table")
	}
}

func TestInsert(t *testing.T) {
	db, _ := InitDB(":memory:")
	file := s.File{
		Hash:     "hash1",
		Type:     "file",
		Filename: "file1.txt",
		Filepath: "files/file1.txt",
		Tags:     []string{"tag1", "tag2"},
	}

	type args struct {
		db   *sql.DB
		file s.File
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Inserts a file successfully",
			args: args{
				db:   db,
				file: file,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				stmt, err := tt.args.db.Prepare("INSERT INTO files(hash, type, filename, filepath, tags) values(?, ?, ?, ?, ?)")
				if err != nil {
					t.Fatal(err)
				}
				defer stmt.Close()
			}

			got, err := Insert(tt.args.db, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteByID(t *testing.T) {
	// initialize an in-memory database for testing
	db, err := InitDB(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// insert a test file into the database
	file := s.File{
		Hash:     "hash1",
		Type:     "file",
		Filename: "file1.txt",
		Filepath: "files/file1.txt",
		Tags:     []string{"tag1", "tag2"},
	}
	id, err := Insert(db, file)
	if err != nil {
		t.Fatal(err)
	}

	// test deleting the file by its ID
	err = DeleteByID(db, id)
	if err != nil {
		t.Errorf("DeleteByID() error = %v, wantErr nil", err)
	}

	// verify that the file is no longer in the database
	rows, err := db.Query("SELECT * FROM files WHERE uid=?", id)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() {
		t.Errorf("DeleteByID() did not delete file with ID %d", id)
	}
}

func TestGetByID_ValidID(t *testing.T) {
	// initialize an in-memory database for testing
	db, err := InitDB(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// insert a file into the database
	file := s.File{
		Hash:     "abc123",
		Type:     "file",
		Filename: "file.txt",
		Filepath: "files/file.txt",
		Tags:     []string{"tag1", "tag2"},
	}
	_, err = Insert(db, file)
	if err != nil {
		t.Fatal(err)
	}

	// get the file by ID
	dbFile, err := GetByID(db, 1)
	fmt.Println(dbFile)
	if err != nil {
		t.Fatal(err)
	}

	// verify that the file was retrieved correctly
	expectedFile := s.FileDatabase{
		ID:        1,
		Hash:      "abc123",
		Type:      "file",
		Filename:  "file.txt",
		Filepath:  "files/file.txt",
		Tags:      []string{"tag1", "tag2"},
		CreatedAt: dbFile.CreatedAt,
	}
	if !reflect.DeepEqual(dbFile, expectedFile) {
		t.Errorf("GetByID() returned %+v, expected %+v", dbFile, expectedFile)
	}
}

func TestGetAll(t *testing.T) {
	// initialize an in-memory database for testing
	db, err := InitDB(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// insert test data into the database
	file1 := &s.File{Hash: "hash1", Type: "file", Filename: "file1.txt", Filepath: "files/file1.txt", Tags: []string{"tag1", "tag2"}}
	file2 := &s.File{Hash: "hash2", Type: "file", Filename: "file2.txt", Filepath: "files/file2.txt", Tags: []string{"tag3", "tag4"}}
	_, err = Insert(db, *file1)

	if err != nil {
		t.Fatal(err)
	}
	_, err = Insert(db, *file2)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve all files from the database
	files, err := GetAll(db)
	if err != nil {
		t.Fatal(err)
	}

	// check that the correct number of files were retrieved
	if len(files) != 2 {
		t.Errorf("expected %d files, got %d", 2, len(files))
	}

	// check that the retrieved files have the correct data
	for i, file := range files {
		if file.ID != int64(i+1) {
			t.Errorf("expected file %d to have ID %d, got %d", i+1, i+1, file.ID)
		}
		if file.Hash != fmt.Sprintf("hash%d", i+1) {
			t.Errorf("expected file %d to have hash %s, got %s", i+1, fmt.Sprintf("hash%d", i+1), file.Hash)
		}
		if file.Type != fmt.Sprintf("%s", file.Type) {
			t.Errorf("expected file %d to have Type %s, got %s", i+1, fmt.Sprintf("Type%d", i+1), file.Type)
		}
		if file.Filename != fmt.Sprintf("file%d.txt", i+1) {
			t.Errorf("expected file %d to have filename %s, got %s", i+1, fmt.Sprintf("file%d.txt", i+1), file.Filename)
		}
		if file.Filepath != fmt.Sprintf("files/file%d.txt", i+1) {
			t.Errorf("expected file %d to have filepath %s, got %s", i+1, fmt.Sprintf("files/file%d.txt", i+1), file.Filepath)
		}
		if !reflect.DeepEqual(file.Tags, []string{fmt.Sprintf("tag%d", i*2+1), fmt.Sprintf("tag%d", i*2+2)}) {
			t.Errorf("expected file %d to have tags %v, got %v", i+1, []string{fmt.Sprintf("tag%d", i*2+1), fmt.Sprintf("tag%d", i*2+2)}, file.Tags)
		}
	}
}
