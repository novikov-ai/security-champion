package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "user", Value: "guest"})
	http.SetCookie(w, &http.Cookie{Name: "password", Value: "qwerty123"})

	helloTemplate := `
	<html>
	<head></head>
	<body>
		<p>Hello, <b>%s</b></p>
	</body>
	</html>
	`
	name := req.URL.Query().Get("name")

	fmt.Fprintf(w, helloTemplate, html.EscapeString(name))
}
