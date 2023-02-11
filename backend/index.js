const downloadRoute = require("./router/download.js");
const deleteRoute = require("./router/delete.js");
const uploadRoute = require("./router/upload.js");
const fileUpload = require("express-fileupload");
const getRoute = require("./router/get.js");
const express = require("express");

const app = express();

app.use(fileUpload());
app.use("/", getRoute);
app.use("/", uploadRoute);
app.use("/", deleteRoute);
app.use("/download", downloadRoute);

app.listen(3000, () => {
  console.log("Server listening on port 3000...");
});
