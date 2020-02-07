# URL Shortener
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d5e298b1d6ce44b78e5829bbb67995c8)](https://www.codacy.com/manual/fritzjo-git/url-shortener?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=FritzJo/url-shortener&amp;utm_campaign=Badge_Grade)

## Description
This repository contains the code for my personal URL shortening service (like [bit.ly](https://bitly.com/), or [goo.gl](https://goo.gl/)).
My goal is to learn more about building websites with MaterializeCSS and backends with Go, as well as key-value stores.

## How to
### Manual deployment
``` bash
# Install requirements
go get github.com/gorilla/mux
go get github.com/boltdb/bolt

# Clone the repository 
git clone https://github.com/FritzJo/url-shortener.git

# Run the server
go run *.go
```

### Docker
``` bash
# Clone the repository 
git clone https://github.com/FritzJo/url-shortener.git

# Build the image
cd url-shortener/
docker build -t urlshortener:v1 . 
docker run -p 8080:8080 urlshortener
```

Open http://localhost:8080/ to use the application.

## Roadmap
* Actually shorten URLs... (I currently use MD5 hashes for shortening, which are longer in many cases)
* Add more server options (custom port, ...) 
* Add more options for shortening
  * one-time-links
  * expiration date 
  * custom short urls
* Improve performance
* Improve code quality
