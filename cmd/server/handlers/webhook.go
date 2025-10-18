package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"orbit/cmd/env"
	"orbit/cmd/server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                               WEBHOOK HANDLER                              */
/* -------------------------------------------------------------------------- */
func WebhookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := verifySignature(c.GetHeader("X-Hub-Signature-256"), body); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var payload models.WorkflowRunPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		fmt.Println("================================")
		fmt.Println(payload)
		// run := payload.WorkFlowRun
		// repo := payload.Repository.FullName

		// var startedAt, completedAt *time.Time
		// if run.CreatedAt != "" {
		// 	t, _ := time.Parse(time.RFC3339, run.CreatedAt)
		// 	startedAt = &t
		// }
		// if run.UpdatedAt != "" {
		// 	t, _ := time.Parse(time.RFC3339, run.UpdatedAt)
		// 	completedAt = &t
		// }
		// var duration *int
		// if startedAt != nil && completedAt != nil {
		// 	d := int(completedAt.Sub(*startedAt).Seconds())
		// 	duration = &d
		// }

		// record := entities.Run{
		// 	GithubRunID:     run.ID,
		// 	Repository:      repo,
		// 	WorkflowName:    run.Name,
		// 	Status:          run.Status,
		// 	Conclusion:      run.Conclusion,
		// 	Ref:             run.Event,
		// 	Branch:          run.HeadBranch,
		// 	CommitSHA:       run.HeadSha,
		// 	CommitMessage:   run.HeadCommit.Message,
		// 	AuthorName:      run.HeadCommit.Author.Name,
		// 	AuthorEmail:     run.HeadCommit.Author.Email,
		// 	StartedAt:       startedAt,
		// 	CompletedAt:     completedAt,
		// 	DurationSeconds: duration,
		// 	HTMLURL:         run.HtmlURL,
		// 	RawPayload:      body,
		// }

		// var existing entities.Run
		// tx := db.Where("github_run_id = ?", run.ID).First(&existing)
		// if tx.Error == gorm.ErrRecordNotFound {
		// 	db.Create(&record)
		// } else if tx.Error == nil {
		// 	record.ID = existing.ID
		// 	db.Model(&existing).Updates(record)
		// } else {
		// 	c.AbortWithStatus(http.StatusInternalServerError)
		// 	return
		// }
		c.Status(http.StatusAccepted)
	}
}

func verifySignature(header string, body []byte) error {
	secret := env.GithubWebhookSecret.GetValue()
	if secret == "" {
		return fmt.Errorf("missing webhook secret")
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expected := "sha256=" + hex.EncodeToString(mac.Sum(nil))
	if !hmac.Equal([]byte(expected), []byte(header)) {
		return fmt.Errorf("signature mismatch")
	}
	return nil
}
