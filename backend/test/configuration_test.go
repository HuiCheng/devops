package test

import (
	"net/http"
	"testing"

	"github.com/HuiCheng/devops/backend/api/v1/configuration"

	"github.com/gavv/httpexpect"
)

var (
	n, k string
)

func TestGetNamespace(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	e.GET("/v1/configuration/namespace").Expect().
		Status(http.StatusOK)
}

func TestPostNamespace(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	body := e.POST("/v1/configuration/namespace").
		WithJSON(configuration.Namespace{Lable: "TestPostNamespace"}).
		Expect().
		Status(http.StatusOK).Body().Raw()
	t.Logf("%+v\n", body)
}

func TestPostKey(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	body := e.POST("/v1/configuration/key").
		WithJSON(configuration.Key{Lable: "TestPostKey"}).
		Expect().
		Status(http.StatusOK).Body().Raw()
	t.Logf("%+v\n", body)
}

func TestPostConfig(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	body := e.POST("/v1/configuration/config").
		Expect().
		Status(http.StatusOK).Body().Raw()
	t.Logf("%+v\n", body)
}
