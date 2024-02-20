package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Person struct {
	Name string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	helloTemplate := `{{.}}`
	tmpl := template.New("hello")
	tmpl, _ = tmpl.Parse(helloTemplate)

	name := req.URL.Query().Get("name")

	w.Header().Set("Content-Type", "application/json")

	_ = tmpl.Execute(w, fmt.Sprintf(`{"Name": "%s"}`, name))
}
