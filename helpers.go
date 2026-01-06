// Package helpers provides utility functions for the payment-processor project.
package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// GetRandomString generates a random string of a specified length.
func GetRandomString(length int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var b []byte
	for i := 0; i < length; i++ {
		b = append(b, charSet[rand.Intn(len(charSet))])
	}
	return string(b)
}

// GetLogger returns a new logger with a specified name.
func GetLogger(name string) *logrus.Logger {
	return logrus.New().WithFields(logrus.Fields{
		"logger": name,
	})
}

// GetHTTPError returns an HTTP error response.
func GetHTTPError(w http.ResponseWriter, status int, message string) {
	http.Error(w, message, status)
}

// GetRouter returns a new router with a specified name.
func GetRouter(name string) *mux.Router {
	return mux.NewRouter().PathPrefix(fmt.Sprintf("/%s", name)).Subrouter()
}

// DecodeJSON decodes a JSON string into a struct.
func DecodeJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// GetRandomNumber generates a random number between a specified range.
func GetRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

// GetUUID generates a random UUID.
func GetUUID() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", 
		rand.Int63(), rand.Int63(), rand.Int63(), rand.Int63(), rand.Int63())
}

// GetTimeNow returns the current time.
func GetTimeNow() time.Time {
	return time.Now()
}