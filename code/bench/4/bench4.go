package bench4

import (
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func Main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}
