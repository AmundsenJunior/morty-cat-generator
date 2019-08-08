# joke-name-generator
*coding challenge - web service that pulls names and jokes from other services
and responds to user requests with a custom joke*

## API docs:
* https://uinames.com/
* http://www.icndb.com/api/

## service concepts needed:
* ~~mux routing to handle load of requests~~
* benchmark tests of service loading
* error handling of different http codes
* ~~unmarshalling json responses~~
* ~~handling unicode/utf-8 characters from uinames~~
* ~~using go modules to manage project dependencies~~
* Add a health endpoint of 200/500 status with json returned of app up and
  access to apis status (red, yellow, green)

## requirements:
* ~~README.md of running and using service~~
* TODO comments of pieces missing
* return zip of project by email

## Issues

* uinames.com was unavailable at the time of my challenge (`ERR_NAME_NOT_RESOLVED`),
  so I ran the PHP server locally for the time being, via https://github.com/thm/uinames:
  1. `sudo apt install -y apache2 php libapache2-mod-php`
  1. `cd ~/dev`
  1. `git clone git@github.com:thm/uinames.git`
  1. `ln -s /home/user/dev/uinames/uinames.com /var/www/html/uinames`
  1. `curl -X GET http://localhost/uinames/api/`

## Usage
1. Configure the Inputs services in `main.init()`
1. Download dependencies via `go mod download` (Go 1.11+ required)
1. Build and run the service: `go build; ./joke-name-generator`
1. Call the service to receive joke responses: `curl localhost:8080/`