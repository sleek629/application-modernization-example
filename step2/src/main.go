package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Data corresponds to the column of word_tb
type Data struct {
	Word string `db:"word"`
	Num  int    `db:"num"`
}

// Output is the data to pass to template file "index.html"
type Output struct {
	Input string
	Data  []Data
}

func getWords() (words []Data, err error) {
	rows, err := db.Query("SELECT word, num FROM word_tb ORDER BY num DESC")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var wordDB Data
		err = rows.Scan(&wordDB.Word, &wordDB.Num)
		words = append(words, wordDB)
	}

	return words, nil
}

func updateWords(word string) error {
	stmtSel, err := db.Prepare("SELECT word FROM word_tb WHERE word = ?")
	if err != nil {
		return err
	}
	defer stmtSel.Close()

	rows, err := stmtSel.Query(word)
	// If rows exists, the word is already in word_tb.
	// If not, the word needs to be inserted to word_tb.
	if rows.Next() {
		stmtUp, err := db.Prepare("UPDATE word_tb SET num = num + 1 WHERE word = ?")
		if err != nil {
			return err
		}
		defer stmtUp.Close()

		_, err = stmtUp.Exec(word)
		if err != nil {
			return err
		}
	} else {
		stmtIn, err := db.Prepare("INSERT INTO word_tb (word, num) VALUES (?, 1)")
		if err != nil {
			return err
		}
		defer stmtIn.Close()

		_, err = stmtIn.Exec(word)
		if err != nil {
			return err
		}
	}
	return nil
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
			err = updateWords(input)
			if err != nil {
				log.Println(err)
				fmt.Fprintf(w, err.Error())
				return
			}
		}
	}

	data, err := getWords()
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
	db, _ = sql.Open("mysql", "user:Password@123@/word_db")
	defer db.Close()

	// checking DB connection
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", mainPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
