package main

import (
	"bytes"
	"fmt"
)

func PureFormatUser(s User) string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("name: %v\n", s.Name))
	buf.WriteString(fmt.Sprintf("age: %v\n", s.Age))

	return buf.String()
}
