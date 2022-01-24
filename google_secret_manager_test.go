package gokko_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"agekko/gokko-sdk"
)

func Test_SecretManager_NewSecretGetter_GetKey(t *testing.T) {
	secretGetter := gokko.SecretManager_NewSecretGetter("MYSECRET")

	assert.Equal(t, "MYSECRET", secretGetter.GetKey())
}

func TestSecretManager_NewSecretGetter_GetClientAndContext(t *testing.T) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./../gekko.json")
	defer os.Unsetenv("GOOGLE_APPLICATION_CREDENTIALS")

	secretGetter := gokko.SecretManager_NewSecretGetter("MYSECRET")
	client, ctx := secretGetter.GetClientAndContext()
	defer client.Close()

	assert.NotNil(t, client)
	assert.NotNil(t, ctx)
}

func TestSecretManager_NewSecretGetter_GetClientAndContext_ErrorIfInvalidCredentials(
	t *testing.T,
) {
	secretGetter := gokko.SecretManager_NewSecretGetter("MYSECRET")
	client, ctx := secretGetter.GetClientAndContext()

	assert.Nil(t, client)
	assert.Nil(t, ctx)
}
