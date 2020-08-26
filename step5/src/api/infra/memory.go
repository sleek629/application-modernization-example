package infra

import (
	"api/model"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLHandler is implementation of DatabaseHandler
type MemoryHandler struct {
	data map[string]int
}

// NewMemoryHandler returns MemoryHandler
func NewMemoryHandler() *MemoryHandler {
	data := map[string]int{}
	return &MemoryHandler{data}
}

func (memoryHandler *MemoryHandler) convert(input map[string]int) (output []*model.Data) {
	for k, v := range input {
		output = append(output, &model.Data{k, v})
	}
	return
}

func (memoryHandler *MemoryHandler) GetWords() (data []*model.Data, err error) {
	data = memoryHandler.convert(memoryHandler.data)
	return data, nil
}

func (memoryHandler *MemoryHandler) UpdateWord(word string) (err error) {
	_, ok := memoryHandler.data[word]
	if ok == true {
		memoryHandler.data[word]++
		return nil
	}
	memoryHandler.data[word] = 1
	return nil
}
