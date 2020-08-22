package infra

import "app/model"

// DatabaseHandler is interface for function of storing data
type DatabaseHandler interface {
	GetWords() (data []*model.Data, err error)
	UpdateWord(word string) (err error)
}
