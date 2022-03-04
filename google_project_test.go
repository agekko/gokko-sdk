package gokko_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/agekko/gokko-sdk"
)

func Test_Project_GetProjectId(t *testing.T) {
	t.Parallel()

	os.Setenv("GCP_PROJECT_ID", "agekko")
	defer os.Unsetenv("GCP_PROJECT_ID")

	assert.Equal(
		t,
		"agekko",
		gokko.Project_GetProjectId(),
	)
}
