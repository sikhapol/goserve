package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

var (
	rootPath = "."
	port     = flag.String("p", "8080", "port")
)

func init() {
	log.SetPrefix("goserve: ")
}

func main() {

	flag.Parse()

	if flag.NArg() == 1 {
		rootPath = flag.Arg(0)
	} else if flag.NArg() > 1 {
		// error, only 1 root path is allowed
	}

	absRootPath, err := filepath.Abs(rootPath)
	if err != nil {
		panic(err)
	}

	log.Println("start serving", *port)
	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(absRootPath)))
}
