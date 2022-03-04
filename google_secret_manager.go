package gokko

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	gax "github.com/googleapis/gax-go/v2"
	log "github.com/sirupsen/logrus"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type GoogleSecretManagerClientInterface interface {
	Close() error

	AccessSecretVersion(
		context.Context,
		*secretmanagerpb.AccessSecretVersionRequest,
		...gax.CallOption,
	) (
		*secretmanagerpb.AccessSecretVersionResponse,
		error,
	)
}

type SecretGetterInterface interface {
	GetKey() string
	GetClientAndContext() (GoogleSecretManagerClientInterface, context.Context)
}

type SecretGetter struct {
	Key string
}

func (sg *SecretGetter) GetKey() string {
	return sg.Key
}

func (sg *SecretGetter) GetClientAndContext() (GoogleSecretManagerClientInterface, context.Context) {
	return getSecretManagerClientAndContext()
}

func SecretManager_NewSecretGetter(secretKey string) SecretGetterInterface {
	return &SecretGetter{
		Key: secretKey,
	}
}

func SecretManager_GetSecret(
	secretGetter SecretGetterInterface,
) string {
	client, ctx := secretGetter.GetClientAndContext()
	defer client.Close()

	projectId := Project_GetProjectId()
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf(
			"projects/%s/secrets/%s/versions/%s",
			projectId,
			secretGetter.GetKey(),
			"latest",
		),
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return ""
	}

	return string(result.Payload.Data)
}

func getSecretManagerClientAndContext() (
	GoogleSecretManagerClientInterface,
	context.Context,
) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Warn(err.Error())
		return nil, nil
	}

	return client, ctx
}
