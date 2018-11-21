package include

import "flag"

// Var
var (
	Bind     string
	gormType string
	gormHost string
)

// InitFlag Func
func InitFlag() {
	flag.StringVar(&Bind, "bind", ":8081", "For example, 127.0.0.1:8081 0.0.0.0:8081")
	flag.StringVar(&gormType, "gormType", "sqlite", "For example, sqlite | mysql")
	flag.StringVar(&gormHost, "gormHost", "./db.sqlite", "For example, ./db.sqlite | <username>:<password>@tcp(<host>:<port>)/<dbname>?charset=utf8")
	flag.Parse()
}
