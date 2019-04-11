package reviewing

import "github.com/tea-go/tea-go-web-boilerplate/pkg/storage"

// Event defines possible outcomes from the "adding actor"
type Event int

const (
	// Done means finished processing successfully
	Done Event = iota

	// Failed means processing did not finish successfully
	Failed

	// We could also have a Queued Event which would mean queued for processing
	Queued
)

// Repository provides access to the review storage.
type Repository interface {
	// AddReview saves a given review.
	AddReview(r storage.Review) (*storage.Review, error)
}

// Service provides reviewing operations.
type Service interface {
	AddSampleReviews([]storage.Review) <-chan Event
}

type service struct {
	rR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddSampleReviews adds some sample reviews to the database
func (s *service) AddSampleReviews(data []storage.Review) <-chan Event {
	results := make(chan Event)

	go func() {
		defer close(results)

		for _, b := range data {
			_, err := s.rR.AddReview(b)
			if err != nil {
				results <- Failed
			}
			results <- Done
		}
	}()

	return results
}
