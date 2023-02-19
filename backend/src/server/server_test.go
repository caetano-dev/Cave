package server

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/drull1000/notetaking-app/src/database"
)

func TestEnv_HealthCheck(t *testing.T) {
	db, _ := database.InitDB(":memory:")
	env := &Env{DB: db}

	request := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	response := httptest.NewRecorder()
	env.HealthCheck(response, request)
	if response.Code != 200 {
		t.Errorf("Expected %d, received %d", 200, response.Code)
	}
}

func TestEnv_FilesUpload(t *testing.T) {
	db, _ := database.InitDB(":memory:")
	env := &Env{DB: db}

	r, w := io.Pipe()
	defer r.Close()

	m := multipart.NewWriter(w)
	tmpdir := t.TempDir()

	go func() {
		file, err := os.Create(filepath.Join(tmpdir, "file.txt"))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.WriteString("Hello World!")
		if err != nil {
			log.Fatal(err)
		}

		part, err := m.CreateFormFile("myFile", "file.txt")
		if err != nil {
			// The error is returned from read on the pipe.
			w.CloseWithError(err)
			return
		}

		if _, err := io.Copy(part, file); err != nil {
			// The error is returned from read on the pipe.
			w.CloseWithError(err)
			return
		}

		w.Close()
	}()

	request := httptest.NewRequest(http.MethodPost, "/files/upload", r)
	request.Header.Add("Content-Type", m.FormDataContentType())

	wr := httptest.NewRecorder()
	env.FilesUpload(wr, request)

	response := wr.Result()
	defer response.Body.Close()

	if wr.Code != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, wr.Code)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	hash := string(data)
	if hash != "" {
		t.Errorf("Expected %s, received %s", "", hash)
	}
}
