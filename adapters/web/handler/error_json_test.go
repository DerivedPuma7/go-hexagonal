package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "hello json"

	result := jsonError(msg)

	require.Equal(t, string([]byte(`{"message":"hello json"}`)), string(result))
}
