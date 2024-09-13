# Simple webserver with Golang

> Moodle-assignment: Create a simple webserver with multiple routes, one of which is static.
> Create a parallell, non-blocking process that takes some time to complete.
> Implement graceful shutdown and test the API routes with curl.

## Setup

1. Install GoLang 1.23 or newer
2. Download repo
3. cd into the folder
4. run command: go run main.go startHTTPServer.go
5. to shut down: CTRL-C

## Use

Once the server is up and running, you can test the API calls either by visiting any of the routes below in your browser, or use curl:

`curl http://localhost:8080/*insert route here*`

Routes:

- http://localhost:8080/
- http://localhost:8080/about
- http://localhost:8080/data
- http://localhost:8080/process
- http://localhost:8080/static
