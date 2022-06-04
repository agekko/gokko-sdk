package gokko

import (
	"context"

	"cloud.google.com/go/storage"
)

type Storage struct {
	Key         string
	FileName    string
	FileContent string
}

func (sw *Storage) GetBucket() string {
	return sw.Key
}

func (sw *Storage) GetFileName() string {
	return sw.FileName
}

func (sw *Storage) GetFileContent() string {
	return sw.FileContent
}

func Storage_NewStorage(
	key string, fileName string, fileContent string,
) Storage {
	return Storage{
		Key:         key,
		FileName:    fileName,
		FileContent: fileContent,
	}
}

func Storage_Write(storage Storage) error {
	client, ctx := getStorageClientAndContext()

	writer := client.Bucket(storage.GetBucket()).
		Object(storage.GetFileName()).
		NewWriter(ctx)
	writer.ContentType = "text/plain"

	_, err := writer.Write([]byte(storage.GetFileContent()))
	if err != nil {
		return err
	}
	writer.Close()

	return nil
}

func getStorageClientAndContext() (*storage.Client, context.Context) {
	ctx := context.Background()
	client, _ := storage.NewClient(ctx)

	return client, ctx
}
