package infra

import (
	pb "wordservice/proto"
)

// DatabaseHandler is interface for function of storing data
type DatabaseHandler interface {
	GetWords() (*pb.WordCounts, error)
	UpdateWord(*pb.Word) error
}
