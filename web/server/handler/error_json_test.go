package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrorJson(t *testing.T) {
	m := "example"
	r := jsonError(m)
	require.Equal(t, []byte(`{"message":"example"}`), r)
}
