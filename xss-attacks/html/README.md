# Code is unsafe, you can run any script!

~~~go
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

	fmt.Fprintf(w, helloTemplate, name)
}
~~~

![unprotected](../resources/input_unprotected.png)

Example attack:
~~~js
<script>alert(document.cookie)</script>
~~~~

# Let's fix it

1. [Encode the input](main.go#L32)

![encoded](../resources/input_encoded.png)