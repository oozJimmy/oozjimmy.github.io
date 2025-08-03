package main

import (
	"net/http"
	"fmt"
)

func main(){
	port := 3000
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Printf("Starting server listening to port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d",port),nil)
}
