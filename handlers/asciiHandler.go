package handlers

import (
	asciiArt "asciiArt/asciiArt"
	"net/http"
	"strings"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		str := r.FormValue("input")
		font := r.FormValue("font")

		switch strings.ToLower(font) {
		case "standard", "shadow", "thinkertoy":
			// font is valid, do nothing
		default:
			http.Error(w, "Invalid font option", http.StatusBadRequest)
			return
		}

		lines := strings.Split(str, "\r\n")

		var asciiLines []string
		for _, line := range lines {
			if line == "" {
				continue
			}
			asciiLine, err := asciiArt.AsciiLine(line, font)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			asciiLines = append(asciiLines, asciiLine...)
		}

		art := asciiArt.PrintAscii(asciiLines)

		err := templates.ExecuteTemplate(w, "result.html", art)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}
