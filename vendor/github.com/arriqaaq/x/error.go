package x

/*
	Common error handling library
	Merged from Dgraph X library
*/

import (
	"fmt"
	"log"

	// "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// Check logs fatal if err != nil.
func FatalCheck(err error, format string, args ...interface{}) {
	if err != nil {
		log.Fatalf("%+v", Wrapf(err, format, args...))
	}
}

// Check logs if err != nil.
func Check(err error) {
	if err != nil {
		log.Printf("%+v", Wrap(err))
	}
}

// Checkf is Check with extra info.
func Checkf(err error, format string, args ...interface{}) {
	if err != nil {
		log.Printf("%+v", Wrapf(err, format, args...))
	}
}

func CheckfNoTrace(err error) {
	if err != nil {
		log.Printf(err.Error())
	}
}

// AssertTrue asserts that b is true. Otherwise, it would log fatal.
func AssertTrue(b bool) {
	if !b {
		log.Fatalf("%+v", errors.Errorf("Assert failed"))
	}
}

// AssertTruef is AssertTrue with extra info.
func AssertTruef(b bool, format string, args ...interface{}) {
	if !b {
		log.Fatalf("%+v", errors.Errorf(format, args...))
	}
}

func AssertTruefNoTrace(b bool, format string, args ...interface{}) {
	if !b {
		log.Fatalf("%+v", fmt.Errorf(format, args...))
	}
}

// Wrap wraps errors from external lib.
func Wrap(err error) error {
	return errors.Wrap(err, "")
}

// Wrapf is Wrap with extra info.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf(format+" error: %+v", append(args, err)...)
}

// Errorf creates a new error with stack trace, etc.
func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// Fatalf logs fatal.
func Fatalf(format string, args ...interface{}) {
	log.Fatalf("%+v", errors.Errorf(format, args...))
}
