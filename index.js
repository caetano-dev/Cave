require("dotenv").config();
const fileUpload = require("express-fileupload");
const express = require("express");
const mysql = require("mysql2");
const fs = require("fs");
const path = require("path");
const app = express();

const host = process.env.DB_HOST;
const user = process.env.DB_USER;
const password = process.env.DB_PASSWORD;
const database = process.env.DB_NAME;

const connection = mysql.createConnection({
  host,
  user,
  password,
  database,
});

connection.connect((err) => {
  if (err) throw err;
  console.log("Connected to the MySQL database.");
});

const filesDirectory = "./files";
console.log(filesDirectory)

// Use fileUpload middleware to handle file uploads
app.use(fileUpload());

// Route for downloading a file
app.get("/download/:fileName", (req, res) => {
  const fileName = req.params.fileName;
  const filePath = path.join(filesDirectory, fileName);
  res.download(filePath);
});

// Route for uploading a file
app.post("/upload", async (req, res) => {
  if (!req.files || Object.keys(req.files).length === 0) {
    return res.status(400).send("No files were uploaded.");
  }

  const file = req.files.file;
  const fileName = file.name;
  const filePath = path.join(filesDirectory, fileName);
  const tags = req.body.tags;
  let sql, data, status;

  try {
    if (fs.readdirSync(filesDirectory).includes(fileName)) {
      await fs.promises.unlink(filePath);
      console.log(`Old file ${fileName} was deleted.`);
      sql =
        "UPDATE files SET path = ?, last_modified = ?, tags = ? WHERE filename = ?";
      data = [filePath, new Date(), tags, fileName];
      status = "File updated";
    } else {
      sql =
        "INSERT INTO files (filename, path, last_modified, tags) VALUES (?, ?, ?, ?)";
      data = [fileName, filePath, new Date(), tags];
      status = "File uploaded!";
    }

    await file.mv(filePath);
    await connection.promise().query(sql, data);
    console.log("File information saved/updated in the database.");
    return res.send(status);
  } catch (err) {
    console.error(err);
    return res.status(500).send(err);
  }
});

app.delete("/delete", async (req, res) => {
  const file = req.files.file;
  const fileName = file.name;
  const filePath = path.join(filesDirectory, fileName);
  const tags = req.body.tags;
  let sql, data, status;

  try {
    if (fs.readdirSync(filesDirectory).includes(fileName)) {
      await fs.promises.unlink(filePath);
      console.log(`File ${fileName} was deleted.`);
      sql =
        "DELETE FROM files WHERE filename = ?";
      data = [fileName];
      status = "File deleted";
      await connection.promise().query(sql, data);
      console.log("File deleted from the database.");
      return res.send(status);
    } else {
      res.send("File does not exist.");
    }
  } catch (err) {
    console.error(err);
    return res.status(500).send(err);
  }
});
//list files in the files folder
app.get("/files", async (req, res) => {
  try {
    const [rows] = await connection
      .promise()
      .query("SELECT filename FROM files");
    const files = rows.map(({ filename }) => filename);
    console.log("Files retrieved from the database.");
    return res.send(files);
  } catch (err) {
    console.error(err);
    return res.status(500).send(err);
  }
});

app.listen(3000, () => {
  console.log("Server listening on port 3000...");
});
