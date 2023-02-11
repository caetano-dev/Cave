const express = require("express");
const path = require("path");
const router = express.Router();
const filesDirectory = "./files";

router.get("/:fileName", (req, res) => {
  const fileName = req.params.fileName;
  const filePath = path.join(filesDirectory, fileName);
  res.download(filePath);
});

module.exports = router;
