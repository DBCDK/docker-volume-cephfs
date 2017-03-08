package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
Test functions
*/
func TestIsDirectory(t *testing.T) {
	assert.True(t, IsDirectory("."))
	assert.False(t, IsDirectory("/dont/exists"))
}
