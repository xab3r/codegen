package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

var requestBody = []byte(`
	{
		"name": "Joe Doe",
		"email": "joe@doe.org",
		"age": 30,
		"activated": true
	}
`)

//easyjson:json
type User struct {
	Name      string `json:"name" print:"true"`
	Email     string `json:"email"`
	Age       int    `json:"age" print:"true"`
	Activated bool   `json:"activated" print:"false"`
}

func main() {
	u := ParseUser(requestBody)

	user := FormatUser(u)
	fmt.Print(user)
}

func ParseUser(userJson []byte) User {
	u := User{}
	json.Unmarshal(requestBody, &u)

	return u
}

func PureParseUser(userJson []byte) User {
	u := User{}
	u.UnmarshalJSON(userJson)

	return u
}

func FormatUser(u User) string {
	v := reflect.ValueOf(u)
	var buf bytes.Buffer

	for i := 0; i < v.NumField(); i++ {
		typeF := v.Type().Field(i)
		tag := typeF.Tag

		if tag.Get("print") == "true" {
			buf.WriteString(fmt.Sprintf("%s: %v\n", tag.Get("json"), v.Field(i)))
		}
	}

	return buf.String()
}
