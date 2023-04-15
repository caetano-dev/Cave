package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	s "github.com/drull1000/cave/pkgs/structs"
	u "github.com/drull1000/cave/pkgs/utils"
	"github.com/go-chi/chi/v5"
)

// List displays all of the files from the database.
func (rs *FilesResource) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", u.FrontentAddress)

	filesDatabase, err := DBClient.GetAll()
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

// Get displays one file that is chosen by its ID
func (rs *FilesResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", u.FrontentAddress)

	uid := chi.URLParam(r, "date")
	if uid == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	file, err := DBClient.GetByID(id)
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

// FilesCreate is the function that uploads a file.
func (rs *FilesResource) FilesCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", u.FrontentAddress)

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

	id, err := DBClient.Insert(f)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, %d\n", f.Hash, f.Filename, f.Tags, id)
}

// Create is the function that uploads a file.
func (rs *FilesResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", u.FrontentAddress)

	var file s.File

	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filename := file.Filename
	tags := file.Tags

	files_path := filepath.Join(".", "files")
	err = os.MkdirAll(files_path, os.ModePerm)
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
	mw := io.MultiWriter(dst, hash, os.Stdout)

	if _, err := io.Copy(mw, r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashString := hex.EncodeToString(hash.Sum(nil))

	f := s.File{
		Hash:     hashString,
		Filename: filename,
		Filepath: filesystemPath,
		Tags:     tags,
	}

	id, err := DBClient.Insert(f)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, %d\n", f.Hash, f.Filename, f.Tags, id)
}

// Delete is the function that deletes a file by id
func (rs *FilesResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", u.FrontentAddress)

	var body s.RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := body.ID
	file, err := DBClient.GetByID(id)
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

	err = DBClient.DeleteByID(id)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File %s deleted successfully (id: %d)\n", file.Filename, file.ID)
}

// Update ...
func (rs *FilesResource) Update(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, PUT, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	header.Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	var body s.RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, err := DBClient.GetByID(body.ID)

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

// HealthCheck ...
func (rs *FilesResource) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
