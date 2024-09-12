# Simple webserver with Golang

> Moodle-assignment: Create a simple webserver with multiple routes, one of which is static.
> Create a parallell, non-blocking process that takes some time to complete.
> Implement graceful shutdown and test the API routes with curl.

## Setup

1. Install GoLang 1.23 or newer
2. Download repo
3. cd into the folder
4. run command: go run main.go
5. to shut down: CTRL-C

## Use

Once the server is up and running, you can test the API calls either by visiting localhost:8080 in your browser, or use curl:
`curl http://localhost:8080/*insert route here*`
