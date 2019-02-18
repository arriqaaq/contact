package api

import (
	// "go.uber.org/zap"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) CreateBook(name string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"service", EVENT_ADD_BOOK,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.CreateBook(name)
}

func (s *loggingService) GetBook(id uint) (resp Book, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"service", EVENT_GET_BOOK,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.GetBook(id)

}

func (s *loggingService) GetAllBooks() (resp []Book, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"service", EVENT_GET_ALL_BOOKS,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.GetAllBooks()
}

func (s *loggingService) UpdateBook(name string, id uint) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"service", EVENT_UPDATE_BOOK,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.UpdateBook(name, id)
}

func (s *loggingService) DeleteBook(id uint) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"service", EVENT_DELETE_BOOK,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.DeleteBook(id)
}

func (s *loggingService) CreateContact(name string, email string, bookID uint) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"service", EVENT_ADD_CONTACT,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.CreateContact(name, email, bookID)
}

func (s *loggingService) GetContact(bookID uint, contactID uint) (resp Contact, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"service", EVENT_GET_CONTACT,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.GetContact(bookID, contactID)

}

func (s *loggingService) UpdateContact(name string, email string, bookID uint, contactID uint) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"service", EVENT_UPDATE_CONTACT,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.UpdateContact(name, email, bookID, contactID)
}

func (s *loggingService) DeleteContact(bookID uint, contactID uint) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"service", EVENT_DELETE_CONTACT,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.DeleteContact(bookID, contactID)
}

func (s *loggingService) SearchContacts(name string, email string, page uint) (resp []Contact, count uint, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"service", EVENT_SEARCH_CONTACT,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.SearchContacts(name, email, page)
}
