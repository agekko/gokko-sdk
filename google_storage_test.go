package gokko_test

import (
	"testing"

	"github.com/agekko/gokko-sdk"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_Storage_NewStorage(t *testing.T) {
	t.Parallel()

	storage := gokko.Storage_NewStorageFile("BUCKET", "FILE", "CONTENT", "CONTENT/Type")

	assert.Equal(t, "BUCKET", storage.GetBucketName())
	assert.Equal(t, "FILE", storage.GetFileName())
	assert.Equal(t, "CONTENT", storage.GetFileContent())
	assert.Equal(t, "CONTENT/Type", storage.GetContentType())
}

func Test_Storage_Write(t *testing.T) {
	godotenv.Load()

	storage := gokko.Storage_NewStorageFile("BUCKET", "FILE", "CONTENT", "CONTENT/Type")

	err := gokko.Storage_Write(storage)

	assert.Nil(t, err)
}
