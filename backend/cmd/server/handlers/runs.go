package handlers

import (
	"net/http"
	"orbit/cmd/database"
	"orbit/cmd/repo"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ListWorkflowRuns handles GET /api/runs
func ListWorkflowRuns(ctx *gin.Context) {
	db, _ := database.Connection()

	runRepo := repo.NewWorkflowRunRepository(db)

	// Parse query parameters
	options := parseListOptions(ctx)

	// Get workflow runs from repository
	workflowRuns, total, err := runRepo.List(options)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch workflow runs"})
		return
	}

	// Calculate pagination info
	pagination := gin.H{
		"total":     total,
		"page":      options.Page,
		"page_size": options.PageSize,
		"pages":     calculateTotalPages(total, options.PageSize),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       workflowRuns,
		"pagination": pagination,
	})
}

// parseListOptions extracts and validates query parameters
func parseListOptions(ctx *gin.Context) repo.ListOptions {
	options := repo.ListOptions{}

	// Filters
	// options.Repository = ctx.Query("repository")
	// options.Status = ctx.Query("status")
	// options.Branch = ctx.Query("branch")
	// options.Event = ctx.Query("event")
	// options.Conclusion = ctx.Query("conclusion")

	// Sorting - only allow "run_started_at" or "status"
	sortBy := ctx.Query("sort_by")
	if sortBy == "status" {
		options.SortBy = "status"
	} else {
		options.SortBy = "run_started_at" // default
	}

	// Sort order
	sortOrder := ctx.Query("sort_order")
	if sortOrder == "asc" {
		options.SortOrder = "asc"
	} else {
		options.SortOrder = "desc" // default
	}

	// Pagination
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page < 1 {
		page = 1
	}
	options.Page = page

	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	if pageSize < 1 {
		pageSize = 20 // default page size
	} else if pageSize > 100 {
		pageSize = 100 // max page size
	}
	options.PageSize = pageSize

	return options
}

func calculateTotalPages(total int64, pageSize int) int64 {
	if pageSize == 0 {
		return 0
	}
	pages := total / int64(pageSize)
	if total%int64(pageSize) > 0 {
		pages++
	}
	return pages
}

// GetWorkflowRun handles GET /api/runs/:id
func GetWorkflowRun(ctx *gin.Context) {
	db, _ := database.Connection()

	runRepo := repo.NewWorkflowRunRepository(db)

	// Get run ID from URL parameter
	idParam := ctx.Param("id")
	runID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid run ID"})
		return
	}

	// Get workflow run from repository
	workflowRun, err := runRepo.GetByRunID(runID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Workflow run not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch workflow run"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": workflowRun,
	})
}
