package handlers

import (
	"encoding/json"
	"math"
	"time"
)

// Helper to safely dereference string pointers (default to empty string)
func derefString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

// Helper to safely dereference optional runner fields
func derefRunnerString(p *string) *string {
	if p == nil {
		return nil
	}
	return p
}

// Helper to convert []string labels to json.RawMessage
func labelsToJSON(labels []string) json.RawMessage {
	b, _ := json.Marshal(labels)
	return b
}

// Helper to compute duration in seconds (rounded)
func computeDuration(start, end *time.Time) *int64 {
	if start == nil || end == nil {
		return nil
	}
	dur := end.Sub(*start).Seconds()
	rounded := int64(math.Round(dur))
	return &rounded
}
