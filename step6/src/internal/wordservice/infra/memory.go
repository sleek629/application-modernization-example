package infra

import (
	pb "wordservice/proto"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLHandler is implementation of DatabaseHandler
type MemoryHandler struct {
	wordCounts map[string]int32
}

// NewMemoryHandler returns MemoryHandler
func NewMemoryHandler() *MemoryHandler {
	wordCounts := map[string]int32{}
	return &MemoryHandler{wordCounts}
}

func (memoryHandler *MemoryHandler) convert(input map[string]int32) (output *pb.WordCounts) {
	output = &pb.WordCounts{}
	for k, v := range input {
		output.Wc = append(output.Wc, &pb.WordCount{Word: k, Count: v})
	}
	return
}

func (memoryHandler *MemoryHandler) GetWords() (wordCounts *pb.WordCounts, err error) {
	wordCounts = memoryHandler.convert(memoryHandler.wordCounts)
	return wordCounts, nil
}

func (memoryHandler *MemoryHandler) UpdateWord(word *pb.Word) (err error) {
	_, ok := memoryHandler.wordCounts[word.GetWord()]
	if ok == true {
		memoryHandler.wordCounts[word.GetWord()]++
		return nil
	}
	memoryHandler.wordCounts[word.GetWord()] = 1
	return nil
}
