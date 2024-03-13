package main

import (
	"html/template"
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
		<p><a href="{{.}}">Go back</a></p>
	</body>
	</html>
	`

	returnUrl := req.URL.Query().Get("return_url")

	tmpl, _ := template.New("my-template").Parse(helloTemplate)
	_ = tmpl.Execute(w, returnUrl)
}
