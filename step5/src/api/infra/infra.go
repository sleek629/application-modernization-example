package infra

import "api/model"

// DatabaseHandler is interface for function of storing data
type DatabaseHandler interface {
	GetWords() (wordCounts []*model.WordCount, err error)
	UpdateWord(word string) (err error)
}
