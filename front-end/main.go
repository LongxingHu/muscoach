package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("tag9 正在监听80端口....")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Panic(err)
	}
}
