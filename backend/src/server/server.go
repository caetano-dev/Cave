package server

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/drull1000/notetaking-app/src/database"
	s "github.com/drull1000/notetaking-app/src/structs"
	u "github.com/drull1000/notetaking-app/src/utils"
)

// Env struct is the database env
type Env struct {
	DB *sql.DB
}

// FilesShowAll displays all of the files from the database.
func (env *Env) FilesShowAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", u.FrontentAddress)

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	filesDatabase, err := database.GetAll(env.DB)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var files []s.ResponseFileContent
	for _, fileDB := range filesDatabase {
		content, err := ioutil.ReadFile(fileDB.Filepath)
		if err != nil {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		fileContent := s.ResponseFileContent{
			FileInformation: *fileDB,
			Content:         string(content),
		}

		files = append(files, fileContent)
	}

	allFiles := s.ResponseAllFiles{
		Files: files,
	}
	fileBytes, err := json.Marshal(allFiles)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(fileBytes)
}

// FilesShow displays one file that is chosen by its ID
func (env *Env) FilesShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", u.FrontentAddress)

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	uid := r.FormValue("ID")
	if uid == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	file, err := database.GetByID(env.DB, id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if file.ID == 0 {
		fmt.Fprintf(w, "Not found.")
		return
	}

	fmt.Fprintf(w, "%d, %s, %s, %s, %s\n", file.ID, file.Hash, file.Filename, file.Tags, file.CreatedAt)
}

// FilesCreate is the function that creates a file.
func (env *Env) FilesCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", u.FrontentAddress)

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	filename := "New file"
	tags := strings.Split("", ",")
	files_path := filepath.Join(".", "files")

	err := os.MkdirAll(files_path, os.ModePerm)
	if err != nil {
		http.Error(w, fmt.Sprintf("error retrieving file: %s", err), http.StatusInternalServerError)
		return
	}

	filesystemPath := u.UniqueFilesystemPath(filepath.Join(files_path, filename))

	dst, err := os.Create(filesystemPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	hash := sha256.New()
	hashString := hex.EncodeToString(hash.Sum(nil))

	f := s.File{
		Hash:     hashString,
		Filename: filename,
		Filepath: filesystemPath,
		Tags:     tags,
	}

	id, err := database.Insert(env.DB, f)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, %d\n", f.Hash, f.Filename, f.Tags, id)
}

// FilesDelete is the function that deletes a file.
func (env *Env) FilesDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", u.FrontentAddress)

	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var body s.RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := body.ID
	file, err := database.GetByID(env.DB, id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if file.ID == 0 {
		fmt.Fprintf(w, "file not found")
		return
	}

	//err = os.Remove(filepath.Join("./files", file.Filepath))
	err = os.Remove(file.Filepath)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = database.DeleteByID(env.DB, id)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File %s deleted successfully (id: %d)\n", file.Filename, file.ID)
}

func (env *Env) FileEditContent(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, PUT, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	header.Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var body s.RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, err := database.GetByID(env.DB, body.ID)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if file.ID == 0 {
		fmt.Fprintf(w, "Not found.")
		return
	}

	if err := os.WriteFile(file.Filepath, []byte(body.Value), 0666); err != nil {
		log.Fatal(err)
	}
}

func (env *Env) FileEditName(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, PUT, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	header.Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var body s.RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.UpdateFilename(env.DB, body.ID, body.Value)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func (env *Env) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
