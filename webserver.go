package main

import (
	"net/http"
	"os"
)

func Server() {
	http.Handle("/",http.FileServer(http.Dir("./frontend")))

	http.HandleFunc("/doc", func( w http.ResponseWriter, r *http.Request) {
		data,err := os.ReadFile("data.txt")
		if err != nil {
			http.Error(w,"Couldn't read file.", http.StatusInternalServerError)
		}
		w.Write(data)
	})
	http.ListenAndServe(":8080",nil)

}