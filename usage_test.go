package datacentred

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsage(t *testing.T) {
	r := initRecorder("fixtures/usage")
	defer r.Stop()

	usage, _ := FindUsage(2017, 7)

	assert.Equal(t, 5, len(usage.Projects), "they should be equal")
	assert.Equal(t, 744.0, usage.Projects[0].Usage.Instances[0].Usage[0].Value)
}

func TestUsageErrors(t *testing.T) {
	r1 := initRecorder("fixtures/usage_errors")
	defer r1.Stop()

	_, err := FindUsage(1990, 5)
	assert.Equal(t, "404 Not Found.", err.Error(), "they should be equal")
}
