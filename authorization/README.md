# [Interservice authorization](https://cheatsheetseries.owasp.org/cheatsheets/Microservices_Security_Cheat_Sheet.html)

Code is unsafe:
~~~go
package main

import (
	"html/template"
	"log"
	"net/http"
)

type Server struct {
}

func (s *Server) HelloHandler(w http.ResponseWriter, req *http.Request) {
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
~~~

# Let's fix it
1. [Add interservice auth](main.go#L24)