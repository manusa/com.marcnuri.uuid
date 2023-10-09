package main

import (
	"strings"
	"testing"
)

func TestNewUUID_HasValidProperties(t *testing.T) {
	t.Run("UUID has length of 36", func(t *testing.T) {
		length := len(NewUUID())
		if length != 36 {
			t.Error("Expected length of 36, got ", length)
		}
	})
	t.Run("UUID has 4 dashes", func(t *testing.T) {
		dashCount := strings.Count(NewUUID(), "-")
		if dashCount != 4 {
			t.Error("Expected 4 dashes, got ", dashCount)
		}
	})
}
