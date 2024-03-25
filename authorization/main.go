package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	secret := os.Getenv("SERVICES_SECRET")
	if secret == "" {
		return
	}

	if secret != req.Header.Get("Authorization") {
		return
	}

	helloTemplate := `
	<html>
	<head></head>
	<body>
		<p>Some secret data</p>
	</body>
	</html>
	`
	tmpl, err := template.New("name").Parse(helloTemplate)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, "")
}
