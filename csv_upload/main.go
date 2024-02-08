package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error reading file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		var data []map[string]string

		headers, err := reader.Read()
		if err != nil {
			http.Error(w, "Error reading CSV headers", http.StatusInternalServerError)
			return
		}

		for {
			record, err := reader.Read()0
			if err == io.EOF {
				break
			}
			if err != nil {
				http.Error(w, "Error reading CSV record", http.StatusInternalServerError)
				return
			}

			rowData := make(map[string]string)
			for i, value := range record {
				rowData[headers[i]] = value
			}
			data = append(data, rowData)
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Serve static files (HTML, CSS, JS)
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Handle the file upload
	http.HandleFunc("/upload", uploadHandler)

	port := 9513
	fmt.Printf("Server listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
