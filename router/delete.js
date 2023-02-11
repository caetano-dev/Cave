const express = require("express");
const fs = require("fs");
const path = require("path");
const connection = require("../config/database");
const router = express.Router();
const filesDirectory = "./files";

router.delete("/delete", async (req, res) => {
  const file = req.files.file;
  const fileName = file.name;
  const filePath = path.join(filesDirectory, fileName);
  let sql, data, status;

  try {
    if (fs.readdirSync(filesDirectory).includes(fileName)) {
      await fs.promises.unlink(filePath);
      console.log(`File ${fileName} was deleted.`);
      sql = "DELETE FROM files WHERE filename = ?";
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

module.exports = router;
