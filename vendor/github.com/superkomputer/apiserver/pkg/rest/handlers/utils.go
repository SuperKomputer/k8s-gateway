package handlers

import (
	"sync/atomic"
)

// NewID creates a
func NewID(lastID *int32) int32 {
	return atomic.AddInt32(lastID, 1)
}
