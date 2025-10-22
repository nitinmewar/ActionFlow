package commands

import (
	"orbit/cmd/env"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestMigrateTables(t *testing.T) {
	env.Load()

	cmd := &cobra.Command{}

	cmd.AddCommand(Migrate())
	err := Migrate().RunE(cmd, nil)
	assert.Empty(t, err)
}
