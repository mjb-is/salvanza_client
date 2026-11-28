package main

import (
	"fmt"
	"strconv"
	"time"
)

// uniqueHistoryID returns id, or id with "-2", "-3"... after it when an entry
// in history already has it, so no two Reports entries share an ID. An empty
// id gets one from the clock. Most IDs are the second a run started, so two
// Backup Sets scheduled for the same time used to get the same ID: Reports
// then showed both as selected and opened the same run for either (#173).
// IDs that do not clash are kept exactly as they were.
func uniqueHistoryID(id string, history []JobHistory) string {
	if id == "" {
		id = strconv.FormatInt(time.Now().Unix(), 10)
	}
	taken := make(map[string]bool, len(history))
	for _, h := range history {
		taken[h.ID] = true
	}
	if !taken[id] {
		return id
	}
	for n := 2; ; n++ {
		if c := fmt.Sprintf("%s-%d", id, n); !taken[c] {
			return c
		}
	}
}
