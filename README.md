# Go-REST

Create a restful api with Golang.<br />
The main idea behind this project was to gain some knowledge about implementing 
a rest-api in Golang.

## Setup
Clone the project and use the following command to set up the server:
```shell
go run cmd/main.go
```

Result should be something like this:
```shell
2021/12/18 11:09:19 Server started ...
```

## Endpoints
- **/api/books**
  - returns the list of the books
  - method = GET

- **/api/books/{id}**
  - returns information of a single book
  - method = GET

- **/api/books**
  - creates a new book and returns the created book
  - method = POST

- **/api/books/{id}**
  - updates a book by its id 
  - method = PUT

- **/api/books/{id}**
  - deletes a book  
  - method = DELETE

## Features
- use mux package to create routes
- rest-api implementation

## Tools
- go (version 1.17)
- mux (version 8) 
