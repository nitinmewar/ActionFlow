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
	var runs entities.Run
	var jobs entities.Job
	var steps entities.Step

	runsM := Migrate{TableName: "runs",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&runs) }}
	jobsM := Migrate{TableName: "jobs",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&jobs) }}
	stepsM := Migrate{TableName: "steps",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&steps) }}

	return []Migrate{
		runsM,
		jobsM,
		stepsM,
	}
}
