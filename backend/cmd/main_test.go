package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HuiCheng/devops/backend/router"
	"github.com/gavv/httpexpect"
)

func TestGetNamespace(t *testing.T) {
	tts := httptest.NewServer(router.MKRouter())
	defer tts.Close()

	e := httpexpect.New(t, tts.URL)
	e.GET("/v1/configuration/namespace").Expect().
		Status(http.StatusOK)
}

func TestPostNamespace(t *testing.T) {
	tts := httptest.NewServer(router.MKRouter())
	defer tts.Close()

	e := httpexpect.New(t, tts.URL)
	e.POST("/v1/configuration/namespace").Expect().
		Status(http.StatusOK)
}
