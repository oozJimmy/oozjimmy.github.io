package main

import (
	"net/http"
	"fmt"
)

func main(){
	port := 3000

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/img", http.FileServer(http.Dir("./static/img")))

	fmt.Printf("Starting server listening to port %d\n", port)

	//http.ListenAndServe(fmt.Sprintf(":%d",port),nil)
	err := http.ListenAndServe(":3000",nil)
	fmt.Println("Error:",err)
}
