package handlers

import (
	errors "asciiArt/errors"
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
		defer file.Close()

		_, err = file.WriteString(art)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			templates.ExecuteTemplate(w, "error.html", errors.InternalServer("failed to write to file"))
		}

		// Set the appropriate headers to indicate that it's a file download
		w.Header().Set("Content-Disposition", "attachment; filename=result.txt")
		w.Header().Set("Content-Type", "text/plain")

		// Write the file content to the response
		http.ServeFile(w, r, "result.txt")
	default:
		// http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		templates.ExecuteTemplate(w, "error.html", errors.MethodNotAllowed)
	}
}
