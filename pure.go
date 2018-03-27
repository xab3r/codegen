// DO NOT EDIT

package main

import (
	"bytes"
	"fmt"
)

func PureFormatUser(s User) string {
	var buf = &bytes.Buffer{}

	fmt.Fprintf(buf, "name: %v\n", s.Name)
	fmt.Fprintf(buf, "age: %v\n", s.Age)

	return buf.String()
}
