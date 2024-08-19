package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	expected := `{"message":"Hello Json"}`
	require.JSONEq(t, expected, string(result))
}
