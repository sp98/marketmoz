package kite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOHLCDownSamplingTasks(t *testing.T) {
	tasks, err := GetOHLCDownSamplingTasks()
	assert.NoError(t, err)
	assert.True(t, len(*tasks) > 0)
}
