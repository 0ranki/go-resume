package main

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
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
