package main

import "testing"

var testUser = User{
	Name:      "Jane",
	Email:     "qwe@qwe.com",
	Age:       25,
	Activated: true,
}

// go test -bench=ParseUse . -benchmem
func BenchmarkParseUserEncodingJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseUser(requestBody)
	}
}

func BenchmarkParseUserEasyJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PureParseUser(requestBody)
	}
}

// go test -bench=FormatUser . -benchmem
func BenchmarkFormatUserReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FormatUser(testUser)
	}
}

func BenchmarkFormatUserCodegen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PureFormatUser(testUser)
	}
}
