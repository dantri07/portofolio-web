// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	// Sajikan file statis dari folder "public"
	fs := http.FileServer(http.Dir("./public/"))
	http.Handle("/", fs)

	// Jalankan server di port 8080
	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
