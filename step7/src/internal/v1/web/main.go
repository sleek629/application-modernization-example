package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	pb "web/proto"

	"google.golang.org/grpc"
)

const (
	defaultAddress = "localhost:50000"
)

var client pb.WordAPIClient
var hostname string

// Output is the data to pass to template file "index.html"
type Output struct {
	Input      string
	Hostname   string
	WordCounts *pb.WordCounts
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	var input string
	// If http method is POST, the database needs to be updated.
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		input = r.FormValue("word")
		if input != "" {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			_, err = client.UpdateWord(ctx, &pb.Word{Word: input})
			if err != nil {
				log.Println(err)
				return
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	wordCounts, err := client.GetWords(ctx, &pb.Empty{})
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	// html/template already has xss countermeasure function
	tpl := template.Must(template.ParseFiles("template/index.html"))
	output := Output{
		Input:      input,
		Hostname:   hostname,
		WordCounts: wordCounts,
	}

	tpl.Execute(w, output)
}

func main() {
	hostname, _ = os.Hostname()
	var address string
	apiAddress := os.Getenv("WORDAPI_ADDRESS")
	if apiAddress != "" {
		address = apiAddress
	} else {
		address = defaultAddress
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client = pb.NewWordAPIClient(conn)
	http.HandleFunc("/", mainPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
