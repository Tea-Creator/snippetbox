package postgres

import (
	"github.com/Tea-Creator/snippetbox/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// SnippetModel is used for manipulation with postgresql for working with Snippet
type SnippetModel struct {
	DB *pgxpool.Pool
}

// Insert inserts a new record into db
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get returns snippet with specified id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest returns 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
