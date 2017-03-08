package main

import (
	"testing"
	"github.com/stretchr/testify/assert"

)


/**
Test functions
 */
func TestIsDirectory( t *testing.T ) {
	assert.True(t, IsDirectory("."))
	assert.False(t, IsDirectory("/dont/exists"))
}
