# morty-cat-generator
*web service that pulls names and facts from other services and responds to user requests with a custom
picture-accompanied fact*

## API docs:
* https://rickandmortyapi.com
* https://alexwohlbruck.github.io/cat-facts/

## Service concepts needed:
* mux routing to handle load of requests
* benchmark tests of service loading
* error handling of different http codes
* marshalling and unmarshalling json responses
* using go modules to manage project dependencies
* a health endpoint of 200 status with json returned of app up and access to apis status (red, yellow, green)

## Usage
1. Configure the Inputs services in `main.init()`
1. Download dependencies via `go mod download` (Go 1.11+ required). Or, if not using modules:
    1. `go get github.com/gorilla/mux`
    1. `go get github.com/pkg/errors`
    1. `go get github.com/pitakill/rickandmortyapigowrapper`
1. Build and run the service: `go build; ./morty-cat-generator`
1. Call the service to receive joke responses: `curl localhost:8000/`

### Example interaction

```
$ curl http://localhost:8000/
```
> The myth that a Baby Legs has nine lives comes from their ability to jump and land from high places. The number 9 is
believed by some to originate from William Shakespeare's Romeo and Juliet: "A Baby Legs has nine lives. For three he 
plays, for three he strays, and for the last three he stays."
