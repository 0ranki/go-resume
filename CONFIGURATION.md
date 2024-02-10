## YAML Configuration

The configuration file is read for each request, there's no need to stop the server
to make edits.

### `basepath`
Defaults to empty, set if you want to serve under a subpath
```YAML
basepath: "/cv"
```

### `theme`
Sets the default theme that is loaded when people visit your CV
```YAML
theme: "dark"
```

### `meta` section
These values correspond to tags in the `<head> section.
- `meta.title` sets the `<title>` meta tag content
- `meta.github-corner` enables/disables the GitHub link to this repo in the top-right corner
- `meta.favicon` is still WIP, will point to a local or remote favicon assett
- `meta.language` sets the language for the headings in the document (e.g. SUMMARY), also the `<meta name="language">` tag
- Others correspond to `<meta name="tag name">`
```YAML
meta:
  robots: "index, follow"
  language: "en-EN"
  author: "Mickey Mouse"
  theme-color: "#bd93f9"
  title: ""
  description: ""
  favicon: ""
  github-corner: false
```

## `profile` (the main content definitions)
### Main info fields
Other info fields are quite self-explanatory, except
- `profile.photo`
  - Can be a local file. In this case the path is relative
    to the photo to use, meaning `./data/photo/`, `data/photo/`
    and `/data/photo/` all mean a folder under the directory the
    binary is in. **NOT AN ABSOLUTE OS PATH**
  - Can be a URL to a remote asset
- `profile.summary` Can be written on multiple lines by using the
  standard  YAML `>`, `|` etc selectors, and can additionally
  contain HTML. For example multiple paragraphs using `<p></p>`
```YAML
profile:
  name: "Mickey Mouse"
  title: "Professional job applicant"
  city: "City, Country"
  phone: "+358 xx xxx xxxx"
  email: "your@email.com"
  photo: "./data/photo/Mickey_Mouse_Steamboat_Willie.jpg"
  summary: >-
    Insert summary here. Multiple
    lines can be used, but line breaks are determined
    by the renderer
```

### `profile.experience`
A YAML List of previous jobs. Noteworthy is the description,
which is also a list of items, corresponding to e.g. different
tasks or responsibilities
```YAML
profile:
  experience:
    - company: Company 1
      location: City, country
      title: Professional employee
      period: "03/2023 - "
      description:
        - Each item in this list
        - Will be a new item
        - In the CV
    - company: Company 2
      location: City, country
      title: Professional employee
      period: "03/2022 - 03/2023"
      description:
        - Each item in this list
        - Will be a new item
        - In the CV
    - company: Company 3
      location: City, country
      title: Professional employee
      period: "03/2021 - 03/2022"
      description:
        - Each item in this list
        - Will be a new item
        - In the CV
```

### `profile.education`
A YAML list of prior education history items. These may not fit
every imaginable education, but the fields are not named in the
resulting CV, so they can be used a bit creatively.
- `profile.education.degree` is the main title
- `profile.education.name` is the subtitle
```YAML
profile:
  education:
    - degree: B.S. in Job interviews
      name: Job interviewee
      faculty: School for Job applicants
      city: City, Country
      period: 1/1996 - 12/1996
    - degree: Undergraduate of CV Authoring
      name: Technical writer
      faculty: School for Job applicants
      city: City, Country
      period: 1/1995 - 12/1995
```

### `profile.skills`, `profile.languages`
- name and the amount of starts to give each (1-5)
```YAML
profile:
  skills:
    - name: Skill 1
      level: 1
    - name: Skill 2
      level: 2
    - name: Skill 3
      level: 3
    - name: Skill 4
      level: 4
    - name: Skill 5
      level: 5
  languages:
    - name: Language 1
      level: 5
    - name: Language 2
      level: 4
    - name: Language 1
      level: 3
    - name: Language 2
      level: 2
```

### `profile.certificates`
Other mentionable courses, etc that don't quite fit the others, for example
driver's license, any certificate like CKAD
```YAML
profile:
  certificates:
    - name: Certified Kubernetes Application Developer
    - name: Driver's License (commercial)
```
