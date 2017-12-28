package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path == "/test/" {
		fmt.Fprintf(w, "Привет денис!")
		return
	}
	http.NotFound(w, r)
	return
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/test/", &myHandler{})
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":"+port, nil)
}
