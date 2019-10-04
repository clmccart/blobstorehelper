package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2017-06-01/storage"
	// "github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"

	"blobstorehelper/config"
	"blobstorehelper/iam"
)


func getStorageAccountsClient() storage.AccountsClient {
	groupsClient := storage.NewAccountsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	groupsClient.Authorizer = a
	groupsClient.AddToUserAgent(config.UserAgent())
	return groupsClient
}

func CreateStorageAccount(ctx context.Context, accountName string, rg string) (storage.Account, error) {
	// your logic here
}

func DeleteStorageAccount(ctx context.Context, accountName string, rg string) error {
	// your logic here
}
