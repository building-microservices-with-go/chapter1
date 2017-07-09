package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func BenchmarkHelloWorldHandler(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r, _ := http.Post(
			"http://localhost:8080/helloworld",
			"application/json",
			bytes.NewBuffer([]byte(`{"Name":"World"}`)),
		)

		var response helloWorldResponse
		decoder := json.NewDecoder(r.Body)

		_ = decoder.Decode(&response)
	}
}

func init() {
	go server()
}
