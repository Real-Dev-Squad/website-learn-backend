package resourceSchema

import "time"

type Resources struct {
	Id            string    `json:"id"  firestore:"id"`
	UserId        string    `json:"userId"  firestore:"userId"`
	Title         string    `json:"title"  firestore:"title"`
	Body          string    `json:"body"  firestore:"body"`
	Image         string    `json:"image"  firestore:"image"`
	Link          string    `json:"link" validate:"required,url" firestore:"link"`
	PublishedDate time.Time `json:"publishedDate" firestore:"publishedDate"`
	Type          string    `json:"type"  firestore:"type"`
	Filters       []string  `json:"filters" firestore:"filters"`
}
