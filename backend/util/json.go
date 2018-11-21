package util

import (
	"encoding/json"
)

// InterMap2Byte Func
func InterMap2Byte(m interface{}) (s []byte) {
	var err error
	if s, err = json.Marshal(m); err != nil {
		panic(err.Error())
	}
	return s
}
