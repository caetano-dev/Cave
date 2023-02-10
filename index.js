require("dotenv").config();
const fileUpload = require("express-fileupload");
const express = require("express");
const mysql = require("mysql2");
const fs = require("fs");
const app = express();
const host = process.env.DB_HOST;
const user = process.env.DB_USER;
const password = process.env.DB_PASSWORD;
const database = process.env.DB_NAME;

const connection = mysql.createConnection({
  host: host,
  user: user,
  password: password,
  database: database,
});

connection.connect(function (err) {
  if (err) throw err;
  console.log("Connected to the MySQL database.");
});

// Use fileUpload middleware to handle file uploads
app.use(fileUpload());

// Route for downloading a file
app.get("/download/:fileName", (req, res) => {
  const fileName = req.params.fileName;
  res.download(`./files/${fileName}`);
});

// Route for uploading a file
app.post("/upload", (req, res) => {
  if (!req.files || Object.keys(req.files).length === 0) {
    return res.status(400).send("No files were uploaded.");
  }
  // The name of the input field (i.e. "file") is used to retrieve the uploaded file
  let file = req.files.file;
  let tags = req.body.tags;
  let fileName = file.name;

  if (fs.readdirSync("./files/").includes(file.name)) {
    fs.unlink(`./files/${fileName}`, (err) => {
      if (err) return res.status(500).send(err);
      console.log(`Old file ${fileName} was deleted.`);
    });
    // Use the mv() method to place the file in upload directory (i.e. "./files")
    file.mv(`./files/${fileName}`, function (err) {
      if (err) return res.status(500).send(err);
      console.log(`File ${fileName} was updated.`);
      res.send("File updated");
      connection.query(
        "UPDATE files SET path = ?, last_modified = ?, tags = ? WHERE filename = ?",
        [`./files/${file}`, new Date(), tags, fileName],
        function (error, results) {
          if (error) throw error;
          console.log("File information updated in the database.");
        }
      );
    });
  } else {
    // Use the mv() method to place the file in upload directory (i.e. "./files")
    file.mv(`./files/${file.name}`, function (err) {
      if (err) return res.status(500).send(err);
      res.send("File uploaded!");
      connection.query(
        "INSERT INTO files (filename, path, last_modified, tags) VALUES (?, ?, ?, ?)",
        [file.name, `./files/${file.name}`, new Date(), tags],
        function (error, results) {
          if (error) throw error;
          console.log("File information saved in the database.");
        }
      );
    });
  }
});

//list files in the files folder
app.get("/files", (req, res) => {
    connection.query(
      "SELECT filename FROM files",
      function (error, results) {
        if (error) throw error;
        console.log("Files retrieved from the database.")
        let files = [];
        for (let i=0; i<results.length; i++) {
          files.push(results[i].filename);
        }
        res.send(files);
      }
    );
});

app.listen(3000, () => {
  console.log("Server listening on port 3000...");
});
