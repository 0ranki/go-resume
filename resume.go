package main

type Resume struct {
	Theme string `yaml:"theme"`
	Meta  *Meta  `yaml:"meta"`
}

type job struct {
	Company     string   `yaml:"company"`
	Location    string   `yaml:"location"`
	Title       string   `yaml:"title"`
	Period      string   `yaml:"period"`
	Description []string `yaml:"description"`
}

type jobs []job

type education struct {
	Name    string `yaml:"name"`
	City    string `yaml:"city"`
	Degree  string `yaml:"degree"`
	Period  string `yaml:"period"`
	Faculty string `yaml:"faculty"`
}

type educations []education

type skill struct {
	Name  string `yaml:"name"`
	Level string `yaml:"level"`
}

type skills []skill

type social struct {
	Media  string `yaml:"media"`
	UserId string `yaml:"userId"`
	Link   string `yaml:"link"`
}

type socials []social

type Meta struct {
	Language    string `yaml:"language"`
	Author      string `yaml:"author"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Robots      string `yaml:"robots"`
	ThemeColor  string `yaml:"theme-color"`
}
