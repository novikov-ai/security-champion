# [Cross-Site Request Forgery Prevention (CSRF)](https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html)

Code is unsafe:
~~~go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

type Person struct {
	Name string
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	var p Person
	_ = json.NewDecoder(req.Body).Decode(&p)
	name := p.Name

	if len(name) > 1 {
		fmt.Fprintf(w, "Got your request!")
	}

}
~~~

### Let's fix it
- [Add Content-Type checking](main.go#L23)