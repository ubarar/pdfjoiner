const express = require('express')
const app = express()
const multer = require('multer')
const { exec } = require('child_process');
const fs = require('fs');
const https = require('https');

// config storage for file uploads
var storage = multer.diskStorage({
  destination: function (req, file, cb) {
    cb(null, '/app/storage/input')
  },
  filename: function (req, file, cb) {
    cb(null, file.fieldname + '-' + Date.now())
  }
})

var upload = multer({ storage: storage })

function clearInputs(callback) {
  exec("rm -f /app/storage/input/*", (err, stdout, stderr) => {
    if (err)
      callback(err);
    else
      callback();
  });
}

function clearOutputs(callback) {
  exec("rm -f /app/storage/output/*", (err, stdout, stderr) => {
    if (err)
      callback(err);
    else
      callback();
  });
}

app.get('/status', (req, res) => res.send('ok'));

app.get('/', (req, res) => res.sendFile('/app/index.html'));

app.get('/clear', (req, res) => clearInputs((e) => res.send(e ? e : 'ok')));

app.post('/uploadmultiple', upload.array('myFiles', 12), (req, res, next) => {

  let files = "";
  fs.readdirSync('/app/storage/input').forEach(file => {
    files += '/app/storage/input/' + file + " ";
  });

  torun = 'pdfunite ' + files + ' /app/storage/output/output.pdf && rm -f /app/storage/input/*';
  console.log("about to run " + torun);

  exec(torun, (err, stdout, stderr) => {
    if (err) {
      res.send("Something went wrong, please try again. " + err);
    } else {
      console.log("pdf unite worked successfully")
    }
  });
  res.redirect('/output');
})

app.post('/uploadcompress', upload.array('myFiles', 12), (req, res, next) => {
  console.log('running upload compress')

  let file = "";
  fs.readdirSync('/app/storage/input').forEach(file => {
    file = '/app/storage/input/' + file;
  });

  torun = 'gs -sDEVICE=pdfwrite -dCompatibilityLevel=1.4 -dPDFSETTINGS=/ebook -dNOPAUSE -dQUIET -dBATCH -sOutputFile=/app/storage/output/output.pdf ' + file;
  console.log("about to run " + torun);

  exec(torun, (err, stdout, stderr) => {
    if (err) {
      res.send("Something went wrong, please try again. " + err);
    } else {
      console.log("pdf unite worked successfully")
    }
  });
  res.redirect('/output');
})


app.get('/output', (req, res) => {
  clearInputs((err) => {
    if (err) {
      console.error(err);
      res.send("Something went wrong, please try again. " + err);
    } else {
      res.sendFile('/app/storage/output/output.pdf');
      clearOutputs(() => {});
    }
  });
})

https.createServer({
  key: fs.readFileSync('/app/key.pem'),
  cert: fs.readFileSync('/app/cert.pem')
}, app).listen(8080, () => console.log('server start'));
