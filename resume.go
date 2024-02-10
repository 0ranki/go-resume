package main

import (
	"fmt"
	"html/template"
)

type Resume struct {
	Theme   string   `yaml:"theme"`
	Meta    *Meta    `yaml:"meta"`
	Profile *Profile `yaml:"profile"`
	Year    string
}

type Profile struct {
	Name         string        `yaml:"name"`
	Title        string        `yaml:"title"`
	City         string        `yaml:"city"`
	Phone        string        `yaml:"phone"`
	Email        string        `yaml:"email"`
	Socials      Socials       `yaml:"socials"`
	Photo        string        `yaml:"photo"`
	Summary      string        `yaml:"summary"`
	Experience   []Job         `yaml:"experience"`
	Education    []Education   `yaml:"education"`
	Skills       []Skill       `yaml:"skills"`
	Projects     any           `yaml:"projects"`
	Languages    []Language    `yaml:"languages"`
	Certificates []Certificate `yaml:"certificates"`
}

func (p *Profile) GetTagWidth(t string) string {
	l := 0
	switch t {
	case "skills":
		for _, skill := range p.Skills {
			if len(skill.Name) > l {
				l = len(skill.Name)
			}
		}
	case "languages":
		for _, skill := range p.Languages {
			if len(skill.Name) > l {
				l = len(skill.Name)
			}
		}
	}
	return fmt.Sprintf("%dch", l)
}

//func (job *Job) html(raw string) template.HTML {
//	return template.HTML(raw)
//}

type Job struct {
	Company     string   `yaml:"company"`
	Location    string   `yaml:"location"`
	Title       string   `yaml:"title"`
	Period      string   `yaml:"period"`
	Description []string `yaml:"description"`
}

type Education struct {
	Name    string `yaml:"name"`
	Degree  string `yaml:"degree"`
	Faculty string `yaml:"faculty"`
	City    string `yaml:"city"`
	Period  string `yaml:"period"`
}

type Skill struct {
	Name  string `yaml:"name"`
	Level int    `yaml:"level"`
}

func _visualLevel(level int) template.HTML {
	tagString := ""
	if level > 5 {
		level = 5
	} else if level < 1 {
		level = 1
	}
	for i := 0; i < level; i++ {
		tagString += "<span class=\"tag is-info is-success is-size-6\" style=\"color: black\">☆</span>"
	}
	return template.HTML(tagString)
}
func (skill *Skill) VisualLevel(level int) template.HTML {
	return _visualLevel(level)
}

func (language *Language) VisualLevel(level int) template.HTML {
	return _visualLevel(level)
}

type Social struct {
	Media  string `yaml:"media"`
	UserId string `yaml:"userId"`
	Link   string `yaml:"link"`
}

type Socials []Social

type Meta struct {
	Language     string `yaml:"language"`
	Author       string `yaml:"author"`
	Title        string `yaml:"title"`
	Description  string `yaml:"description"`
	Robots       string `yaml:"robots"`
	ThemeColor   string `yaml:"theme-color"`
	GithubCorner bool   `yaml:"github-corner"`
}

type Language struct {
	Name  string `yaml:"name"`
	Level int    `yaml:"level"`
}

type Certificate struct {
	Name string `yaml:"name"`
}
