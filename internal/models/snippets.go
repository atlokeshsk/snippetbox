package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets(title, content, created, expires)
			 values (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY));`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id,title,content,created,expires
			 FROM snippets
			 WHERE expires > UTC_TIMESTAMP() AND id = ?`

	snippet := &Snippet{}
	err := m.DB.QueryRow(stmt, id).Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.CreatedAt, &snippet.ExpiresAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	return snippet, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id,title,content,created,expires
			FROM snippets
			WHERE expires > UTC_TIMESTAMP()
			ORDER BY id DESC
			LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*Snippet{}
	for rows.Next() {
		snippet := &Snippet{}
		err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.CreatedAt, &snippet.ExpiresAt)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, snippet)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return snippets, nil
}
