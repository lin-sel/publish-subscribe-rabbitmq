package util

import uuid "github.com/satori/go.uuid"

// GenerateUUID Return uuid
func GenerateUUID() uuid.UUID {
	id := uuid.NewV4()
	return id
}
