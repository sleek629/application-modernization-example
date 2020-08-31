package infra

import (
	"database/sql"
	"log"
	pb "wordservice/proto"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLHandler is implementation of DatabaseHandler
type MySQLHandler struct {
	db *sql.DB
}

// NewSQLHandler returns sqlHandler
func NewSQLHandler(conn string) *MySQLHandler {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &MySQLHandler{db}
}

func (mySQLHandler *MySQLHandler) GetWords() (wordCounts *pb.WordCounts, err error) {
	wordCounts = &pb.WordCounts{}
	rows, err := mySQLHandler.db.Query("SELECT word, count FROM word_tb ORDER BY count DESC")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var wordCount pb.WordCount
		err = rows.Scan(&wordCount.Word, &wordCount.Count)
		wordCounts.Wc = append(wordCounts.Wc, &wordCount)
	}

	return wordCounts, nil
}

func (mySQLHandler *MySQLHandler) UpdateWord(word *pb.Word) (err error) {
	stmtSel, err := mySQLHandler.db.Prepare("SELECT word FROM word_tb WHERE word = ?")
	if err != nil {
		return err
	}
	defer stmtSel.Close()

	rows, err := stmtSel.Query(word.GetWord())
	// If rows exists, the word is already in word_tb.
	// If not, the word needs to be inserted to word_tb.
	if rows.Next() {
		stmtUp, err := mySQLHandler.db.Prepare("UPDATE word_tb SET count = count + 1 WHERE word = ?")
		if err != nil {
			return err
		}
		defer stmtUp.Close()

		_, err = stmtUp.Exec(word.GetWord())
		if err != nil {
			return err
		}
	} else {
		stmtIn, err := mySQLHandler.db.Prepare("INSERT INTO word_tb (word, count) VALUES (?, 1)")
		if err != nil {
			return err
		}
		defer stmtIn.Close()

		_, err = stmtIn.Exec(word.GetWord())
		if err != nil {
			return err
		}
	}
	return nil
}
