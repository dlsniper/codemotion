package bench4_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dlsniper/codemotion/code/bench/4"
)

var resp []byte

func BenchmarkServeHello(b *testing.B) {
	responses := make([]*httptest.ResponseRecorder, b.N, b.N)
	for i := 0; i < b.N; i++ {
		responses[i] = httptest.NewRecorder()
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		b.Logf("Failed to create request. Got error: %q", err)
		b.FailNow()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bench4.Hello(responses[i], req)
	}
	b.StopTimer()
	resp = responses[0].Body.Bytes()
}
