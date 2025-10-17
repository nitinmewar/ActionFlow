package entities

import "time"

type Step struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	JobID      string    `gorm:"type:uuid;index" json:"job_id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	Number     *int      `json:"number"`
	RawPayload []byte    `gorm:"type:jsonb" json:"raw_payload"`
	CreatedAt  time.Time `json:"created_at"`
}
