package api

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/go-kit/kit/endpoint"
)

type createBookRequest struct {
	Name string `json:"name" valid:"required"`
}

func makeCreateBookEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(createBookRequest)

		//validation
		result, err := govalidator.ValidateStruct(&req)
		if err != nil {
			fmt.Println("token:validation:error:", result, err)
			return nil, err
		}

		err = svc.CreateBook(req.Name)
		return nil, err
	}
}

type getBookRequest struct {
	ID uint
}

type getBookResponse struct {
	Result Book `json:"result"`
}

func makeGetBookEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(getBookRequest)
		resp := getBookResponse{}
		result, err := svc.GetBook(req.ID)
		if err != nil {
			return nil, err
		}
		resp.Result = result
		return resp, nil
	}
}

type updateBookRequest struct {
	Name string `json:"name"`
	ID   uint
}

func makeUpdateBookEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(updateBookRequest)

		//validation
		result, err := govalidator.ValidateStruct(&req)
		if err != nil {
			fmt.Println("token:validation:error:", result, err)
			return nil, err
		}

		err = svc.UpdateBook(req.Name, req.ID)
		return nil, err
	}
}

type deleteBookRequest struct {
	Name string `json:"name"`
	ID   uint
}

func makeDeleteBookEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(deleteBookRequest)

		//validation
		result, err := govalidator.ValidateStruct(&req)
		if err != nil {
			fmt.Println("token:validation:error:", result, err)
			return nil, err
		}

		err = svc.DeleteBook(req.ID)
		return nil, err
	}
}

type searchContactRequest struct {
	Name  string
	Email string
	Page  uint
}

type searchContactResponse struct {
	Results []Contact `json:"results"`
	Count   uint      `json:"total_count"`
	Next    string    `json:"next,omitempty"`
	Prev    string    `json:"prev,omitempty"`
}

func makeQuery(baseUrl string, name string, email string, page uint) string {
	url := fmt.Sprintf("/v1/book/search?page=%d", page)
	if !isEmptyStr(name) {
		url += "&name=" + name
	}
	if !isEmptyStr(email) {
		url += "&email=" + email
	}
	return url

}

func makeSearchBookEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(searchContactRequest)
		resp := searchContactResponse{}
		results, count, err := svc.SearchContacts(req.Name, req.Email, req.Page)
		if err != nil {
			return nil, err
		}
		resp.Results = results
		resp.Count = count
		offset := (req.Page + 1) * MAX_PAGE_LIMIT
		if offset < resp.Count {
			baseQueryURL := "/v1/book/search?page=%d"
			resp.Next = makeQuery(baseQueryURL, req.Name, req.Email, req.Page+1)
			if req.Page > 0 {
				resp.Prev = makeQuery(baseQueryURL, req.Name, req.Email, req.Page-1)
			}
		}
		return resp, nil
	}
}

type createContactRequest struct {
	Name   string `json:"name" valid:"required"`
	Email  string `json:"email" valid:"required"`
	BookID uint
}

func makeCreateContactEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(createContactRequest)

		//validation
		result, err := govalidator.ValidateStruct(&req)
		if err != nil {
			fmt.Println("token:validation:error:", result, err)
			return nil, err
		}

		err = svc.CreateContact(req.Name, req.Email, req.BookID)
		return nil, err
	}
}

type getContactRequest struct {
	ContactID uint
	BookID    uint
}

type getContactResponse struct {
	Result Contact `json:"result"`
}

func makeGetContactEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(getContactRequest)
		resp := getContactResponse{}

		//validation
		result, err := govalidator.ValidateStruct(&req)
		if err != nil {
			fmt.Println("token:validation:error:", result, err)
			return nil, err
		}

		res, err := svc.GetContact(req.BookID, req.ContactID)
		if err != nil {
			return nil, err
		}
		resp.Result = res
		return resp, nil
	}
}

type putContactRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	BookID    uint
	ContactID uint
}

func makeUpdateContactEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(putContactRequest)

		//validation
		result, err := govalidator.ValidateStruct(&req)
		if err != nil {
			fmt.Println("token:validation:error:", result, err)
			return nil, err
		}

		err = svc.UpdateContact(req.Name, req.Email, req.BookID, req.ContactID)
		return nil, err
	}
}

type deleteContactRequest struct {
	BookID    uint
	ContactID uint
}

func makeDeleteContactEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(deleteContactRequest)

		//validation
		result, err := govalidator.ValidateStruct(&req)
		if err != nil {
			fmt.Println("token:validation:error:", result, err)
			return nil, err
		}

		err = svc.DeleteContact(req.BookID, req.ContactID)
		return nil, err
	}
}

type getAllBooksResponse struct {
	Result []Book `json:"result"`
}

func makeGetAllBookEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		resp := getAllBooksResponse{}
		result, err := svc.GetAllBooks()
		if err != nil {
			return nil, err
		}
		resp.Result = result
		return resp, nil
	}
}
