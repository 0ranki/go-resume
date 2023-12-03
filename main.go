package main

import (
	"embed"
	"html/template"
	"log/slog"
	"net/http"
	"slices"
)

//go:embed "static/css/*.css" "templates/*.html"
var static embed.FS
var templates map[string]*template.Template

func main() {
	readConfig("")
	staticFileServer := http.FileServer(http.FS(static))
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/light", home)
	mux.HandleFunc("/dark", home)
	mux.Handle("/static/", staticFileServer)

	slog.Info("Starting go-resume server, listening on port 3000")
	err := http.ListenAndServe(":3000", mux)
	slog.Error(err.Error())
}

func home(w http.ResponseWriter, r *http.Request) {
	acceptedPages := []string{"/", "/light", "/dark"}
	if !slices.Contains(acceptedPages, r.URL.Path) {
		http.NotFound(w, r)
	}
	templateFiles := []string{
		"templates/index.html",
		"templates/metatag.html",
		"templates/githubCorner.html",
	}
	tmpl, err := template.ParseFS(static, templateFiles...)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
	data, err := readConfig("")
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, *data)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}
