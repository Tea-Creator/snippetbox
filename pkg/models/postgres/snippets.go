package postgres

import (
	"context"
	"errors"

	"github.com/Tea-Creator/snippetbox/pkg/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// SnippetModel is used for manipulation with postgresql for working with Snippet
type SnippetModel struct {
	DB *pgxpool.Pool
}

// Insert inserts a new record into db
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `insert into snippets(title, content, created, expires)
	values ($1, $2, now(), now() + ($3 || ' days')::interval) returning id;`

	id := 0

	err := m.DB.QueryRow(context.Background(), stmt, title, content, expires).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get returns snippet with specified id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `select id, title, content, created, expires from snippets
	where expires > now() and id = $1`

	s := &models.Snippet{}

	err := m.DB.QueryRow(context.Background(), stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return s, nil
}

// Latest returns 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
