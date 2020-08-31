package main

import (
	"context"
	"log"
	"net"
	"os"
	"wordservice/infra"

	"google.golang.org/grpc"

	pb "wordservice/proto"
)

const (
	defaultPort = "50000"
)

type wordAPI struct {
	databaseHandler infra.DatabaseHandler
}

func main() {
	// If you want to use MySQL DB, set MYSQL_CONNECTION in your shell.
	conn := os.Getenv("MYSQL_CONNECTION")
	var databaseHandler infra.DatabaseHandler
	if conn != "" {
		databaseHandler = infra.NewSQLHandler(conn)
	} else {
		databaseHandler = infra.NewMemoryHandler()
	}

	port := ":" + defaultPort
	if value, ok := os.LookupEnv("PORT"); ok {
		port = ":" + value
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	svc := &wordAPI{databaseHandler}
	pb.RegisterWordAPIServer(srv, svc)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (w *wordAPI) GetWords(ctx context.Context, in *pb.Empty) (*pb.WordCounts, error) {
	wordCounts, err := w.databaseHandler.GetWords()
	if err != nil {
		return nil, err
	}
	return wordCounts, nil
}

func (w *wordAPI) UpdateWord(ctx context.Context, word *pb.Word) (*pb.Empty, error) {
	err := w.databaseHandler.UpdateWord(word)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
