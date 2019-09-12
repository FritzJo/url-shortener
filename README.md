# URL Shortener
WIP Repository

## Description
This repository contains the code for my personal URL shortening service (like [bit.ly](https://bitly.com/), or [goo.gl](https://goo.gl/)).
My goal is to learn more about building websites with MaterializeCSS and backends with Go, as well as key-value stores.

## How to
``` bash
# Install requirements
go get github.com/gorilla/mux
go get github.com/boltdb/bolt

# Clone the repository 
git clone https://github.com/FritzJo/url-shortener.git

# Run the server
go run *.go
```
Open http://localhost:8080 to use the application.

## Roadmap
* Actually shorten URLs... (I currently use MD5 hashes for shortening, which are longer in many cases)
* Add more server options (custom port, ...) 
* Docker support
* Add more options for shortening
  * one-time-links
  * expiration date 
  * custom short urls
* Improve performance
* Improve code quality
