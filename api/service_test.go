package api

import (
	"database/sql"
	"errors"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type ServiceSuite struct {
	suite.Suite
	/*
		The suite is defined as a struct, with the store and db as its
		attributes. Any variables that are to be shared between tests in a
		suite should be stored as attributes of the suite instance
	*/
	store Service
	db    *sql.DB
}

func (s *ServiceSuite) SetupSuite() {

	connString := "host=0.0.0.0" + " user=postgres" + " sslmode=disable password=docker"
	fmt.Println("postgres connection string: ", connString)
	storageDb, stErr := gorm.Open(
		"postgres", connString,
	)
	if stErr != nil {
		s.T().Fatal(stErr)
	}

	// storageDb.LogMode(true)
	storageDb.DropTableIfExists(&Contact{}, &Book{})
	storageDb.AutoMigrate(&Book{}, &Contact{})
	storageDb.Model(&Contact{}).AddForeignKey("book_id", "books(id)", "CASCADE", "CASCADE")

	var logger kitlog.Logger
	logger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)

	var cs Service
	cs = NewService(storageDb, logger)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = cs
}

func (s *ServiceSuite) SetupTest() {
	/*
		Felete all entries from the table before each test runs, to ensure a
		consistent state before our tests run.
	*/
	_, err := s.db.Exec("DELETE FROM contacts")
	if err != nil {
		s.T().Fatal(err)
	}
	_, err = s.db.Exec("DELETE FROM books")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *ServiceSuite) TearDownSuite() {
	// Close the connection after all tests in the suite finish
	s.store.Close()
	s.db.Close()
}

// This is the actual "test" as seen by Go, which runs the tests defined below
func TestServiceSuite(t *testing.T) {
	s := new(ServiceSuite)
	suite.Run(t, s)
}

func (s *ServiceSuite) TestCreateBook() {
	err := s.store.CreateBook("isco")
	if err != nil {
		s.T().Fatal(err)
	}

	// Query the database for the entry we just created
	res, err := s.db.Query(`SELECT COUNT(*) FROM books WHERE name='isco'`)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the count result
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}
	// Assert that there must be one entry with the properties of the bird that we just inserted (since the database was empty before this)
	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}

func (s *ServiceSuite) TestCreateDuplicateBook() {
	err := s.store.CreateBook("isco")
	if err != nil {
		s.T().Fatal(err)
	}

	// Query the database for the entry we just created
	res, err := s.db.Query(`SELECT COUNT(*) FROM books WHERE name='isco'`)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the count result
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}
	// Assert that there must be one entry with the properties of the bird that we just inserted (since the database was empty before this)
	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}

	err = s.store.CreateBook("isco")
	if err == nil {
		s.T().Fatal(errors.New("test failed, duplicate record created on unique column"))
	}
}

func (s *ServiceSuite) TestGetBook() {
	_, err := s.db.Exec(`INSERT INTO books(id,name,active) VALUES(10,'pascal',true)`)
	if err != nil {
		s.T().Fatal(err)
	}

	book, err := s.store.GetBook(10)
	if err != nil {
		s.T().Fatal(err)
	}

	// Assert that the details of the book is the same as the one we inserted
	expectedBook := Book{Name: "pascal"}
	if book.Name != expectedBook.Name {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedBook, book)
	}

	_, err = s.store.GetBook(200)
	if err == nil {
		s.T().Fatal(errors.New("got a record which does not exist"))
	}
}

func (s *ServiceSuite) TestUpdateBook() {
	_, err := s.db.Exec(`INSERT INTO books(id,name,active) VALUES(10,'pascal',true)`)
	if err != nil {
		s.T().Fatal(err)
	}

	err = s.store.UpdateBook("newton", 10)
	if err != nil {
		s.T().Fatal(err)
	}

	book, err := s.store.GetBook(10)
	if err != nil {
		s.T().Fatal(err)
	}

	expectedBook := Book{Name: "newton"}
	if book.Name != expectedBook.Name {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedBook, book)
	}
}

func (s *ServiceSuite) TestDeleteBook() {
	_, err := s.db.Exec(`INSERT INTO books(id,name,active) VALUES(10,'pascal',true)`)
	if err != nil {
		s.T().Fatal(err)
	}

	err = s.store.DeleteBook(10)
	if err != nil {
		s.T().Fatal(err)
	}

	_, err = s.store.GetBook(10)
	if err == nil {
		s.T().Fatal(err)
	}
}

func (s *ServiceSuite) TestDeleteBookWithContacts() {
	_, err := s.db.Exec(`INSERT INTO books(id,name,active) VALUES(1,'rooney',true)`)
	if err != nil {
		s.T().Fatal(err)
	}

	_, err = s.db.Exec(`INSERT INTO contacts(id,name,active,book_id) VALUES(1,'kai',true,1)`)
	if err != nil {
		s.T().Fatal(err)
	}

	err = s.store.DeleteBook(1)
	if err != nil {
		s.T().Fatal(err)
	}

	// Query the database for the entry we just created
	res, err := s.db.Query(`SELECT COUNT(*) FROM books WHERE name='rooney' and active=true`)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the count result
	var bcount int
	for res.Next() {
		err := res.Scan(&bcount)
		if err != nil {
			s.T().Error(err)
		}
	}
	// Assert that there must be one entry with the properties of the bird that we just inserted (since the database was empty before this)
	if bcount != 0 {
		s.T().Errorf("incorrect count, wanted 1, got %d", bcount)
	}

	res, err = s.db.Query(`SELECT COUNT(*) FROM contacts WHERE book_id=1 and active=true`)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the count result
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}
	// Assert that there must be one entry with the properties of the bird that we just inserted (since the database was empty before this)
	if count != 0 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}
