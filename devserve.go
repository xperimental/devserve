package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Println("usage: devserve <directory> <port>")
		return
	}

	dir := flag.Arg(0)
	port := flag.Arg(1)

	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
