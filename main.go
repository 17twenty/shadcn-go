package main

import (
	"fmt"
	"net/http"
	"shadcn-go/shadcn"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
}
func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		next.ServeHTTP(w, r)
	})
}

func main() {
	// component := shadcn.Hello("World")

	// foo := templ.NewCSSMiddleware(messageHandler("hello"))

	// http.Handle("/", templ.Handler(component))

	// http.Handle("/", templ.Handler(shadcn.Demo()))

	router := mux.NewRouter().StrictSlash(true)
	staticRouter := router.PathPrefix("/static/")
	staticRouter.Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.Handle("/", templ.Handler(shadcn.Demo()))

	fmt.Println("Listening on :9005")
	// http.ListenAndServe(":9005", foo)
	http.ListenAndServe(":9005", router)
}
