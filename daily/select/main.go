package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter , r *http.Request) {
		fmt.Println(w, "Hello")
	})
	go func() {
		if err := http.ListenAndServe(":8081",nil) ; err != nil {
			log.Fatal(err)
		}
	}()
	select {
	}
}