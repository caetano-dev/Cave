const connection = require("../config/database");
const express = require("express");
const router = express.Router();

router.get("/files", async (req, res) => {
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

module.exports = router;
