package handlers

import (
	asciiArt "asciiArt/asciiArt"
	myErrors "asciiArt/errors"
	"errors"
	"net/http"
	"strings"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		art, err1 := getAsciiArtFromRequest(r)
		if err1 != nil {
			w.WriteHeader(http.StatusBadRequest)
			templates.ExecuteTemplate(w, "error.html", myErrors.BadRequest(err1.Error()))
			return
		}

		err := templates.ExecuteTemplate(w, "result.html", art)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			templates.ExecuteTemplate(w, "error.html", myErrors.InternalServer(err.Error()))
			return
		}

	default:
		// http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		templates.ExecuteTemplate(w, "error.html", myErrors.MethodNotAllowed)
	}

}

func getAsciiArtFromRequest(r *http.Request) (string, error) {
	str := r.FormValue("input")
	font := r.FormValue("font")

	switch strings.ToLower(font) {
	case "standard", "shadow", "thinkertoy":
		// font is valid, do nothing
	default:
		return "", errors.New("invalid font option")
	}

	lines := strings.Split(str, "\r\n")

	var asciiLines []string
	for _, line := range lines {
		if line == "" {
			continue
		}
		asciiLine, err := asciiArt.AsciiLine(line, font)
		if err != nil {
			return "", err
		}
		asciiLines = append(asciiLines, asciiLine...)
	}

	art := asciiArt.PrintAscii(asciiLines)
	return art, nil
}
