package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test/", TestHandler)
	http.HandleFunc("/signup/", SignUpHandler)
	http.HandleFunc("/deleteaccount/", DeleteAccountHandler)
	http.HandleFunc("/send/", SendHandler)
	http.HandleFunc("/receive/", ReceiveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
