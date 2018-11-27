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
	ns []configuration.Namespace
	ks []configuration.Key
	c  configuration.Config
)

func TestGetNamespace(t *testing.T) {
	e := httpexpect.New(t, tts.URL)
	e.GET("/v1/configuration/namespace").Expect().
		Status(http.StatusOK)
}

func TestPostNamespace(t *testing.T) {
	n := configuration.Namespace{}
	e := httpexpect.New(t, tts.URL)
	for _, i := range []string{"1", "2"} {
		post := configuration.Namespace{Lable: "TestPostNamespace" + i}
		data := e.POST("/v1/configuration/namespace").
			WithJSON(post).
			Expect().
			Status(http.StatusOK).JSON().Object().Raw()["data"]

		if err := json.Unmarshal(util.InterMap2Byte(data), &n); err != nil {
			t.Fatalf("%+v\n", err.Error())
		}
		ns = append(ns, n)
		t.Logf("\n%+v\n%+v\n%+v\n", post, data, n)
	}
}

func TestPostKey(t *testing.T) {
	k := configuration.Key{}
	e := httpexpect.New(t, tts.URL)
	for _, i := range []string{"1", "2"} {
		post := configuration.Key{Lable: "TestPostKey" + i}
		data := e.POST("/v1/configuration/key").
			WithJSON(post).
			Expect().
			Status(http.StatusOK).JSON().Object().Raw()["data"]

		if err := json.Unmarshal(util.InterMap2Byte(data), &k); err != nil {
			t.Fatalf("%+v\n", err.Error())
		}
		ks = append(ks, k)

		t.Logf("\n%+v\n%+v\n%+v\n", post, data, k)
	}

}

func mkVaules() (r []configuration.Value) {
	for _, n := range ns {
		for _, k := range ks {
			r = append(r, configuration.Value{
				Namespace: n,
				Key:       k,
				Value:     n.Lable + " " + k.Lable,
			})
		}
	}
	return
}

func TestPostConfig(t *testing.T) {

	e := httpexpect.New(t, tts.URL)
	post := configuration.Config{
		Lable:      "lable",
		Note:       "note",
		Namespaces: ns,
		Keys:       ks,
		Values:     mkVaules(),
	}

	data := e.POST("/v1/configuration/config").
		WithJSON(post).
		Expect().
		Status(http.StatusOK).JSON().Object().Raw()

	t.Logf("\n%+v\n%+v\n", post, data)
}

func TestGetConfig(t *testing.T) {
	e := httpexpect.New(t, tts.URL)

	data := e.GET("/v1/configuration/config").
		Expect().
		Status(http.StatusOK).JSON().Object().Raw()
	j, _ := json.Marshal(data)
	t.Logf("\n%+v\n", string(j))
}
