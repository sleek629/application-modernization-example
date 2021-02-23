package infra

import (
	"sort"
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
	// Sort input by value(count)
	d := DictList{}
	for k, v := range input {
		d = append(d, Dict{k, int(v)})
	}
	sort.Sort(sort.Reverse(d))

	output = &pb.WordCounts{}
	for _, v := range d {
		output.Wc = append(output.Wc, &pb.WordCount{Word: v.Key, Count: int32(v.Value)})
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

// Dict is a data structure to hold key/value pairs
type Dict struct {
	Key   string
	Value int
}

// DictList is list for dict to enable sort
type DictList []Dict

func (d DictList) Len() int           { return len(d) }
func (d DictList) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d DictList) Less(i, j int) bool { return d[i].Value < d[j].Value }
