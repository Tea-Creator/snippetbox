package models

import (
	"errors"
	"time"
)

// ErrNoRecord is error which returned when there is no results for some query
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is piece of text that user stores in our app
type Snippet struct {
	// ID is unique identifier of snippet
	ID int
	// Title is short-display name of snippet
	Title string
	// Content is body of snippet
	Content string
	// Created stores time when snippet was created
	Created time.Time
	// Expires stores time when it's necessary to stop showing snippet
	Expires time.Time
}
