package main

import (
	"github.com/HuiCheng/devops/backend/include"
)

func main() {
	include.InitFlag()
	include.InitDB("prod")
	include.MKRouter().Run(include.Bind)
}
