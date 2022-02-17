package main

import (
	"example.com/handler"
	"example.com/wiki"
	"log"
	"net/http"
)

func main() {
	wiki.InsertDummyData()
	//fmt.Println("Hello")
	//
	//p2, _ := wiki.LoadPage("TestPage")
	//fmt.Println(string(p2.Body))

	handler.Route()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
