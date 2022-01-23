package gokko

import (
	"context"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func SecretManager_GetSecret(secretKey string) string {
	client, ctx := getSecretManagerClientAndContext()
	defer client.Close()

	projectId := Project_GetProjectId()
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/%s", projectId, secretKey, "latest"),
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return ""
	}

	return result.Payload.String()
}

func getSecretManagerClientAndContext() (*secretmanager.Client, context.Context) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}
