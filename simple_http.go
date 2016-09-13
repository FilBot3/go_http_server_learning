package main 

import (
  "fmt"
  "net/http"
)

// handler takes two inputs from the HandleFunc method that passes
// to this function. 
// http.ResponseWriter 
// http.Request
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Fprintf is a formatted printer. 
		// r holds the information from http.Requsst. Part of the
		// data available is Path, which is a slice. 
		fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
	}
	if r.Method == "POST" {
		fmt.Fprintf(w, "Hi there, You posted that you love %s", r.URL.Path[1:])
	}
}

func main() {
	// HandleFunc listens for a particular URL Context/Route in this
	// case its root "/". Then pass the functionality to the method
	// called handler. 
	http.HandleFunc("/", handler)
	// This method starts the web server on a specified port,
	// and pass of to a router. However, we're not using a rounter
	// so we pass it nil. 
	fmt.Println("Listening and Serving on port 8080")
	http.ListenAndServe(":8080", nil)
}