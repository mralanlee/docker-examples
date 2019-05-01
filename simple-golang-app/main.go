package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 3000

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><img src='https://i.kym-cdn.com/entries/icons/facebook/000/027/475/Screen_Shot_2018-10-25_at_11.02.15_AM.jpg'></html>")
}

func main() {

	http.HandleFunc("/", handler)
	log.Println("Listening on:", port)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
