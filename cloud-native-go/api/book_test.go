package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native GO", Author: "Nithin C", ISBN: "121212"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native GO","author":"Nithin C","isbn":"121212"}`, string(json), "Book JSON marshaling wrong")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native GO","author":"Nithin C","isbn":"121212"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Cloud Native GO", Author: "Nithin C", ISBN: "121212"}, book, "something wrong")
}
