package migrator

import (
	"orbit/cmd/database/entities"

	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(db *gorm.DB) []Migrate {
	return []Migrate{
		{
			TableName: "workflow_runs",
			Run: func(d *gorm.DB) error {
				return d.AutoMigrate(&entities.WorkflowRun{})
			},
		},
	}
}
