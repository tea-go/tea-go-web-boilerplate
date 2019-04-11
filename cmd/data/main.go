package main

import (
	"fmt"

	"time"

	"github.com/tea-go/tea-go-web-boilerplate/pkg/adding"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/reviewing"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/storage"
)

type Message interface{}

func main() {
	// create a mysql storage
	s, _ := storage.NewStorage()
	s.AutoMigrate()

	// create the available services
	adder := adding.NewService(s)       // adding "actor"
	reviewer := reviewing.NewService(s) // reviewing "actor"

	resultsBeer := adder.AddSampleBeers(adding.DefaultBeers)
	resultsReview := reviewer.AddSampleReviews(reviewing.DefaultReviews)

	go func() {
		for result := range resultsBeer {
			fmt.Printf("Added sample beer with result %s.\n", result.GetMeaning()) // human-friendly
		}
	}()

	go func() {
		for result := range resultsReview {
			fmt.Printf("Added sample review with result %d.\n", result) // machine-friendly
		}
	}()

	time.Sleep(2 * time.Second) // this is here just to get the output from goroutines printed

	fmt.Println("No more data to add!")
}
