package datacentred

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsage(t *testing.T) {
	r := initRecorder("fixtures/usage")
	defer r.Stop()

	usage, _ := ShowUsage(2017, 7)

	assert.Equal(t, 5, len(usage.Projects), "they should be equal")
	assert.Equal(t, 744.0, usage.Projects[0].Usage.Instances[0].Usage[0].Value)
}
