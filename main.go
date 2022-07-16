package main

import (
	"fmt"

	"log"
	"net/http"
	"strconv"

	"zirconlabz.com/main/mypackage"
)

var times = 0

func main2() {
	fmt.Println("Hello!")
	mypackage.PrintHello()
	http.HandleFunc("/", myHandler)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	times += 1
	fmt.Fprintf(w, "Hello0000: "+strconv.Itoa(times))
}
