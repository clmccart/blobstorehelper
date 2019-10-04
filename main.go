package main

import (
	"fmt"
	"context"
	"blobstorehelper/storage"
	"blobstorehelper/config"
)

func main() {
	config.LoadSettings()
	ctx := context.Background()

	accountName := "testsorageacc"
	rg := "clm-go-dev"

	_, err := storage.CreateStorageAccount(ctx, accountName, rg)
	if err != nil {
		fmt.Println("Was unable to create storage account: %v", err.Error())
	} else {
		fmt.Println("Storage account CREATED successfully")
	}
	
	err = storage.DeleteStorageAccount(ctx, accountName, rg)
	if err != nil {
		fmt.Println("Was unable to delete storage account: %v", err.Error())
	} else {
		fmt.Println("Storage account DELETED successfully")

	}
}