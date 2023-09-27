package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessBatch(t *testing.T) {
	err := ProcessBatch(Config, metricsCh)
	assert.NoError(t, err)
}
