const fs = require('fs');
require('dotenv').config();
const mysql = require('mysql2');
const host = process.env.DB_HOST;
const user = process.env.DB_USER;
const password = process.env.DB_PASSWORD;
const database = process.env.DB_NAME;

const connection = mysql.createConnection({
  host: host,
  user: user,
  password: password,
  database: database 
});


const sql = fs.readFileSync('./files.sql').toString();

connection.query(sql, function(err, result) {
  if (err) throw err;
  console.log("Table created");
});

connection.end();
