package datacentred

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiscServerErrors(t *testing.T) {
	r := initRecorder("fixtures/server_error")
	defer r.Stop()

	_, err := Request("POST", "roles", nil)

	assert.Equal(t, "400 Bad Request.", err.Error(), "they should be the same")
}
