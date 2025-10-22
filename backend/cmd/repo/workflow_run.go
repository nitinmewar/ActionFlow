package repo

import (
	"orbit/cmd/database/entities"

	"gorm.io/gorm"
)

type WorkflowRunRepository struct {
	db *gorm.DB
}

func NewWorkflowRunRepository(db *gorm.DB) *WorkflowRunRepository {
	return &WorkflowRunRepository{db: db}
}

// CreateOrUpdate creates a new workflow run or updates if it already exists
func (r *WorkflowRunRepository) CreateOrUpdate(workflowRun *entities.WorkflowRun) error {
	var existing entities.WorkflowRun
	result := r.db.Where("run_id = ?", workflowRun.RunID).First(&existing)

	if result.Error == gorm.ErrRecordNotFound {
		// Create new record
		return r.db.Create(workflowRun).Error
	} else if result.Error != nil {
		return result.Error
	}

	// Update existing record
	return r.db.Model(&existing).Updates(workflowRun).Error
}

// GetByRunID retrieves a workflow run by GitHub run ID
func (r *WorkflowRunRepository) GetByRunID(runID int64) (*entities.WorkflowRun, error) {
	var workflowRun entities.WorkflowRun
	result := r.db.Where("run_id = ?", runID).First(&workflowRun)
	if result.Error != nil {
		return nil, result.Error
	}
	return &workflowRun, nil
}

// GetByID retrieves a workflow run by internal ID
func (r *WorkflowRunRepository) GetByID(id uint) (*entities.WorkflowRun, error) {
	var workflowRun entities.WorkflowRun
	result := r.db.First(&workflowRun, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &workflowRun, nil
}

// ListOptions defines filtering and sorting options for listing workflow runs
type ListOptions struct {
	Repository string // Filter by repository (owner/repo)
	Status     string // Filter by status
	Branch     string // Filter by branch
	Event      string // Filter by event type
	Conclusion string // Filter by conclusion

	// Sorting - only these two fields are supported
	SortBy    string // "run_started_at", "status"
	SortOrder string // "asc", "desc"

	// Pagination
	Page     int
	PageSize int
}

// List retrieves workflow runs with filtering, sorting, and pagination
func (r *WorkflowRunRepository) List(options ListOptions) ([]entities.WorkflowRun, int64, error) {
	var workflowRuns []entities.WorkflowRun
	var total int64

	// Build query with filters
	query := r.db.Model(&entities.WorkflowRun{})

	if options.Repository != "" {
		query = query.Where("repository = ?", options.Repository)
	}
	if options.Status != "" {
		query = query.Where("status = ?", options.Status)
	}
	if options.Branch != "" {
		query = query.Where("head_branch = ?", options.Branch)
	}
	if options.Event != "" {
		query = query.Where("event = ?", options.Event)
	}
	if options.Conclusion != "" {
		query = query.Where("conclusion = ?", options.Conclusion)
	}

	// Count total records for pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting - only allow run_started_at and status
	sortField := "run_started_at" // default
	if options.SortBy == "status" {
		sortField = "status"
	}

	sortOrder := "DESC" // default
	if options.SortOrder == "asc" {
		sortOrder = "ASC"
	}

	query = query.Order(sortField + " " + sortOrder)

	// Apply pagination
	if options.PageSize > 0 {
		offset := (options.Page - 1) * options.PageSize
		query = query.Offset(offset).Limit(options.PageSize)
	}

	// Execute query
	if err := query.Find(&workflowRuns).Error; err != nil {
		return nil, 0, err
	}

	return workflowRuns, total, nil
}
