package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"orbit/cmd/database"
	"orbit/cmd/database/entities"
	"orbit/cmd/env"
	"orbit/cmd/repo"
	"orbit/cmd/server/models"
	"time"

	"github.com/gin-gonic/gin"
)

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

func WebhookHandler(ctx *gin.Context) {
	db, _ := database.Connection()

	runRepo := repo.NewWorkflowRunRepository(db)

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := verifySignature(ctx.GetHeader("X-Hub-Signature-256"), body); err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Parse URL-encoded form data
	values, err := url.ParseQuery(string(body))
	if err != nil {
		fmt.Println("Error parsing form data:", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	payloadStr := values.Get("payload")
	if payloadStr == "" {
		fmt.Println("No payload found in request")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// URL decode the payload
	decodedPayload, err := url.QueryUnescape(payloadStr)
	if err != nil {
		fmt.Println("Error URL decoding payload:", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var webhookPayload models.GitHubWebhookPayload
	if err := json.Unmarshal([]byte(decodedPayload), &webhookPayload); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Only process workflow_run events with "completed" action
	if webhookPayload.WorkflowRun != nil {
		workflowRun := mapWebhookToWorkflowRun(webhookPayload)

		// Store in database
		if err := runRepo.CreateOrUpdate(workflowRun); err != nil {
			fmt.Println("Error storing workflow run:", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		fmt.Printf("Successfully stored workflow run %d with conclusion: %s\n",
			workflowRun.RunID, workflowRun.Conclusion)
	}

	ctx.JSON(http.StatusOK, "entry created")
}

// Helper function to map webhook data to your struct
func mapWebhookToWorkflowRun(payload models.GitHubWebhookPayload) *entities.WorkflowRun {
	wr := payload.WorkflowRun

	// Calculate duration if completed
	var duration *int64
	var completedAt *time.Time

	if payload.Action == "completed" && wr.UpdatedAt.After(wr.RunStartedAt) {
		dur := int64(wr.UpdatedAt.Sub(wr.RunStartedAt).Seconds())
		duration = &dur
		completedAt = &wr.UpdatedAt
	}

	// Extract short SHA (first 7 characters)
	headSHAShort := ""
	if len(wr.HeadSHA) >= 7 {
		headSHAShort = wr.HeadSHA[:7]
	}

	return &entities.WorkflowRun{
		RunID:                wr.ID,
		WorkflowName:         wr.Name,
		WorkflowID:           wr.WorkflowID,
		Repository:           payload.Repository.FullName,
		HeadBranch:           wr.HeadBranch,
		HeadSHA:              wr.HeadSHA,
		HeadSHAShort:         headSHAShort,
		DisplayTitle:         wr.DisplayTitle,
		WorkflowPath:         wr.Path,
		CheckSuiteID:         wr.CheckSuiteID,
		RunNumber:            wr.RunNumber,
		RunAttempt:           wr.RunAttempt,
		Event:                wr.Event,
		Status:               wr.Status,
		Conclusion:           wr.Conclusion,
		ActorLogin:           wr.Actor.Login,
		TriggeringActorLogin: wr.TriggeringActor.Login,
		GitHubURL:            wr.HTMLURL,
		RunStartedAt:         wr.RunStartedAt,
		CreatedAt:            wr.CreatedAt,
		UpdatedAt:            wr.UpdatedAt,
		CompletedAt:          completedAt,
		Duration:             duration,
		CommitMessage:        wr.HeadCommit.Message,
		CommitTimestamp:      wr.HeadCommit.Timestamp,
		CommitAuthorName:     wr.HeadCommit.Author.Name,
		CommitAuthorEmail:    wr.HeadCommit.Author.Email,
	}
}
