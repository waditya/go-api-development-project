# go-api-development-project 

## Web Server and API 

### Web Server

* The primary role of a web server is to serve web pages for a website

* If you want to host your web application on the internet, in many cases you will need a web server.

* A web page can be rendered from a single HTML file, or a complex assortment of resources fitted 
together. (like JS, scripts, CSS, files, etc)

### Request-Response Model 

Client sends the data and server responds to the request with a response

### use cases for a web server 

* Servers HTML, CSS, and JS files
* Servers images and videos
* Handles HTTP error messaging
* Handles user requests, often concurrently
* Directs URL matching and rewriting
* Processes and servers dynamic content 

### API 

* API stands for **application programming interface**, which is a set of definitions and protocols 
for building and integrating application software.

* APIs let your product or service communnicate with other producst and services without having to 
know how they are implemented. Example - If yor application wants to have navigation feature, it can 
use Google Maps API. It doesn't have ti know how Google Maps API is implemented.

* This can simplify application development and save time and money.

## Creating an API Server

* We use ListenAndServe function for the http package to create a server and make it listen to a specific 
port. ListenAndServe listens on a TCP address and port and then calls the Server function within http package.
The serve accepts incoming HTTP requests on the listener l, and creates a new service goroutine for each request.
The service goroutines read requests and then call handler to reply to them.
The handler is typically nil, in whih case DefaultServeMux is used.

DefaultServeMux is the default ServeMux used by Serve.

ServeMux is Go struct. This struct is an  HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

* A Web requets is used to access or modify a resource on the server. To convey the details of the resource,
we use endpoints. Endpoints helps client to convey which specific resource it wants access to.

* When a client requests for access to a specific endpoint, that request has to be handled/fulfilled by some entity. This entity is the handler. There can be many different endpoints that needed to be served. Different 
handlers can serve different endpoints. In order to tie endpoint requested by client request with correct handler
that can serve the request, we have HandlerFunc function in the http package. It connects together endpoint 
with the handler.

## HTTP Verbs 

* HTTP defines a set of **request methods** to indicate the desired action to be dtored for a given resource

* Although they can also be nounds, these requets methods are sometimes referred to as HTTP verbs

### GET - 

* The GET method requests a representation of the specified resource
* Reuqests using GET should only retrieve data 
* GET represents the READ part of the CRUD operation

### POST - 

* The POST method submits an entity to the specified resource, often causing a change in state or 
side effects on the server

* POST represents the CREATE part of the CRUD operation 

### PUT - 

* The PUT method replaces all current representatios of the target resource with the
request payload

### DELETE - 

* The DELETE method deletes the specified resource

### PATCH - 

* The PATCH method replces a specific atribute of the target resource with

