package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1: solved go test, warning: no tests to run
// https://cloud.tencent.com/developer/article/1820021

//2: run test from terminal
//go to rest_errors folder run "go test", or "go test -cover"
func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("new internal server error", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "new internal server error", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])

	// if want to print out what is err struct looks like
	// errBytes, _ := json.Marshal(err)
	// fmt.Println(string(errBytes))
}

func TestNewError(t *testing.T) {
	err := NewError("new error")
	assert.NotNil(t, err)
	assert.EqualValues(t, "new error", err.Error())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("new bad request error")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "new bad request error", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
	assert.Nil(t, err.Causes)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("new not found error")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "new not found error", err.Message)
	assert.EqualValues(t, "not_found", err.Error)
	assert.Nil(t, err.Causes)
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("new unauthorized error")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status)
	assert.EqualValues(t, "new unauthorized error", err.Message)
	assert.EqualValues(t, "unauthorized_error", err.Error)
	assert.Nil(t, err.Causes)
}
