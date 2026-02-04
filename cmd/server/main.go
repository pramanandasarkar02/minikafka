package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"

	_ "net/http/pprof"

	_ "github.com/mattn/go-sqlite3"
)

const (
	uploadDir     = "./uploads"
	maxUploadSize = 100 << 20 // 200MB
)

type WeatherInfo struct {
	ID          int32     `json:"id"`
	Lon         float32   `json:"lon"`
	Lat         float32   `json:"lat"`
	Temperature float32   `json:"temperature"`
	Pressure    float32   `json:"pressure"`
	Timestamp   time.Time `json:"timestamp"`
}

var DB *sql.DB

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./weather.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS weather(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		lon FLOAT,
		lat FLOAT,
		temperature FLOAT,
		presure FLOAT,
		timestamp TIMESTAMP
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%s: %v", sqlStmt, err)
	}
}


func weatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var wf WeatherInfo

		err := json.NewDecoder(r.Body).Decode(&wf)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		log.Println(wf)

		_, err = DB.Exec(
			`INSERT INTO weather (lon, lat, temperature, presure, timestamp)
     VALUES (?, ?, ?, ?, ?)`,
			wf.Lon, wf.Lat, wf.Temperature, wf.Pressure, wf.Timestamp,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Weather Info inserted successfully"})
	} else if r.Method == http.MethodGet {

		rows, err := DB.Query(`select * from weather`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var wfs []WeatherInfo
		for rows.Next() {
			var wf WeatherInfo
			if err := rows.Scan(
				&wf.ID,
				&wf.Lon,
				&wf.Lat,
				&wf.Temperature,
				&wf.Pressure,
				&wf.Timestamp,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			wfs = append(wfs, wf)
		}

		json.NewEncoder(w).Encode(wfs)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

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

func main() {

	go func() {
		log.Println("pprof running at http://localhost:6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	initDB()
	defer DB.Close()

	http.HandleFunc("/weather", weatherHandler)
	http.HandleFunc("/upload", uploadHandler)
	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
