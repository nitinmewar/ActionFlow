package models

import "time"

type RunListItem struct {
	ID              string     `json:"id"`
	GithubRunID     int64      `json:"github_run_id"`
	Repository      string     `json:"repository"`
	WorkflowName    string     `json:"workflow_name"`
	Status          string     `json:"status"`
	Conclusion      string     `json:"conclusion"`
	Branch          string     `json:"branch"`
	CommitSHA       string     `json:"commit_sha"`
	StartedAt       *time.Time `json:"started_at"`
	DurationSeconds *int       `json:"duration_seconds"`
}

type RunDetail struct {
	RunListItem
	CommitMessage string `json:"commit_message"`
	AuthorName    string `json:"author_name"`
	AuthorEmail   string `json:"author_email"`
	HTMLURL       string `json:"html_url"`
}
