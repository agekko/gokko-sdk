package gokko

import (
	"context"

	"cloud.google.com/go/storage"
)

type StorageFile struct {
	BucketName  string
	FileName    string
	FileContent string
	ContentType string
}

func (storageFile *StorageFile) GetBucketName() string {
	return storageFile.BucketName
}

func (storageFile *StorageFile) GetFileName() string {
	return storageFile.FileName
}

func (storageFile *StorageFile) GetFileContent() string {
	return storageFile.FileContent
}

func (storageFile *StorageFile) GetContentType() string {
	return storageFile.ContentType
}

func Storage_NewStorageFile(
	bucketName string,
	fileName string,
	fileContent string,
	contentType string,
) StorageFile {
	return StorageFile{
		BucketName:  bucketName,
		FileName:    fileName,
		FileContent: fileContent,
		ContentType: contentType,
	}
}

func Storage_Write(storage StorageFile) error {
	client, ctx := getStorageClientAndContext()

	writer := client.Bucket(storage.GetBucketName()).
		Object(storage.GetFileName()).
		NewWriter(ctx)
	writer.ContentType = storage.GetContentType()

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
