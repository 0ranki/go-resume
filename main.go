package main

import (
	"embed"
	"html/template"
	"log/slog"
	"net/http"
)

//go:embed "static/css/*.css"
var static embed.FS

func home(w http.ResponseWriter, r *http.Request) {
	templateFiles := []string{
		"./templates/index.html",
		"./templates/jsonLd.html",
		"./templates/metatag.html",
		"./templates/githubCorner.html",
	}
	tmpl, err := template.ParseFiles(templateFiles...)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

//func css(w http.ResponseWriter, r *http.Request) {
//	staticFileServer := http.FileServer(http.FS(static))
//	mime.AddExtensionType(".css", "text/css")
//	staticFileServer.ServeHTTP(w, r)
//}

func main() {
	staticFileServer := http.FileServer(http.FS(static))
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.Handle("/static/", staticFileServer)

	err := http.ListenAndServe(":3000", mux)
	slog.Error(err.Error())
}
