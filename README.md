# morty-cat-generator
*web service that pulls names and facts from other services
and responds to user requests with a custom picture-accompanied fact*

## API docs:
* https://rickandmortyapi.com
* https://alexwohlbruck.github.io/cat-facts/

## Service concepts needed:
* mux routing to handle load of requests
* benchmark tests of service loading
* error handling of different http codes
* unmarshalling json responses
* handling unicode/utf-8 characters from uinames
* using go modules to manage project dependencies
* Add a health endpoint of 200 status with json returned of app up and
  access to apis status (red, yellow, green)

## Usage
1. Configure the Inputs services in `main.init()`
1. Download dependencies via `go mod download` (Go 1.11+ required and project
  should exist outside `$GOPATH`). Or, if not using modules:
    1. `go get github.com/gorilla/mux`
    1. `go get github.com/pkg/errors`
1. Build and run the service: `go build; ./joke-name-generator`
1. Call the service to receive joke responses: `curl localhost:8080/`
