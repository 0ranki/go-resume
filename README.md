# Go-resume ![Static Badge](https://img.shields.io/badge/WIP-yellow)

A dynamic resume site server using [Bulma CSS](https://github.com/jgthms/bulma) written in Go.
Get up and running easily by defining the content in a single YAML file.

If you have a public server, can serve the website on port 3000 (HTTP). Or run locally only to
generate a PDF.

### Live example: [https://oranki.net/cv](https://oranki.net/cv)

### Screenshots:
#### Desktop
![kuva](https://github.com/0ranki/go-resume/assets/50285623/4d918f31-a922-45d7-b2e6-67d4a44ebedb)

#### Mobile:
![kuva](https://github.com/0ranki/go-resume/assets/50285623/3d0252c2-2450-47b2-b2c4-bcf448db9f75)


### PDF
The resume can be exported fairly well to PDF by using your browser's built-in printing function.
Just configure your resume.yaml, run the server locally, open the page in your browser and press
"Print to PDF". Feedback is welcome, the result is not perfect, but subjectively as good as most
of the free Word templates I've come across.

**Note:** Printing the dark version doesn't quite work yet

## Usage (see [CONFIGURATION.md](CONFIGURATION.md))
- Download one of the pre-built binaries from releases or build yourself
- Update YAML (see separate document for details)
- Start the server
- The content is updated live, so you can make edits while the server is
  running

### Reverse proxy
Running behind a reverse proxy is simple, e.g. for NGINX:
```
server {
  listen 80;
  listen [::]:80;
  server_name my-cv.example.com;
  location / {
    return 301 https://$host$request_uri;
  }
}
server {
  listen 443 ssl http2;
  listen [::]:443 ssl http2;
  server_name my-cv.example.com;
  ssl_certificate "/path/to/cert";
  ssl_certificate_key "/path/to/key";
  location / {
    proxy_set_header X-Forwarded-For $real_ip; # Not really necessary
    proxy_pass http://localhost:3000;
  }
}
```
If you want to run under a subpath, set `basepath` in the config YAML and change the
location block to the same, e.g. `basepath: "/cv"` in YAML, and
```
  ...
  location /cv {
    proxy_pass http://localhost:3000;
  }
  ...
```

### Planned, but still missing features:
- Socials list
- Contact / contact info so it won't cause a spam wave


## Special thanks
Shamelessly adapted to Go from [mazipan's](https://github.com/mazipan) [bulma-resume-template](https://github.com/mazipan/bulma-resume-template).
