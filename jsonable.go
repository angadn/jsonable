package jsonable

import (
	"encoding/json"
	"fmt"
)

// JSONable is used to embed JSON-generating and Scanning code. However, we're choosing
// NOT to embed it so as to not worry about weird entity states, were this struct to gain
// state over time.
type JSONable struct {
}

// New is a convenience constructor for JSONable.
func New() JSONable {
	return JSONable{}
}

// Scan can be used to implement https://golang.org/pkg/database/sql/#Scanner.
func (jsonable JSONable) Scan(dest interface{}, src interface{}, tag string) error {
	if b, ok := src.([]byte); ok {
		return json.Unmarshal(b, dest)
	}

	return fmt.Errorf("cannot convert %v to string for JSON-parsing %s", src, tag)
}

// JSON is useful to share JSON-generating logic.
func (jsonable JSONable) JSON(target interface{}, err *error) string {
	var (
		e   error
		ret []byte
	)

	ret, e = json.Marshal(target)
	if *err == nil {
		*err = e
	}

	return string(ret)
}
