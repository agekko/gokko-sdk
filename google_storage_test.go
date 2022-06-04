package gokko_test

import (
	"testing"

	"github.com/agekko/gokko-sdk"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_Storage_NewStorage(t *testing.T) {
	t.Parallel()

	storage := gokko.Storage_NewStorage("BUCKET", "FILE", "CONTENT")

	assert.Equal(t, "BUCKET", storage.GetBucket())
	assert.Equal(t, "FILE", storage.GetFileName())
	assert.Equal(t, "CONTENT", storage.GetFileContent())
}

func Test_Storage_Write(t *testing.T) {
	godotenv.Load()

	storage := gokko.Storage_NewStorage("BUCKET", "FILE", "CONTENT")

	err := gokko.Storage_Write(storage)

	assert.Nil(t, err)
}
