package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))

	// Route handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", pageHandler("about.html"))
	http.HandleFunc("/courses", pageHandler("courses.html"))
	http.HandleFunc("/contact", pageHandler("contact.html"))
	http.Handle("/images/", fs)

	fmt.Printf("🚀 Kailash Badu Portfolio running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "./static/home.html")
}

func pageHandler(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/"+filename)
	}
}
