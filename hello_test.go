package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "hello"

	assert.Equal(t, got, want)
}