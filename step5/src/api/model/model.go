package model

// WordCount corresponds to the column of word_tb
type WordCount struct {
	Word  string `db:"word"`
	Count int    `db:"count"`
}

// Request is struct for data exchanged between Web and API
type Request struct {
	Word string `json:"word"`
}
