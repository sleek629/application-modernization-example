package model

// Data corresponds to the column of word_tb
type Data struct {
	Word string `db:"word"`
	Num  int    `db:"num"`
}
