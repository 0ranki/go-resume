package main

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"path"
	"strings"
)

func readConfig(configFilePath string) (resumeConfig *Resume, err error) {
	if configFilePath == "" {
		configFilePath = "./data/resume.yaml"
	}
	_, err = os.Stat(configFilePath)
	if err != nil {
		return
	}
	configBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(configBytes, &resumeConfig)
	if err != nil {
		slog.Error(err.Error())
	}
	return
}

func photoIsLocal(photoLocation string) bool {
	return !strings.HasPrefix(photoLocation, "http")
}

func getPhotoPaths(cfg *Resume) (dir, relDir string) {
	photo := cfg.Profile.Photo
	if strings.HasPrefix(photo, ".") {
		photo = strings.TrimPrefix(photo, ".")
	}
	dir, _ = path.Split(photo)
	relDir = "." + dir
	return
}
