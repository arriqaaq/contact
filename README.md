Create curl calls in this README file


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


TODO:

Add unit test cases

Create v2 branch with the following:
Add caching/redis
