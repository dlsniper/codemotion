package bench6_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dlsniper/codemotion/code/bench/6"
)

func generateSimpleRequests(b *testing.B) ([]*httptest.ResponseRecorder, []*http.Request) {
	responses := make([]*httptest.ResponseRecorder, b.N)
	requests := make([]*http.Request, b.N)

	for i := 0; i < b.N; i++ {
		responses[i] = httptest.NewRecorder()

		req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
		if err != nil {
			b.Logf("got unexpected error: %q", err.Error())
			b.Fail()
		}
		requests[i] = req
	}

	return responses, requests
}

func BenchmarkHandler(b *testing.B) {
	mux := bench6.Mux()
	responses, requests := generateSimpleRequests(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(responses[i], requests[i])
	}
}
