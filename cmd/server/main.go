package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	_ "net/http/pprof"
)

const (
	uploadDir     = "./uploads"
	maxUploadSize = 100 << 20 // 200MB
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	//  limit the max size
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
	}
	defer file.Close()

	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	filename := uuid.New().String() + ext
	filepath := filepath.Join(uploadDir, filename)

	dst, err := os.Create(filepath)
	if err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
	}
	defer dst.Close()

	// stream copy to disk

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Upload Failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Uploaded file successfully: %s\n", filename)

}


func weatherHandler(w http.ResponseWriter, r *http.Request){
	
}

func main() {
	go func() {
		log.Println("pprof running at http://localhost:6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("Server running at http://localhost:8080")
	http.HandleFunc("/weather", weatherHandler)

	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)

}
