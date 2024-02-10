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
var basePath string
var strs map[string]map[string]string

func main() {
	loadStrings()
	cfgFile := flag.String("config", "./data/resume.yaml", "Path to the configuration YAML")
	flag.Parse()

	cfg, err := readConfig(*cfgFile)
	if err != nil {
		slog.Error(fmt.Sprintf("error reading configuration: %s", err.Error()))
		os.Exit(1)
	}
	basePath = "/"
	if cfg.Basepath != "" {
		basePath = cfg.Basepath
		if !strings.HasSuffix(basePath, "/") {
			basePath = basePath + "/"
		}
	}
	configFile = *cfgFile
	staticFileServer := http.FileServer(http.FS(static))
	mux := http.NewServeMux()
	mux.HandleFunc(basePath, home)
	mux.HandleFunc(basePath+"light", home)
	mux.HandleFunc(basePath+"dark", home)
	mux.Handle(basePath+"static/", http.StripPrefix(basePath, staticFileServer))
	if photoIsLocal(cfg.Profile.Photo) {
		dir, relDir := getPhotoPaths(cfg)
		mux.Handle(dir, http.StripPrefix(dir, http.FileServer(http.Dir(relDir))))
	}
	slog.Info("Starting go-resume server, listening on port 3000")
	err = http.ListenAndServe("0.0.0.0:3000", mux)
	slog.Error(err.Error())
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Client: " + r.Header.Get("X-Forwarded-For") + " Request: " + r.URL.Path)
	acceptedPages := []string{basePath, basePath + "light", basePath + "dark"}
	if !slices.Contains(acceptedPages, r.URL.Path) {
		http.NotFound(w, r)
	}
	templateFiles := []string{
		"templates/index.html",
		"templates/metatag.html",
		"templates/githubCorner.html",
	}
	var err error
	tmpl := template.New("index").Funcs(templateFunctions())
	tmpl, err = tmpl.ParseFS(static, templateFiles...)
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
		"translate": func(label, lang string) string {
			if strs[lang][label] == "" {
				return label + " (no translation)"
			}
			return strs[lang][label]
		},
	}
}
