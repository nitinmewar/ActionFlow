package handlers

import (
	"net/http"
	"orbit/cmd/database/entities"
	"orbit/cmd/server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListRunsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var runs []entities.Run
		if err := db.Order("started_at DESC").Limit(25).Find(&runs).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		out := make([]models.RunListItem, len(runs))
		for i, r := range runs {
			out[i] = models.RunListItem{
				ID:              r.ID,
				GithubRunID:     r.GithubRunID,
				Repository:      r.Repository,
				WorkflowName:    r.WorkflowName,
				Status:          r.Status,
				Conclusion:      r.Conclusion,
				Branch:          r.Branch,
				CommitSHA:       r.CommitSHA,
				StartedAt:       r.StartedAt,
				DurationSeconds: r.DurationSeconds,
			}
		}
		c.JSON(http.StatusOK, out)
	}
}

func GetRunDetailHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var run entities.Run
		if err := db.Where("id = ?", id).First(&run).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.AbortWithStatus(http.StatusNotFound)
			} else {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
			return
		}
		resp := models.RunDetail{
			RunListItem: models.RunListItem{
				ID:              run.ID,
				GithubRunID:     run.GithubRunID,
				Repository:      run.Repository,
				WorkflowName:    run.WorkflowName,
				Status:          run.Status,
				Conclusion:      run.Conclusion,
				Branch:          run.Branch,
				CommitSHA:       run.CommitSHA,
				StartedAt:       run.StartedAt,
				DurationSeconds: run.DurationSeconds,
			},
			CommitMessage: run.CommitMessage,
			AuthorName:    run.AuthorName,
			AuthorEmail:   run.AuthorEmail,
			HTMLURL:       run.HTMLURL,
		}
		c.JSON(http.StatusOK, resp)
	}
}
