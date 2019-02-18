package api

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// AuthError represents an authorization error.
type AuthError struct {
	Realm string
}

// StatusCode is an implementation of the StatusCoder interface in go-kit/http.
func (AuthError) StatusCode() int {
	return http.StatusUnauthorized
}

// Error is an implementation of the Error interface.
func (AuthError) Error() string {
	return http.StatusText(http.StatusUnauthorized)
}

// Headers is an implementation of the Headerer interface in go-kit/http.
func (e AuthError) Headers() http.Header {
	return http.Header{
		"Content-Type":     []string{"application/json; charset=utf-8"},
		"WWW-Authenticate": []string{fmt.Sprintf(`Basic realm=%q`, e.Realm)},
	}
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic Zmxhc2g6Zmxhc2g=" returns ([]byte("flash"), []byte("flash"), true).
func parseBasicAuth(auth string) (username, password []byte, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}

	s := bytes.IndexByte(c, ':')
	if s < 0 {
		return
	}
	return c[:s], c[s+1:], true
}

// Returns a hash of a given slice.
func toHashSlice(s []byte) []byte {
	hash := sha256.Sum256(s)
	return hash[:]
}

// AuthMiddleware returns a Basic Authentication middleware for a particular user and password.
func AuthMiddleware(requiredUser, requiredPassword, realm string) endpoint.Middleware {
	requiredUserBytes := toHashSlice([]byte(requiredUser))
	requiredPasswordBytes := toHashSlice([]byte(requiredPassword))

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			auth, ok := ctx.Value(httptransport.ContextKeyRequestAuthorization).(string)
			if !ok {
				fmt.Println("here 1")
				return nil, AuthError{realm}
			}

			givenUser, givenPassword, ok := parseBasicAuth(auth)
			if !ok {
				fmt.Println("here 2", givenUser, givenPassword)
				return nil, AuthError{realm}
			}

			givenUserBytes := toHashSlice(givenUser)
			givenPasswordBytes := toHashSlice(givenPassword)

			if subtle.ConstantTimeCompare(givenUserBytes, requiredUserBytes) == 0 ||
				subtle.ConstantTimeCompare(givenPasswordBytes, requiredPasswordBytes) == 0 {
				return nil, AuthError{realm}
			}

			return next(ctx, request)
		}
	}
}
