const fileUpload = require('express-fileupload');
const express = require('express');
const fs = require('fs')
const app = express();

// Use fileUpload middleware to handle file uploads
app.use(fileUpload());

// Route for downloading a file
app.get('/download/:fileName', (req, res) => {
  const fileName = req.params.fileName;
  res.download(`./files/${fileName}`);
});

// Route for uploading a file
app.post('/upload', (req, res) => {
  if (!req.files || Object.keys(req.files).length === 0) {
    return res.status(400).send('No files were uploaded.');
  }

  // The name of the input field (i.e. "sampleFile") is used to retrieve the uploaded file
  let file = req.files.sampleFile;

  // Use the mv() method to place the file in upload directory (i.e. "./files")
  file.mv(`./files/${file.name}`, function(err) {
    if (err) return res.status(500).send(err);

    res.send('File uploaded!');
  });
});

//list files in the files folder
app.get('/files', (req, res) => {
  files = fs.readdirSync("./files")
  res.send(files)
})


app.listen(3000, () => {
  console.log('Server listening on port 3000...');
});
