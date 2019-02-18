Create curl calls in this README file


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


TODO:

Add redis
Add pagination HATEOAS
Add unit test cases