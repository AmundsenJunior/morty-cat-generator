# joke-name-generator
*coding challenge - web service that pulls names and jokes from other services and responds to user requests with a custom joke*

## API docs:
* https://uinames.com/
* http://www.icndb.com/api/

## service concepts needed:
* mux routing to handle load of requests
* benchmark tests of service loading
* error handling of different http codes
* marshalling and unmarshalling json responses
* handling unicode/utf-8 characters from uinames
* using go modules to manage project dependencies
* Add a health endpoint of 200/500 status with json returned of app up and access to apis status (red, yellow, green)

## requirements:
* README.md of running and using service
* TODO comments of pieces missing
r* eturn zip of project by email
