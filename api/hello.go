package handler

import "net/http"

func Hello(w http.ResponseWriter, _ *http.Request) {
    println("Hello, World!") // <- this you will see in your terminal
	w.Write([]byte("Hello, World!"))
}
