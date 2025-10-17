package entities

import "time"

type Run struct {
	ID              string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	GithubRunID     int64      `gorm:"uniqueIndex;not null" json:"github_run_id"`
	Repository      string     `json:"repository"`
	WorkflowName    string     `json:"workflow_name"`
	Status          string     `json:"status"`
	Conclusion      string     `json:"conclusion"`
	Ref             string     `json:"ref"`
	Branch          string     `json:"branch"`
	CommitSHA       string     `json:"commit_sha"`
	CommitMessage   string     `json:"commit_message"`
	AuthorName      string     `json:"author_name"`
	AuthorEmail     string     `json:"author_email"`
	StartedAt       *time.Time `json:"started_at"`
	CompletedAt     *time.Time `json:"completed_at"`
	DurationSeconds *int       `json:"duration_seconds"`
	HTMLURL         string     `json:"html_url"`
	RawPayload      []byte     `gorm:"type:jsonb" json:"raw_payload"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Jobs            []Job      `json:"jobs" gorm:"foreignKey:RunID"`
}
