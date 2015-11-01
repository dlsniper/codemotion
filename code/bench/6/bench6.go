package bench6

import (
	"html/template"
	"log"
	"net/http"
)

var visitorID uint64

func visitorCounter() uint64 {
	visitorID++
	return visitorID
}

func Hello(w http.ResponseWriter, r *http.Request) {
	response := `<!doctype html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <title>Hello world</title>
    </head>
    <body>
        <p>Hello visitor #{{.Number}}!</p>
    </body>
</html>`

	t := template.New("response")
	t, err := t.Parse(response)
	if err != nil {
		log.Printf("got error: %q", err)
		w.WriteHeader(500)
		return
	}
	err = t.Execute(w, struct{ Number uint64 }{visitorID})
	if err != nil {
		log.Printf("got error: %q", err)
	}
}

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)
	return mux
}

func Server() {
	mux := Mux()
	log.Fatalf("%q", http.ListenAndServe(":8000", mux))
}
