package test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestGetNamespace(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	e.GET("/v1/configuration/namespace").Expect().
		Status(http.StatusOK)
}

func TestPostNamespace(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	e.POST("/v1/configuration/namespace").Expect().
		Status(http.StatusOK)
}
