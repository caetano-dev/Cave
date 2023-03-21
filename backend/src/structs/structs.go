package structs

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

type Config struct {
	Database string
}

type RequestBody struct {
	ID    int64  `json:"id"`
	Value string `json:"value"`
}

type ResponseFileContent struct {
	FileInformation FileDatabase
	Content         string
}
