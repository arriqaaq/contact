# contact

Simple contact book CRUD API app in Go

![contact](https://comps.canstockphoto.com/contact-book-icon-illustration_csp31179527.jpg)




## Features

* CRUD APIs for a contact book app
* Each contact should have a unique email address
* Allow searching by name and email address
* Search should support pagination and should return 10 items by default per invocation
* Unit tests and Integration tests for each functionality
* Basic authentication for the app
* Cache implementation (in memory + redis) [Please check branch cache]

## Example Usage

```bash

Run Build:
	make build

Run tests:
	make integration-test

Run service:
	make docker-up

Stop service:
	make docker-down

```

## CURL examples

```go
Create Book
curl -H 'Authorization: Basic Zmxhc2g6Zmxhc2g=' -XPOST 0.0.0.0:8080/v1/book -d '{"name":"alibaba"}'

Get Book
curl -H 'Authorization: Basic Zmxhc2g6Zmxhc2g=' -XGET 0.0.0.0:8080/v1/book/1

Get All Books
curl -H 'Authorization: Basic Zmxhc2g6Zmxhc2g=' -XGET 0.0.0.0:8080/v1/book

Update Book
curl -H 'Authorization: Basic Zmxhc2g6Zmxhc2g=' -XPUT 0.0.0.0:8080/v1/book/1 -d '{"name":"salah"}'

Delete Book
curl -H 'Authorization: Basic Zmxhc2g6Zmxhc2g=' -XDELETE 0.0.0.0:8080/v1/book/1

Search Contact
curl -H 'Authorization: Basic Zmxhc2g6Zmxhc2g=' -XGET '0.0.0.0:8080/v1/book/search?name=alibaba&page=0'

Create Contact
curl -H 'Authorization: Basic Zmxhc2g6Zmxhc2g=' -XPOST 0.0.0.0:8080/v1/book/1/contact -d '{"name":"alibaba","email":"ali@baba.com"}'

Get Contact
curl -H 'Authorization: Basic Zmxhc2g6Zmxhc2g=' -XGET 0.0.0.0:8080/v1/book/1/contact/1

Update Contact
curl -H 'Authorization: Basic Zmxhc2g6Zmxhc2g=' -XPUT 0.0.0.0:8080/v1/book/1/contact/1 -d '{"name":"salah"}'

Delete Contact
curl -H 'Authorization: Basic Zmxhc2g6Zmxhc2g=' -XDELETE 0.0.0.0:8080/v1/book/1/contact/1

```