package infra

import (
	"app/model"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLHandler is implementation of DatabaseHandler
type MemoryHandler struct {
	WordCounts map[string]int
}

// NewMemoryHandler returns MemoryHandler
func NewMemoryHandler() *MemoryHandler {
	wordCounts := map[string]int{}
	return &MemoryHandler{wordCounts}
}

func (memoryHandler *MemoryHandler) convert(input map[string]int) (output []*model.WordCount) {
	for k, v := range input {
		output = append(output, &model.WordCount{Word: k, Count: v})
	}
	return
}

func (memoryHandler *MemoryHandler) GetWords() (wordCounts []*model.WordCount, err error) {
	wordCounts = memoryHandler.convert(memoryHandler.WordCounts)
	return wordCounts, nil
}

func (memoryHandler *MemoryHandler) UpdateWord(word string) (err error) {
	_, ok := memoryHandler.WordCounts[word]
	if ok == true {
		memoryHandler.WordCounts[word]++
		return nil
	}
	memoryHandler.WordCounts[word] = 1
	return nil
}
