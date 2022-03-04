package gokko_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/agekko/gokko-sdk"
	gokkomock "github.com/agekko/gokko-sdk/mocks"
)

func Test_SecretManager_NewSecretGetter_GetKey(t *testing.T) {
	t.Parallel()

	secretGetter := gokko.SecretManager_NewSecretGetter("MYSECRET")

	assert.Equal(t, "MYSECRET", secretGetter.GetKey())
}

func TestSecretManager_GetSecret(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := gokkomock.NewMockSecretGetterInterface(ctrl)

	m.EXPECT().GetKey().Return("MYSECRET")

	secretExpected := "SECRET_EXPECTED"
	m.EXPECT().
		GetClientAndContext().
		Return(
			gokkomock.NewSecretManagerClient(secretExpected),
			context.Background(),
		)

	secretReceived := gokko.SecretManager_GetSecret(m)

	assert.Equal(t, "SECRET_EXPECTED", secretReceived)
}
