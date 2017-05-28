package main

import "net/http"
import "log"

func main() {
	http.HandleFunc("/",someFunc)
	http.ListenAndServe(":8008",nil)
}

func someFunc (w http.ResponseWriter,req *http.Request) {
	path:= req.URL.Path[2:]
	log.Println(path)
	w.Write([]byte(path))
}
