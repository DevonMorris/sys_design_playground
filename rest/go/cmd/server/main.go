package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct{
  Id string `json:"Id"`
  Title string `json:"Title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

var Articles []Article

func homePage(writer http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(writer, "Welcome to the home page!")
  fmt.Println("Endpoint hit: home page")
}

func returnAllArticles(writer http.ResponseWriter, req *http.Request) {
  fmt.Println("Endpoint hit: articles")
  json.NewEncoder(writer).Encode(Articles)
}

func returnSingleArticle(writer http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  key := vars["id"]

  for _, article := range Articles {
    if article.Id == key {
      json.NewEncoder(writer).Encode(article)
    }
  }
}

func createNewArticle(writer http.ResponseWriter, req *http.Request) {
  reqBody, _ := ioutil.ReadAll(req.Body)

  var article Article
  json.Unmarshal(reqBody, &article)

  Articles = append(Articles, article)

  json.NewEncoder(writer).Encode(article)
}

func deleteArticle(writer http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  key := vars["id"]

  for index, article := range Articles {
    if article.Id == key {
      Articles = append(Articles[:index], Articles[index+1:]...)
    }
  }
}

func handleRequests() {
  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/", homePage)
  myRouter.HandleFunc("/articles", returnAllArticles)
  myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
  myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
  myRouter.HandleFunc("/article/{id}", returnSingleArticle)
  log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
  Articles = []Article{
    {Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
    {Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
  }
  handleRequests()
}
