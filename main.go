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
	fmt.Println(storage.CreateStorageAccount(ctx))
}