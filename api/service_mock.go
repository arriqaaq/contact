package api

import (
	"github.com/stretchr/testify/mock"
)

func NewMockService() *mockService {
	return new(mockService)
}

type mockService struct {
	mock.Mock
}

func (m *mockService) CreateBook(name string) error {
	rets := m.Called(name)
	return rets.Error(0)
}

func (m *mockService) GetBook(id uint) (Book, error) {
	rets := m.Called(id)
	return rets.Get(0).(Book), rets.Error(1)
}

func (m *mockService) GetAllBooks() ([]Book, error) {
	rets := m.Called()
	return rets.Get(0).([]Book), rets.Error(1)
}

func (m *mockService) UpdateBook(name string, id uint) error {
	rets := m.Called(name, id)
	return rets.Error(0)
}

func (m *mockService) DeleteBook(id uint) error {
	rets := m.Called(id)
	return rets.Error(0)
}

func (m *mockService) CreateContact(name string, email string, bookID uint) error {
	rets := m.Called(name, email, bookID)
	return rets.Error(0)
}

func (m *mockService) GetContact(bookID uint, contactID uint) (Contact, error) {
	rets := m.Called(bookID, contactID)
	return rets.Get(0).(Contact), rets.Error(1)
}

func (m *mockService) UpdateContact(name string, email string, bookID uint, contactID uint) error {
	rets := m.Called(name, email, bookID, contactID)
	return rets.Error(0)
}

func (m *mockService) DeleteContact(bookID uint, contactID uint) error {
	rets := m.Called(bookID, contactID)
	return rets.Error(0)
}

func (m *mockService) SearchContacts(name string, email string, page uint) ([]Contact, uint, error) {
	rets := m.Called(name, email, page)
	return rets.Get(0).([]Contact), uint(len(rets.Get(0).([]Contact))), rets.Error(1)
}

func (m *mockService) Close() error {
	return nil
}
