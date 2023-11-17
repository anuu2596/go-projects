package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// creating a REST API that allows us to CREATE, READ, UPDATE and DELETE the articles on our website.
// we’ll have to define our Article structure using structs
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"name"`
	Desc    string `json:"desc"`
	Content string `json:"material"`
}

// simple server which can handle HTTP requests.
// To do this we’ll create a new file called main.go.
// Within this main.go file we’ll want to define 3 distinct functions.
// A homePage function that will handle all requests to our root URL,
// a handleRequests function that will match the URL path hit with a defined function
// a main function which will kick off our API.
func getArticle(w http.ResponseWriter, r *http.Request) {
	//var Articles []Article

	Articles := []Article{
		{Id: "1", Title: "Hello0", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	//fmt.Println("Endpoint Hit: homePage")
	//json.NewEncoder(w).Encode(Articles)
	output, err := json.Marshal(Articles)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(output)
}

func getAnimal(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sher aaya"))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "Key: "+id)

	Articles := []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	}
	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	Articles := []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	}
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := io.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(Articles)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	Articles := []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	}
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(Articles)

}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/article", getArticle).Methods("GET")
	myRouter.HandleFunc("/animals", getAnimal)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	//gorilla mux router we can add variables to our paths &
	//choose article we want to return
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	connectPGSQL()
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
