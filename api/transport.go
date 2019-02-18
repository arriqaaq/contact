package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateBookRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var body createBookRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("bad request: ", err)
		return nil, ERR_INVALID_REQUEST
	}

	return body, nil
}

// r.Handle("/v1/book/{id}", getBookHandler).Methods("GET")           // Get all the books
func decodeGetBookRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var body getBookRequest

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ERR_INVALID_REQUEST
	}
	// Do validation if id is a valid int
	val, err := StringToUInt(id)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}
	body.ID = val
	return body, nil
}

func decodeUpdateBookRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var body updateBookRequest

	vars := mux.Vars(r)
	bid, ok := vars["id"]
	if !ok {
		return nil, ERR_INVALID_REQUEST
	}
	bVal, err := StringToUInt(bid)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("bad request: ", err)
		return nil, ERR_INVALID_REQUEST
	}

	body.ID = bVal
	return body, nil
}

func decodeDeleteBookRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var body deleteBookRequest

	vars := mux.Vars(r)
	bid, ok := vars["id"]
	if !ok {
		return nil, ERR_INVALID_REQUEST
	}
	// Do validation if id is a valid int
	bVal, err := StringToUInt(bid)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}

	body.ID = bVal
	return body, nil
}

// r.Handle("/v1/book/search", searchBookHandler).Methods("GET") // /search?name="isco"&email="isco@alarcon.com"

func decodeSearchContactRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var body searchContactRequest

	body.Name = r.URL.Query().Get("name")
	body.Email = r.URL.Query().Get("email")
	pageNo := r.URL.Query().Get("page")
	if pageNo == "" {
		return nil, ERR_INVALID_REQUEST
	}
	val, err := StringToUInt(pageNo)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}
	body.Page = val
	return body, nil
}

// // CRUD operations for contacts
// r.Handle("/v1/book/{book_id}/contacts", createContactHandler).Methods("POST")

func decodeCreateContactRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var body createContactRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("bad request: ", err)
		return nil, ERR_INVALID_REQUEST
	}

	vars := mux.Vars(r)
	id, ok := vars["book_id"]
	if !ok {
		return nil, ERR_INVALID_REQUEST
	}
	// Do validation if id is a valid int
	val, err := StringToUInt(id)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}
	body.BookID = val

	return body, nil
}

// r.Handle("/v1/book/{book_id}/contacts/{contact_id}", getContactHandler).Methods("GET")
func decodeGetContactRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var body getContactRequest

	vars := mux.Vars(r)
	bid, ok := vars["book_id"]
	if !ok {
		return nil, ERR_INVALID_REQUEST
	}
	cid, ok := vars["contact_id"]
	if !ok {
		return nil, ERR_INVALID_REQUEST
	}
	// Do validation if id is a valid int
	bVal, err := StringToUInt(bid)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}
	cVal, err := StringToUInt(cid)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}

	body.BookID = bVal
	body.ContactID = cVal

	return body, nil
}

// r.Handle("/v1/book/{book_id}/contacts/{contact_id}", updateContactHandler).Methods("PUT")
func decodeUpdateContactRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var body putContactRequest

	vars := mux.Vars(r)
	bid, ok := vars["book_id"]
	if !ok {
		return nil, ERR_INVALID_REQUEST
	}
	cid, ok := vars["contact_id"]
	if !ok {
		return nil, ERR_INVALID_REQUEST
	}
	// Do validation if id is a valid int
	bVal, err := StringToUInt(bid)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}
	cVal, err := StringToUInt(cid)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("bad request: ", err)
		return nil, ERR_INVALID_REQUEST
	}

	body.BookID = bVal
	body.ContactID = cVal
	return body, nil
}

// r.Handle("/v1/book/{book_id}/contacts/{contact_id}", deleteContactHandler).Methods("DELETE")
func decodeDeleteContactRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var body deleteContactRequest

	vars := mux.Vars(r)
	bid, ok := vars["book_id"]
	if !ok {
		return nil, ERR_INVALID_REQUEST
	}
	cid, ok := vars["contact_id"]
	if !ok {
		return nil, ERR_INVALID_REQUEST
	}
	// Do validation if id is a valid int
	bVal, err := StringToUInt(bid)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}
	cVal, err := StringToUInt(cid)
	if err != nil {
		return nil, ERR_INVALID_REQUEST
	}

	body.BookID = bVal
	body.ContactID = cVal
	return body, nil
}

func decodeAllBooksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ERR_INVALID_REQUEST:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// MakeHandler returns a handler for the booking service.
func MakeHandler(
	bs Service,
	logger kitlog.Logger) http.Handler {

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createBookHandler := kithttp.NewServer(
		makeCreateBookEndpoint(bs),
		decodeCreateBookRequest,
		encodeResponse,
		opts...,
	)

	getBookHandler := kithttp.NewServer(
		makeGetBookEndpoint(bs),
		decodeGetBookRequest,
		encodeResponse,
		opts...,
	)

	getAllBooksHandler := kithttp.NewServer(
		makeGetAllBookEndpoint(bs),
		decodeAllBooksRequest,
		encodeResponse,
		opts...,
	)

	updateBookHandler := kithttp.NewServer(
		makeUpdateBookEndpoint(bs),
		decodeUpdateBookRequest,
		encodeResponse,
		opts...,
	)

	deleteBookHandler := kithttp.NewServer(
		makeDeleteBookEndpoint(bs),
		decodeDeleteBookRequest,
		encodeResponse,
		opts...,
	)

	searchBookHandler := kithttp.NewServer(
		makeSearchBookEndpoint(bs),
		decodeSearchContactRequest,
		encodeResponse,
		opts...,
	)

	createContactHandler := kithttp.NewServer(
		makeCreateContactEndpoint(bs),
		decodeCreateContactRequest,
		encodeResponse,
		opts...,
	)
	getContactHandler := kithttp.NewServer(
		makeGetContactEndpoint(bs),
		decodeGetContactRequest,
		encodeResponse,
		opts...,
	)
	updateContactHandler := kithttp.NewServer(
		makeUpdateContactEndpoint(bs),
		decodeUpdateContactRequest,
		encodeResponse,
		opts...,
	)
	deleteContactHandler := kithttp.NewServer(
		makeDeleteContactEndpoint(bs),
		decodeDeleteContactRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	// CRUD operations for book
	r.Handle("/v1/book", createBookHandler).Methods("POST")
	r.Handle("/v1/book", getAllBooksHandler).Methods("GET")
	r.Handle("/v1/book/search", searchBookHandler).Methods("GET") // /search?name="isco"&email="isco@alarcon.com"
	r.Handle("/v1/book/{id}", getBookHandler).Methods("GET")
	r.Handle("/v1/book/{id}", updateBookHandler).Methods("PUT")
	r.Handle("/v1/book/{id}", deleteBookHandler).Methods("DELETE")

	// CRUD operations for contacts
	r.Handle("/v1/book/{book_id}/contact", createContactHandler).Methods("POST")
	r.Handle("/v1/book/{book_id}/contact/{contact_id}", getContactHandler).Methods("GET")
	r.Handle("/v1/book/{book_id}/contact/{contact_id}", updateContactHandler).Methods("PUT")
	r.Handle("/v1/book/{book_id}/contact/{contact_id}", deleteContactHandler).Methods("DELETE")
	r.Handle("/v1/book/search", searchBookHandler).Methods("GET") // /search?name="isco"&email="isco@alarcon.com"

	return r
}
