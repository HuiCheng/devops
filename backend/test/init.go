package test

import (
	"net/http/httptest"

	"github.com/HuiCheng/devops/backend/include"
)

var (
	tts *httptest.Server
)

func init() {
	include.InitFlag()
	include.InitDB("test")

	tts = httptest.NewServer(include.MKRouter())
}
