package infra

import (
	"app/model"
	"database/sql"
	"log"

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

func (mySQLHandler *MySQLHandler) GetWords() (data []*model.Data, err error) {
	rows, err := mySQLHandler.db.Query("SELECT word, num FROM word_tb ORDER BY num DESC")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var wordDB model.Data
		err = rows.Scan(&wordDB.Word, &wordDB.Num)
		data = append(data, &wordDB)
	}

	return data, nil
}

func (mySQLHandler *MySQLHandler) UpdateWord(word string) (err error) {
	stmtSel, err := mySQLHandler.db.Prepare("SELECT word FROM word_tb WHERE word = ?")
	if err != nil {
		return err
	}
	defer stmtSel.Close()

	rows, err := stmtSel.Query(word)
	// If rows exists, the word is already in word_tb.
	// If not, the word needs to be inserted to word_tb.
	if rows.Next() {
		stmtUp, err := mySQLHandler.db.Prepare("UPDATE word_tb SET num = num + 1 WHERE word = ?")
		if err != nil {
			return err
		}
		defer stmtUp.Close()

		_, err = stmtUp.Exec(word)
		if err != nil {
			return err
		}
	} else {
		stmtIn, err := mySQLHandler.db.Prepare("INSERT INTO word_tb (word, num) VALUES (?, 1)")
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
