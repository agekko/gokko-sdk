package gokko_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"gokko"
	gokko_mock "gokko/mocks"
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

	m := gokko_mock.NewMockSecretGetterInterface(ctrl)

	m.EXPECT().GetKey().Return("MYSECRET")

	secretExpected := "SECRET_EXPECTED"
	m.EXPECT().
		GetClientAndContext().
		Return(
			gokko_mock.NewSecretManagerClient(secretExpected),
			context.Background(),
		)

	secretReceived := gokko.SecretManager_GetSecret(m)

	assert.Equal(t, fmt.Sprintf("data:\"%s\"", secretExpected), secretReceived)
}
