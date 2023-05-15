// Contains the snippet schema.

package db

import (
	"github.com/lithammer/shortuuid/v4"
	"gorm.io/gorm"
)

// Snippet is the schema for the snippets table.
type Snippet struct {
	gorm.Model

	// ID is the unique identifier of the snippet.
	ID string

	// Name is the name of the snippet.
	Name string

	// Content is the content of the snippet.
	Content string
}

// BeforeCreate is called before the snippet gets created.
// Here, the ID will get assigned to a short random UUID.
func (s *Snippet) BeforeCreate(tx *gorm.DB) error {
	s.ID = shortuuid.New()
	return nil
}
