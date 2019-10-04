package main

import (
	"fmt"
	"context"
	// "blobstorehelper/storage"
	"blobstorehelper/config"
)

func main() {
	config.LoadSettings()
	ctx := context.Background()

	// your logic here. for example:
	// storage.Create
	// stroage.Delete

	// also, don't forget to set your env variables!
}