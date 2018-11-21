package test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/HuiCheng/devops/backend/api/v1/configuration"
	"github.com/HuiCheng/devops/backend/util"

	"github.com/gavv/httpexpect"
)

var (
	n configuration.Namespace
	k configuration.Key
	c configuration.Config
)

func TestGetNamespace(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	e.GET("/v1/configuration/namespace").Expect().
		Status(http.StatusOK)
}

func TestPostNamespace(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	post := configuration.Namespace{Lable: "TestPostNamespace"}
	data := e.POST("/v1/configuration/namespace").
		WithJSON(post).
		Expect().
		Status(http.StatusOK).JSON().Object().Raw()["data"]

	if err := json.Unmarshal(util.InterMap2Byte(data), &n); err != nil {
		t.Fatalf("%+v\n", err.Error())
	}

	t.Logf("\n%+v\n%+v\n%+v\n", post, data, n)
}

func TestPostKey(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	post := configuration.Key{Lable: "TestPostKey"}
	data := e.POST("/v1/configuration/key").
		WithJSON(post).
		Expect().
		Status(http.StatusOK).JSON().Object().Raw()["data"]

	if err := json.Unmarshal(util.InterMap2Byte(data), &k); err != nil {
		t.Fatalf("%+v\n", err.Error())
	}

	t.Logf("\n%+v\n%+v\n%+v\n", post, data, k)
}

func TestPostConfig(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	post := configuration.Config{
		Lable:      "lable",
		Note:       "note",
		Namespaces: []configuration.Namespace{n},
		Keys:       []configuration.Key{k},
	}
	data := e.POST("/v1/configuration/config").
		WithJSON(post).
		Expect().
		Status(http.StatusOK).JSON().Object().Raw()["data"]

	if err := json.Unmarshal(util.InterMap2Byte(data), &c); err != nil {
		t.Fatalf("%+v\n", err.Error())
	}

	t.Logf("\n%+v\n%+v\n%+v\n", post, data, c)
}
