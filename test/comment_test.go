//go:build e2e
// +build e2e

package test

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestComments(t *testing.T) {
	client := resty.New()
	res, err := client.R().Get(BASE_URL + "/api/comment")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusOK, res.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()
	res, err := client.R().
		SetBody(`{"slug":"/", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode())
}
