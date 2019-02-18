Features:
	Caching is introduced to satisfy the following requirements,
	"The code should scale out for millions of contacts per contact book"
	We will implement a cache of 1 minute for all the GET API records




CURL:

Create Book
curl -XPOST 0.0.0.0:8080/v1/book -d '{"name":"alibaba"}'

Get Book
curl -XGET 0.0.0.0:8080/v1/book/1

Get All Books
curl -XGET 0.0.0.0:8080/v1/book

Update Book
curl -XPUT 0.0.0.0:8080/v1/book/1 -d '{"name":"salah"}'

Delete Book
curl -XDELETE 0.0.0.0:8080/v1/book/1

Search Contact
curl -XGET '0.0.0.0:8080/v1/book/search?name=alibaba&page=0'

Create Contact
curl -XPOST 0.0.0.0:8080/v1/book/1/contact -d '{"name":"alibaba","email":"ali@baba.com"}'

Get Contact
curl -XGET 0.0.0.0:8080/v1/book/1/contact/1

Update Contact
curl -XPUT 0.0.0.0:8080/v1/book/1/contact/1 -d '{"name":"salah"}'

Delete Contact
curl -XDELETE 0.0.0.0:8080/v1/book/1/contact/1


