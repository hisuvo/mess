package validator

import (
	"fmt"
	"strconv"
)

// ParseUintID converts a string ID into uint.
// It rejects invalid, zero, and overflowing IDs.
func ParseUintID(value string) (uint, error) {
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid ID: %w", err)
	}

	if id == 0 {
		return 0, fmt.Errorf("ID must be greater than zero")
	}

	return uint(id), nil
}