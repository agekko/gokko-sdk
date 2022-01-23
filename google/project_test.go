package gokko_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	gokko "agekko/gokko-sdk/google"
)

func Test_Project_GetProjectId(t *testing.T) {
	os.Setenv("GCP_PROJECT_ID", "agekko")

	assert.Equal(
		t,
		"agekko",
		gokko.Project_GetProjectId(),
	)
}
