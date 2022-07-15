package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHalfOfIt(t *testing.T) {
	n, _ := halfOfIt(8)
	assert.Equal(t, int8(4), n)
}
