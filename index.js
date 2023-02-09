const http = require("http");
const fs = require("fs");
const formidable = require("formidable");

const server = http.createServer((req, res) => {
  // upload files
  if (req.url === "/fileupload") {
    const form = new formidable.IncomingForm();
    form.parse(req, (err, fields, files) => {
      if (err) {
        res.write("Error uploading file");
        res.end();
      }
      const oldPath = files.filetoupload.filepath;
      const folderName = "uploads";
      const newPath = `./${folderName}/${files.filetoupload.originalFilename}`;

      // Check if the folder exists
      fs.promises.access(folderName)
        .catch(() => {
          // If the folder doesn't exist, create it
          return fs.promises.mkdir(folderName);
        })
        .then(() => {
          // Move the file to the folder
          fs.rename(oldPath, newPath, (err) => {
            if (err) {
              console.log(err);
              res.write("Error uploading file");
              res.end();
            }
            res.write("File uploaded and moved!");
            res.end();
          });
        });
    });
  } else {
    res.writeHead(200, { "Content-Type": "text/html" });
    res.write(`
      <form action="/fileupload" method="post" enctype="multipart/form-data">
        <input type="file" name="filetoupload">
        <input type="submit">
      </form>
    `);
    return res.end();
  }
});

const port = process.env.PORT || 3000;
server.listen(port, () => {
  console.log(`Server listening on port ${port}`);
});