package api

import (
	"reflect"
	"testing"

	"github.com/arriqaaq/zizou"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

func TestNewService(t *testing.T) {
	type args struct {
		storage *gorm.DB
		logger  log.Logger
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.storage, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateBook(t *testing.T) {
	type fields struct {
		logger  log.Logger
		storage *gorm.DB
		cache   *zizou.Cache
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				logger:  tt.fields.logger,
				storage: tt.fields.storage,
				cache:   tt.fields.cache,
			}
			if err := s.CreateBook(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetBook(t *testing.T) {
	type fields struct {
		logger  log.Logger
		storage *gorm.DB
		cache   *zizou.Cache
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				logger:  tt.fields.logger,
				storage: tt.fields.storage,
				cache:   tt.fields.cache,
			}
			got, err := s.GetBook(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetAllBooks(t *testing.T) {
	type fields struct {
		logger  log.Logger
		storage *gorm.DB
		cache   *zizou.Cache
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				logger:  tt.fields.logger,
				storage: tt.fields.storage,
				cache:   tt.fields.cache,
			}
			got, err := s.GetAllBooks()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetAllBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetAllBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_SearchContacts(t *testing.T) {
	type fields struct {
		logger  log.Logger
		storage *gorm.DB
		cache   *zizou.Cache
	}
	type args struct {
		name  string
		email string
		page  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Contact
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				logger:  tt.fields.logger,
				storage: tt.fields.storage,
				cache:   tt.fields.cache,
			}
			got, err := s.SearchContacts(tt.args.name, tt.args.email, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.SearchContacts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.SearchContacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateContact(t *testing.T) {
	type fields struct {
		logger  log.Logger
		storage *gorm.DB
		cache   *zizou.Cache
	}
	type args struct {
		name   string
		email  string
		bookID uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				logger:  tt.fields.logger,
				storage: tt.fields.storage,
				cache:   tt.fields.cache,
			}
			if err := s.CreateContact(tt.args.name, tt.args.email, tt.args.bookID); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetContact(t *testing.T) {
	type fields struct {
		logger  log.Logger
		storage *gorm.DB
		cache   *zizou.Cache
	}
	type args struct {
		bookID    uint
		contactID uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Contact
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				logger:  tt.fields.logger,
				storage: tt.fields.storage,
				cache:   tt.fields.cache,
			}
			got, err := s.GetContact(tt.args.bookID, tt.args.contactID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetContact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateContact(t *testing.T) {
	type fields struct {
		logger  log.Logger
		storage *gorm.DB
		cache   *zizou.Cache
	}
	type args struct {
		name      string
		email     string
		bookID    uint
		contactID uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				logger:  tt.fields.logger,
				storage: tt.fields.storage,
				cache:   tt.fields.cache,
			}
			if err := s.UpdateContact(tt.args.name, tt.args.email, tt.args.bookID, tt.args.contactID); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteContact(t *testing.T) {
	type fields struct {
		logger  log.Logger
		storage *gorm.DB
		cache   *zizou.Cache
	}
	type args struct {
		bookID    uint
		contactID uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				logger:  tt.fields.logger,
				storage: tt.fields.storage,
				cache:   tt.fields.cache,
			}
			if err := s.DeleteContact(tt.args.bookID, tt.args.contactID); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
