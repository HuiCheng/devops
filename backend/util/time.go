package util

import (
	"fmt"
	"time"
)

// JSONTime Struct
type JSONTime time.Time

// MarshalJSON Func
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
