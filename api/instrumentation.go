package api

import (
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingService struct {
	requestCount        metrics.Counter
	successRequestCount metrics.Counter
	requestLatency      metrics.Histogram
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(
	counter metrics.Counter,
	successCounter metrics.Counter,
	latency metrics.Histogram,
	s Service) Service {
	return &instrumentingService{
		requestCount:        counter,
		successRequestCount: successCounter,
		requestLatency:      latency,
		Service:             s,
	}
}

func (s *instrumentingService) CreateBook(name string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", EVENT_ADD_BOOK, "error", "false"}
		s.requestCount.With(lvs...).Add(1)
		s.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		if err == nil {
			lvs := []string{"method", EVENT_ADD_BOOK, "error", "true"}
			s.successRequestCount.With(lvs...).Add(1)
		}
	}(time.Now())

	return s.Service.CreateBook(name)
}

func (s *instrumentingService) GetBook(id uint) (resp Book, err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", EVENT_GET_BOOK, "error", "false"}
		s.requestCount.With(lvs...).Add(1)
		s.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		if err == nil {
			lvs := []string{"method", EVENT_GET_BOOK, "error", "true"}
			s.successRequestCount.With(lvs...).Add(1)
		}
	}(time.Now())

	return s.Service.GetBook(id)
}

func (s *instrumentingService) GetAllBooks() (resp []Book, err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", EVENT_GET_ALL_BOOKS, "error", "false"}
		s.requestCount.With(lvs...).Add(1)
		s.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		if err == nil {
			lvs := []string{"method", EVENT_GET_ALL_BOOKS, "error", "true"}
			s.successRequestCount.With(lvs...).Add(1)
		}
	}(time.Now())

	return s.Service.GetAllBooks()
}

func (s *instrumentingService) UpdateBook(name string, id uint) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", EVENT_UPDATE_BOOK, "error", "false"}
		s.requestCount.With(lvs...).Add(1)
		s.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		if err == nil {
			lvs := []string{"method", EVENT_UPDATE_BOOK, "error", "true"}
			s.successRequestCount.With(lvs...).Add(1)
		}
	}(time.Now())

	return s.Service.UpdateBook(name, id)
}

func (s *instrumentingService) DeleteBook(id uint) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", EVENT_DELETE_BOOK, "error", "false"}
		s.requestCount.With(lvs...).Add(1)
		s.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		if err == nil {
			lvs := []string{"method", EVENT_DELETE_BOOK, "error", "true"}
			s.successRequestCount.With(lvs...).Add(1)
		}
	}(time.Now())

	return s.Service.DeleteBook(id)
}

func (s *instrumentingService) CreateContact(name string, email string, bookID uint) (err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", EVENT_ADD_CONTACT, "error", "false"}
		s.requestCount.With(lvs...).Add(1)
		s.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		if err == nil {
			lvs := []string{"method", EVENT_ADD_CONTACT, "error", "true"}
			s.successRequestCount.With(lvs...).Add(1)
		}
	}(time.Now())

	return s.Service.CreateContact(name, email, bookID)
}

func (s *instrumentingService) GetContact(bookID uint, contactID uint) (resp Contact, err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", EVENT_GET_CONTACT, "error", "false"}
		s.requestCount.With(lvs...).Add(1)
		s.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		if err == nil {
			lvs := []string{"method", EVENT_GET_CONTACT, "error", "true"}
			s.successRequestCount.With(lvs...).Add(1)
		}
	}(time.Now())

	return s.Service.GetContact(bookID, contactID)

}

func (s *instrumentingService) UpdateContact(name string, email string, bookID uint, contactID uint) (err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", EVENT_UPDATE_CONTACT, "error", "false"}
		s.requestCount.With(lvs...).Add(1)
		s.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		if err == nil {
			lvs := []string{"method", EVENT_UPDATE_CONTACT, "error", "true"}
			s.successRequestCount.With(lvs...).Add(1)
		}
	}(time.Now())

	return s.Service.UpdateContact(name, email, bookID, contactID)
}

func (s *instrumentingService) DeleteContact(bookID uint, contactID uint) (err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", EVENT_DELETE_CONTACT, "error", "false"}
		s.requestCount.With(lvs...).Add(1)
		s.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		if err == nil {
			lvs := []string{"method", EVENT_DELETE_CONTACT, "error", "true"}
			s.successRequestCount.With(lvs...).Add(1)
		}
	}(time.Now())

	return s.Service.DeleteContact(bookID, contactID)
}

func (s *instrumentingService) SearchContacts(name string, email string, page uint) (resp []Contact, count uint, err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", EVENT_SEARCH_CONTACT, "error", "false"}
		s.requestCount.With(lvs...).Add(1)
		s.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		if err == nil {
			lvs := []string{"method", EVENT_SEARCH_CONTACT, "error", "true"}
			s.successRequestCount.With(lvs...).Add(1)
		}
	}(time.Now())

	return s.Service.SearchContacts(name, email, page)

}
