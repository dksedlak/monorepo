package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMiddleware(t *testing.T) {
	service := "testando"
	obj := NewMiddleware(service)

	require.NotNil(t, obj)
	require.Equal(t, service, obj.Service)
}
