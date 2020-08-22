package main

import (
	"app/infra"
	"app/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var databaseHandler infra.DatabaseHandler

// Output is the data to pass to template file "index.html"
type Output struct {
	Input string
	Data  []*model.Data
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
			err = databaseHandler.UpdateWord(input)
			if err != nil {
				log.Println(err)
				fmt.Fprintf(w, err.Error())
				return
			}
		}
	}

	data, err := databaseHandler.GetWords()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	// html/template already has xss countermeasure function
	tpl := template.Must(template.ParseFiles("template/index.html"))
	output := Output{
		Input: input,
		Data:  data,
	}

	tpl.Execute(w, output)
}

func main() {
	// If you want to use MySQL DB, set MYSQL_CONNECTION in your shell.
	conn := os.Getenv("MYSQL_CONNECTION")
	if conn != "" {
		databaseHandler = infra.NewSQLHandler(conn)
	} else {
		databaseHandler = infra.NewMemoryHandler()
	}

	http.HandleFunc("/", mainPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
