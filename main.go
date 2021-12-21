package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
type Article struct {
	Title string `json:Title`
	Desc string `json:"desc"`
	Content string `json:"content"`

}
type Articles []Article

func testPostArticles(w http.ResponseWriter, r *http.Request){ 
	fmt.Fprintf(w, "Test POST Endpoint functional")
}
func allArticles(w http.ResponseWriter, r *http.Request){ 
	articles := Articles { 
		Article{Title:"Test Title", Desc: "Test Description", Content: "Hello World!" },
	}
	fmt.Println("End points Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homePage (w http.ResponseWriter, t *http.Request){
	fmt.Fprint(w,"HomePage Endpoint Hit")
}
func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", allArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081",myRouter))

}

func main(){ 
	handleRequests()
}
