const express = require("express");
const fs = require("fs");
const path = require("path");
const connection = require("../config/database");
const router = express.Router();
const filesDirectory = "./files";

const ALLOWED_TYPES = ["image/jpeg", "image/png", "application/pdf", "text/plain"];

const getCurrentDate = () => {
  const date = new Date()
  const year = date.getFullYear()
  const month = ('0' + (date.getMonth() + 1)).slice(-2)
  const day = ('0' + date.getDate()).slice(-2)
  return `${day}-${month}-${year}`
}


router.post("/upload", async (req, res) => {
  if (!req.files || Object.keys(req.files).length === 0) {
    return res.status(400).send("No files were uploaded.");
  }

  const file = req.files.file;
  const fileName = file.name;
  const filePath = path.join(filesDirectory, fileName);
  const tags = req.body.tags;

  if(!ALLOWED_TYPES.includes(file.mimetype)){
    console.log(`file type ${file.mimetype} rejected.`);
    return res.status(415).send("File type not allowed.")
  }

  let sql, data, status;

  try {
    if (fs.readdirSync(filesDirectory).includes(fileName)) {
      await fs.promises.unlink(filePath);
      console.log(`Old file ${fileName} was deleted.`);
      sql =
        "UPDATE files SET path = ?, last_modified = ?, tags = ? WHERE filename = ?";
      data = [filePath, getCurrentDate(), tags, fileName];
      status = "File updated!";
    } else {
      sql =
        "INSERT INTO files (filename, path, last_modified, tags) VALUES (?, ?, ?, ?)";
      data = [fileName, filePath, getCurrentDate(), tags];
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

module.exports = router;
