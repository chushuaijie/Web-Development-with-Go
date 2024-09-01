package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset =utf-8")
	fmt.Fprint(w, "<h1>Welcome to my grerat sites!</h1>")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset =utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch,email me at <a href =\"mailto:jon@calhoun.io\">jon@calhoun.in</a>.")
}
func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset =utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li>Is there a free version?</li>
	</ul>
	`)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
		// w.WriteHeader((http.StatusNotFound))
		// fmt.Fprint(w, "404 not found")
	}
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		// w.WriteHeader((http.StatusNotFound))
// 		// fmt.Fprint(w, "404 not found")
// 	}
// }

func main() {
	// var router http.HandlerFunc = pathHandler
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	// http.HandleFunc("/path", pathHandler)
	fmt.Println("Start the server on 3000...")
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
