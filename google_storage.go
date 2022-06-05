package gokko

import (
	"context"

	"cloud.google.com/go/storage"
)

type StorageFile struct {
	BucketName             string
	FileName               string
	FileContent            string
	FileContentDisposition string
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

func (storageFile *StorageFile) GetFileContentDisposition() string {
	return storageFile.FileContentDisposition
}

func Storage_NewStorage(
	bucketName string,
	fileName string,
	fileContent string,
	fileContentDisposition string,
) StorageFile {
	return StorageFile{
		BucketName:             bucketName,
		FileName:               fileName,
		FileContent:            fileContent,
		FileContentDisposition: fileContentDisposition,
	}
}

func Storage_Write(storage StorageFile) error {
	client, ctx := getStorageClientAndContext()

	writer := client.Bucket(storage.GetBucketName()).
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
