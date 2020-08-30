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

// WordCount corresponds to the column of word_tb
type WordCount struct {
	Word  string `db:"word"`
	Count int    `db:"count"`
}

// Output is the data to pass to template file "index.html"
type Output struct {
	Input      string
	WordCounts []WordCount
}

func getWords() (wordCounts []WordCount, err error) {
	rows, err := db.Query("SELECT word, count FROM word_tb ORDER BY count DESC")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var wordCount WordCount
		err = rows.Scan(&wordCount.Word, &wordCount.Count)
		wordCounts = append(wordCounts, wordCount)
	}

	return wordCounts, nil
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
		stmtUp, err := db.Prepare("UPDATE word_tb SET count = count + 1 WHERE word = ?")
		if err != nil {
			return err
		}
		defer stmtUp.Close()

		_, err = stmtUp.Exec(word)
		if err != nil {
			return err
		}
	} else {
		stmtIn, err := db.Prepare("INSERT INTO word_tb (word, count) VALUES (?, 1)")
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

	wordCounts, err := getWords()
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
