package api

import (
	// x "github.com/arriqaaq/x/convert/strings"
	"fmt"
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

	MAX_PAGE_LIMIT = 10
)

var (
	DateFormat                       = "2006-01-02"
	DefaultInternalCacheEvictionTime = 5 * time.Minute
	DefaultInternalCacheExpiryTime   = 10 * time.Second

	CACHE_BOOK_KEY         = "book:%d"
	CACHE_ALL_BOOK_KEY     = "book:all"
	CACHE_CONTACT_KEY      = "book:%d:contact:%s"
	CACHE_SEARCH_KEY       = "search:name:%s:email:%s:page:%d"
	CACHE_SEARCH_COUNT_KEY = "search:name:%s:email:%s:page:%d:count"
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
	SearchContacts(name string, email string, page uint) ([]Contact, uint, error)

	Close() error
}

func NewService(
	storage *gorm.DB,
	logger log.Logger,
) Service {

	return &service{
		storage: storage,
		logger:  logger,
		cache:   NewL1Cache(),
	}
}

type service struct {
	logger  log.Logger
	storage *gorm.DB
	cache   Cache
}

func (s *service) CreateBook(name string) error {
	bi := Book{Name: name, Active: true}
	return s.storage.Create(&bi).Error
}

func (s *service) GetBook(id uint) (Book, error) {
	pKey := fmt.Sprintf(CACHE_BOOK_KEY, id)

	// l1 cache
	val, found := s.cache.Get(pKey)
	if found {
		return val.(Book), nil
	}

	book := Book{}
	si := &Book{ID: id, Active: true}
	err := s.storage.Preload("Contacts").Where(&si).First(&book).Error
	if err == nil {
		s.cache.Set(pKey, book, DefaultInternalCacheExpiryTime)
	}
	return book, err
}

func (s *service) GetAllBooks() ([]Book, error) {
	pKey := CACHE_ALL_BOOK_KEY

	// l1 cache
	val, found := s.cache.Get(pKey)
	if found {
		return val.([]Book), nil
	}

	books := make([]Book, 0, 10)
	err := s.storage.Preload("Contacts").Where(&Book{Active: true}).Find(&books).Error
	if err == nil {
		s.cache.Set(pKey, books, DefaultInternalCacheExpiryTime)
	}
	return books, err
}

func (s *service) UpdateBook(name string, id uint) error {

	book := Book{}
	err := s.storage.Where(&Book{ID: id, Active: true}).First(&book).Error
	if err != nil {
		return ERR_INVALID_ID
	}
	if name != "" {
		err := s.storage.Model(&book).Updates(map[string]interface{}{"name": name}).Error
		if err == nil {
			pKey := fmt.Sprintf(CACHE_BOOK_KEY, id)
			book.Name = name
			s.cache.Set(pKey, book, DefaultInternalCacheExpiryTime)
		}
		return err

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
	pKey := fmt.Sprintf(CACHE_BOOK_KEY, id)
	s.cache.Delete(pKey)
	return tx.Commit().Error
}

func (s *service) CreateContact(name string, email string, bookID uint) error {
	return s.storage.Create(&Contact{Name: name, Email: email, BookID: bookID, Active: true}).Error
}

func (s *service) GetContact(bookID uint, contactID uint) (Contact, error) {
	pKey := fmt.Sprintf(CACHE_CONTACT_KEY, bookID, contactID)

	// l1 cache
	val, found := s.cache.Get(pKey)
	if found {
		return val.(Contact), nil
	}

	contact := Contact{}
	si := &Contact{BookID: bookID, ID: contactID, Active: true}
	err := s.storage.Where(si).First(&contact).Error
	if err == nil {
		s.cache.Set(pKey, contact, DefaultInternalCacheExpiryTime)
	}
	return contact, err

}

func (s *service) UpdateContact(name string, email string, bookID uint, contactID uint) error {

	var err error
	pKey := fmt.Sprintf(CACHE_CONTACT_KEY, bookID, contactID)

	contact := Contact{}
	si := &Contact{BookID: bookID, ID: contactID, Active: true}
	err = s.storage.Where(si).First(&contact).Error
	if err != nil {
		return ERR_INVALID_ID
	}
	if name != "" && email != "" {
		err = s.storage.Model(&contact).Updates(map[string]interface{}{"name": name, "email": email}).Error
		if err == nil {
			contact.Name = name
			contact.Email = email
			s.cache.Set(pKey, contact, DefaultInternalCacheExpiryTime)
		}
	} else if name != "" && email == "" {
		err = s.storage.Model(&contact).Update("name", name).Error
		if err == nil {
			contact.Name = name
			s.cache.Set(pKey, contact, DefaultInternalCacheExpiryTime)
		}
	} else if name == "" && email != "" {
		err = s.storage.Model(&contact).Update("email", email).Error
		if err == nil {
			contact.Email = email
			s.cache.Set(pKey, contact, DefaultInternalCacheExpiryTime)
		}
	}

	return err
}

func (s *service) DeleteContact(bookID uint, contactID uint) error {
	contact := Contact{}
	si := &Contact{BookID: bookID, ID: contactID, Active: true}
	err := s.storage.Where(si).First(&contact).Error
	if err != nil {
		return ERR_INVALID_ID
	}
	err = s.storage.Model(&contact).Update("active", false).Error
	if err != nil {
		pKey := fmt.Sprintf(CACHE_CONTACT_KEY, bookID, contactID)
		s.cache.Delete(pKey)
	}
	return err
}

func (s *service) SearchContacts(name string, email string, page uint) ([]Contact, uint, error) {
	pKey := fmt.Sprintf(CACHE_SEARCH_KEY, name, email, page)
	pCountKey := fmt.Sprintf(CACHE_SEARCH_COUNT_KEY, name, email, page)

	// l1 cache
	val, found := s.cache.Get(pKey)
	if found {
		cnt, cntFound := s.cache.Get(pCountKey)
		if cntFound {
			return val.([]Contact), cnt.(uint), nil
		}
	}

	var count uint
	contacts := make([]Contact, 0, MAX_PAGE_LIMIT)
	offset := page * MAX_PAGE_LIMIT

	si := &Contact{Name: name, Email: email, Active: true}
	err := s.storage.Model(&Contact{}).Where(&si).Count(&count).Error
	if err != nil {
		return nil, count, err
	}
	if offset > count {
		return nil, count, ERR_INVALID_PAGE
	}
	err = s.storage.Offset(offset).Limit(MAX_PAGE_LIMIT).Where(&si).Find(&contacts).Error
	if err == nil {
		s.cache.Set(pKey, contacts, DefaultInternalCacheExpiryTime)
		s.cache.Set(pCountKey, count, DefaultInternalCacheExpiryTime)
	}
	return contacts, count, err
}

func (s *service) Close() error {
	return s.storage.Close()
}
