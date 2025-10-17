package entities

import "time"

type Job struct {
	ID              string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	RunID           string     `gorm:"type:uuid;index" json:"run_id"`
	GithubJobID     *int64     `json:"github_job_id"`
	Name            string     `json:"name"`
	Status          string     `json:"status"`
	Conclusion      string     `json:"conclusion"`
	StartedAt       *time.Time `json:"started_at"`
	CompletedAt     *time.Time `json:"completed_at"`
	DurationSeconds *int       `json:"duration_seconds"`
	RawPayload      []byte     `gorm:"type:jsonb" json:"raw_payload"`
	CreatedAt       time.Time  `json:"created_at"`
	Steps           []Step     `json:"steps" gorm:"foreignKey:JobID"`
}
