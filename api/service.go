package api

import (
	// x "github.com/arriqaaq/x/convert/strings"
	"fmt"
	"github.com/arriqaaq/zizou"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	EVENT_ADD_BOOK      = "add_book"
	EVENT_GET_BOOK      = "get_book"
	EVENT_GET_ALL_BOOKS = "get_all_books"
	EVENT_UPDATE_BOOK   = "update_book"
	EVENT_DELETE_BOOK   = "delete_book"

	EVENT_ADD_CONTACT    = "add_contact"
	EVENT_GET_CONTACT    = "get_contact"
	EVENT_UPDATE_CONTACT = "update_contact"
	EVENT_DELETE_CONTACT = "delete_contact"
	EVENT_SEARCH_CONTACT = "search_contact"
)

var (
	DateFormat                       = "2006-01-02"
	DefaultInternalCacheEvictionTime = 5 * time.Minute
	DefaultInternalCacheExpiryTime   = 10 * time.Minute

	CACHE_BOOK_KEY     = "book:%d"
	CACHE_ALL_BOOK_KEY = "book:all"
	CACHE_CONTACT_KEY  = "book:%d:contact:%s"
)

type EcpmItem struct {
	Partner string
	Game    string
	Country string
	Ecpm    float64
}

type Service interface {
	CreateBook(name string) error
	GetBook(id uint) (Book, error)
	GetAllBooks() ([]Book, error)
	UpdateBook(name string, id uint) error
	DeleteBook(id uint) error

	CreateContact(name string, email string, bookID uint) error
	GetContact(bookID uint, contactID uint) (Contact, error)
	UpdateContact(name string, email string, bookID uint, contactID uint) error
	DeleteContact(bookID uint, contactID uint) error
	SearchContacts(name string, email string, page uint) ([]Contact, error)
}

func NewService(
	storage *gorm.DB,
	logger log.Logger,
) Service {

	zizouCnf := &zizou.Config{
		SweepTime: DefaultInternalCacheEvictionTime,
		ShardSize: 256,
	}
	l1cache, _ := zizou.New(zizouCnf)
	return &service{
		storage: storage,
		logger:  logger,
		cache:   l1cache,
	}
}

type service struct {
	logger  log.Logger
	storage *gorm.DB
	cache   *zizou.Cache
}

func (s *service) CreateBook(name string) error {
	return s.storage.Create(&Book{Name: name, Active: true}).Error
}

func (s *service) GetBook(id uint) (Book, error) {

	book := Book{}
	si := &Book{ID: id, Active: true}
	err := s.storage.Preload("Contacts").Where(&si).First(&book).Error
	return book, err
}

func (s *service) GetAllBooks() ([]Book, error) {

	books := make([]Book, 0, 10)
	err := s.storage.Preload("Contacts").Where(&Book{Active: true}).Find(&books).Error
	return books, err
}

func (s *service) UpdateBook(name string, id uint) error {
	book := Book{}
	err := s.storage.Where(&Book{ID: id, Active: true}).First(&book).Error
	if err != nil {
		return ERR_INVALID_ID
	}
	if name != "" {
		return s.storage.Model(&book).Updates(map[string]interface{}{"name": name}).Error
	}
	return nil
}

func (s *service) DeleteBook(id uint) error {
	book := Book{}
	err := s.storage.Where(&Book{ID: id, Active: true}).First(&book).Error
	if err != nil {
		return ERR_INVALID_ID
	}

	tx := s.storage.Begin()
	if tx.Error != nil {
		return err
	}
	err = tx.Model(&Contact{}).Where(&Contact{BookID: book.ID}).Update("active", false).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&book).Update("active", false).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *service) CreateContact(name string, email string, bookID uint) error {
	return s.storage.Create(&Contact{Name: name, Email: email, BookID: bookID, Active: true}).Error
}

func (s *service) GetContact(bookID uint, contactID uint) (Contact, error) {
	contact := Contact{}
	si := &Contact{BookID: bookID, ID: contactID, Active: true}
	err := s.storage.Where(si).First(&contact).Error
	return contact, err

}

func (s *service) UpdateContact(name string, email string, bookID uint, contactID uint) error {
	contact := Contact{}
	si := &Contact{BookID: bookID, ID: contactID, Active: true}
	err := s.storage.Where(si).First(&contact).Error
	if err != nil {
		return ERR_INVALID_ID
	}
	if name != "" && email != "" {
		return s.storage.Model(&contact).Updates(map[string]interface{}{"name": name, "email": email}).Error
	} else if name != "" && email == "" {
		return s.storage.Model(&contact).Update("name", name).Error
	} else if name == "" && email != "" {
		return s.storage.Model(&contact).Update("email", email).Error
	}

	return nil
}

func (s *service) DeleteContact(bookID uint, contactID uint) error {
	contact := Contact{}
	si := &Contact{BookID: bookID, ID: contactID, Active: true}
	err := s.storage.Where(si).First(&contact).Error
	if err != nil {
		return ERR_INVALID_ID
	}
	return s.storage.Model(&contact).Update("active", false).Error
}

func (s *service) SearchContacts(name string, email string, page uint) ([]Contact, error) {
	contacts := make([]Contact, 0, 10)
	offset := page * 10
	err := s.storage.Offset(offset).Limit(10).Where(&Contact{Name: name, Email: email, Active: true}).Find(&contacts).Error
	return contacts, err
}
