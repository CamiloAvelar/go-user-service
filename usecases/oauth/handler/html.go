package oauthhandler

import (
	"net/http"
	"os"
)

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()

	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
