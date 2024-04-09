package handlers

import (
	errors "asciiArt/errors"
	"io"
	"log"
	"net/http"
	"os"
)

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		art, err := getAsciiArtFromRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			templates.ExecuteTemplate(w, "error.html", errors.BadRequest(err.Error()))
			return
		}

		// Create a new file and write the data to it
		file, err := os.Create("result.txt")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			templates.ExecuteTemplate(w, "error.html", errors.InternalServer("failed to create file"))
		}

		_, err = file.WriteString(art)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			templates.ExecuteTemplate(w, "error.html", errors.InternalServer("failed to write to file"))
		}

		file.Close()

		// Open the file for reading
		file, err = os.Open("result.txt")
		if err != nil {
			log.Printf("failed to open file: %v", err)
			return
		}
		defer file.Close()

		// Set the appropriate headers to indicate that it's a file download
		w.Header().Set("Content-Disposition", "attachment; filename=result.txt")
		w.Header().Set("Content-Type", "text/plain")
		// w.Header().Set("Content-Length", string(num))

		// Write the file content to the response
		// http.ServeFile(w, r, "result.txt")
		_, err = io.Copy(w, file)
		if err != nil {
			log.Printf("failed to write file to response: %v", err)
			return
		}

		// Ensure the file is closed before trying to delete it
		file.Close()

		// Delete the file after it's served
		err = os.Remove("result.txt")
		if err != nil {
			log.Printf("failed to delete file: %v", err)
		}

	default:
		// http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		templates.ExecuteTemplate(w, "error.html", errors.MethodNotAllowed)
	}
}
