package main

import (
  "log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/GBadila/multi-language-bazel-monorepo/projects/go_hello_world"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request")
    w.Write([]byte(go_hello_world.HelloWorld()))
}

func main() {
    r := mux.NewRouter()

    // IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
    r.HandleFunc("/", YourHandler)

    log.Println("Going to listen on port 5556")
    log.Fatal(http.ListenAndServe(":5556", r))
}
