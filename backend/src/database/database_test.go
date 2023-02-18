package database

import (
	"database/sql"

	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInsert(t *testing.T) {
	db, _ := InitDB(":memory:")
	file := File{
		Hash:     "hash1",
		Filename: "file1.txt",
		Tags:     []string{"tag1", "tag2"},
	}

	type args struct {
		db   *sql.DB
		file File
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
				stmt, err := tt.args.db.Prepare("INSERT INTO files(hash, filename, tags) values(?, ?, ?)")
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
