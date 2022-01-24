package gokko

import "os"

func Project_GetProjectId() string {
	return os.Getenv("GCP_PROJECT_ID")
}
