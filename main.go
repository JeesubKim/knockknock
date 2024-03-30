package main
import (
		"fmt"
		"net/http"
       )

func main(){
	fmt.Println("Hello Go!")


	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	}

	mux.HandleFunc("GET /words", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "get words")
	}
	mux.HandleFunc("POST /words", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "post words")
	}
	if err := http.ListenAndServe("localhost:8080", mux) err != nil {
		fmt.Println(err.Error())
	}
}
