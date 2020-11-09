package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
)


type Article struct {
    Id      string    `json:"Id"`
    Title   string `json:"Title"`
    Subtitle    string `json:"sub"`
    Content string `json:"content"`
    Creationts string `json:"creationts"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "News Articles\n")
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "News Articles\n")
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func search(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "News Articles\n")
	keys := r.URL.Query()["search"]
    key :=keys[0]

    for _, article := range Articles {
       if article.Title == string(key) {
            json.NewEncoder(w).Encode(article)
		}
		if article.Subtitle == string(key) {
            json.NewEncoder(w).Encode(article)
		}
		if article.Content == string(key) {
            json.NewEncoder(w).Encode(article)
        }
    }
}


func createNewArticle(w http.ResponseWriter, r *http.Request) {  
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}


func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
   
   myRouter.HandleFunc("/home", homePage)
   myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
   myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
   myRouter.HandleFunc("/articles", returnAllArticles)
  
   myRouter.HandleFunc("/", search)
	
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    Articles = []Article{
        Article{Id: "1", Title: "Attacks", Subtitle: "Delhi", Content: "Attacks in Delhi took place", Creationts: "12/03/2020"},
        Article{Id: "2 ", Title: "Politics", Subtitle: "Mumbai", Content: "Huge political shift in mumbai", Creationts:"11/03/2019"},
    }
    handleRequests()
}
