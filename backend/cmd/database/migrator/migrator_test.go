package migrator

import (
	"orbit/cmd/env"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMigrator(t *testing.T) {
	env.Load()
	db := &gorm.DB{}

	migrate := AutoMigrate(db)
	assert.NotEmpty(t, migrate)
}
