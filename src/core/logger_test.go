package core

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	t.Run("test new logger", func(t *testing.T) {
		logger := NewLogger()

		if logger == nil {
			t.Error("logger is nil")
		}
	})
}
