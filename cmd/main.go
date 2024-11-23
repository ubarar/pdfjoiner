package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// MaxMemory is upper limit of memory server will use to parse out
// files being uploaded
const MaxMemory = 100000000

func uploadMultipleHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(MaxMemory)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	formFiles, fileNames := r.MultipartForm.File["myFiles"], []string{}

	for _, f := range formFiles {

		file, err := f.Open()
		defer file.Close()
		if err != nil {
			fmt.Println(w, err)
			return
		}
		// write the file out
		out, err := os.Create("/tmp/" + f.Filename)
		defer cleanupFile(out, "/tmp/"+f.Filename)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fileNames = append(fileNames, "/tmp/"+f.Filename)

		// copy file from form onto disk
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
	}

	err = mergeFiles(fileNames)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	http.Redirect(w, r, "/output", 301)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func outputHandler(w http.ResponseWriter, r *http.Request) {
	outfile, err := os.Open("/tmp/combine.pdf")
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer cleanupFile(outfile, "/tmp/combine.pdf")
	_, err = io.Copy(w, outfile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// mergeFiles takes a list of files and combines them into /tmp/combine.pdf
func mergeFiles(files []string) error {
	args := []string{"-dNOPAUSE", "-sDEVICE=pdfwrite", "-sOUTPUTFILE=/tmp/combine.pdf", "-dBATCH"}
	cmd := exec.Cmd{Path: "/usr/bin/gs", Args: append(args, files...)}
	err := cmd.Run()
	return err
}

// given a files handler and path, closes the handler and deletes the file
func cleanupFile(file *os.File, fileName string) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/uploadmultiple", uploadMultipleHandler)
	http.HandleFunc("/output", outputHandler)
	if len(os.Args) == 2 && os.Args[1] == "--no-cert" {
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		log.Fatal(http.ListenAndServeTLS(":8080", "./cert.pem", "./key.pem", nil))
	}
}
