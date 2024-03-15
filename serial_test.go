package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerial(t *testing.T) {
	TestSerial := NewSerial()

	ConfMade := false
	PortOpen := false

	assert.Equal(t, TestSerial.ConfigMade, ConfMade, "ConfigMade should be false")
	assert.Equal(t, TestSerial.PortOpen, PortOpen, "PortOpen should be false")
}
