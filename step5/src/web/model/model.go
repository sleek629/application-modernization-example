package model

// Data corresponds to the column of word_tb
type Data struct {
	Word string `db:"word"`
	Num  int    `db:"num"`
}

// Request is struct for data exchanged between Web and API
type Request struct {
	Word string `json:"word"`
}
