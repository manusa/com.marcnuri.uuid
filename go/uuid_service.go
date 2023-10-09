package main

import (
	"github.com/google/uuid"
)

func NewUUID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
