package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/HuiCheng/devops/backend/include"
)

var (
	r   http.Handler
	tts *httptest.Server
)

func init() {
	include.InitFlag()
	include.InitDB()

	r = include.MKRouter()
	tts = httptest.NewServer(r)
	fmt.Println("Test Init")
}
