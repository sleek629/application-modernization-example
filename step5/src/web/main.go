package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"web/client"
	"web/model"
)

var httpClient client.WordsClient

// Output is the data to pass to template file "index.html"
type Output struct {
	Input      string
	WordCounts []*model.WordCount
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	var input string
	// If http method is POST, the database needs to be updated.
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		input = r.FormValue("word")
		if input != "" {
			err = httpClient.UpdateWord(input)
			if err != nil {
				log.Println(err)
				fmt.Fprintf(w, err.Error())
				return
			}
		}
	}

	wordCounts, err := httpClient.GetWords()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	// html/template already has xss countermeasure function
	tpl := template.Must(template.ParseFiles("template/index.html"))
	output := Output{
		Input:      input,
		WordCounts: wordCounts,
	}

	tpl.Execute(w, output)
}

func main() {
	apiURL := os.Getenv("API_URL")
	if apiURL != "" {
		httpClient = client.NewHttpClient(apiURL)
	} else {
		log.Fatal("API_URL is not set")
	}
	http.HandleFunc("/", mainPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
