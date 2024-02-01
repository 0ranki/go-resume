package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"
)

//go:embed "static/css/*.css" "templates/*.html"
var static embed.FS
var templates map[string]*template.Template
var configFile string

func main() {
	cfgFile := flag.String("config", "./data/resume.yaml", "Path to the configuration YAML")
	flag.Parse()

	_, err := readConfig(*cfgFile)
	if err != nil {
		slog.Error(fmt.Sprintf("error reading configuration: %s", err.Error()))
		os.Exit(1)
	}
	configFile = *cfgFile
	staticFileServer := http.FileServer(http.FS(static))
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/light", home)
	mux.HandleFunc("/dark", home)
	mux.Handle("/static/", staticFileServer)

	slog.Info("Starting go-resume server, listening on port 3000")
	err = http.ListenAndServe("0.0.0.0:3000", mux)
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
	data, err := readConfig(configFile)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
	// TODO: Theme from browser
	if strings.HasSuffix(r.URL.Path, "/light") || strings.HasSuffix(r.URL.Path, "/dark") {
		urlSlice := strings.Split(r.URL.Path, "/")
		data.Theme = urlSlice[len(urlSlice)-1]
	}
	data.Year = time.Now().Format("2006")
	tmpl.Funcs(templateFunctions())
	err = tmpl.Execute(w, *data)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"html": func(raw string) template.HTML {
			return template.HTML(raw)
		},
	}
}
