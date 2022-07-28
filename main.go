package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	name string `json:name`
	age  string `json:age`
}

func main() {
	fmt.Println("starting server")
	http.HandleFunc("/", UserHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	getu := user{
		name: "Amrit",
		age:  "21",
	}
	postu := user{}
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&postu)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintln(w, postu)
		fmt.Fprintln(w, "Post method completed")
	} else if r.Method == "GET" {
		err := json.NewEncoder(w).Encode(getu)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintln(w, "Get method completed")
	}
}
