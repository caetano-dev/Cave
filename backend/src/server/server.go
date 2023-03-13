package server

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/drull1000/notetaking-app/src/database"
)

// Env struct is the database env
type Env struct {
	DB *sql.DB
}

type responseFileContent struct {
	FileInformation database.FileDatabase
	Content         string
}

// FilesShowAll displays all of the files from the database.
func (env *Env) FilesShowAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	files, err := database.GetAll(env.DB)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fileBytes, err := json.Marshal(files)
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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

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
		fmt.Fprintf(w, "not found")
		return
	}

	fmt.Fprintf(w, "%d, %s, %s, %s, %s\n", file.ID, file.Hash, file.Filename, file.Tags, file.CreatedAt)
}

// FilesUpload is the function that uploads a file.
func (env *Env) FilesUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	newFilename := r.FormValue("newFilename")
	tags := r.FormValue("tags")
	splitTags := strings.Split(tags, ",")

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, fmt.Sprintf("error retrieving file: %s", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	files_path := filepath.Join(".", "files")
	err = os.MkdirAll(files_path, os.ModePerm)
	if err != nil {
		http.Error(w, fmt.Sprintf("error retrieving file: %s", err), http.StatusInternalServerError)
		return
	}

	if newFilename == "" {
		newFilename = handler.Filename
	}

	// Create `files` if not exist
	filesystemPath := filepath.Join(files_path, handler.Filename)
	dst, err := os.Create(filesystemPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	hash := sha256.New()
	mw := io.MultiWriter(dst, hash, os.Stdout)

	if _, err := io.Copy(mw, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashString := hex.EncodeToString(hash.Sum(nil))

	f := database.File{
		Hash:     hashString,
		Filename: newFilename,
		Filepath: filesystemPath,
		Tags:     splitTags,
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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	if r.Method != "DELETE" {
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
		fmt.Fprintf(w, "file not found")
		return
	}

	err = os.Remove(filepath.Join("./files", file.Filename))
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

// FileContent function opens a file and returns its content as a response
func (env *Env) FileContent(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	header.Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		ID int64 `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if fmt.Sprint(requestBody.ID) == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(fmt.Sprint(requestBody.ID), 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// get file data from database
	f, err := database.GetByID(env.DB, id)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	// open file
	file, err := os.Open(f.Filepath)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	byteContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	content := string(byteContent[:])

	response := &responseFileContent{
		FileInformation: f,
		Content:         content,
	}

	// Set the content type and write the content to the response
	fileBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(fileBytes)
}

func (env *Env) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
