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

func CreateStorageAccount(ctx context.Context) (storage.Account, error) {
	storageAccountsClient := getStorageAccountsClient()

	var s storage.Account

	result, err := storageAccountsClient.CheckNameAvailability(context.Background(),
		storage.AccountCheckNameAvailabilityParameters{
			Name: to.StringPtr("clmgodevstgacc1"),
			Type: to.StringPtr("Microsoft.Storage/storageAccounts")})
	if err != nil {
		log.Fatalf("%s: %v", "storage account creation failed", err)
	}
	if *result.NameAvailable != true {
		log.Fatalf("%s: %v", "storage account name not available", err)
	}

	// create the account
	future, err := storageAccountsClient.Create(ctx, "clm-go-dev", "clmgodevstgacc1", storage.AccountCreateParameters{
		Sku: &storage.Sku{
			Name: storage.StandardLRS},
		Kind:     storage.Storage,
		Location: to.StringPtr("westus2"),
		AccountPropertiesCreateParameters: &storage.AccountPropertiesCreateParameters{},
	})

	if err != nil {
		return s, fmt.Errorf("failed to start creating storage account: %v\n", err)
	}

	err = future.WaitForCompletionRef(ctx, storageAccountsClient.Client)

	if err != nil {
		return s, fmt.Errorf("failed to finish creating storage account: %v\n", err)
	}

	return future.Result(storageAccountsClient)
}


// func getStorageAccountsClient() storage.AccountsClient {
// 	storageAccountsClient := storage.NewAccountsClient(config.SubscriptionID())
// 	auth, _ := iam.GetResourceManagementAuthorizer()
// 	storageAccountsClient.Authorizer = auth
// 	storageAccountsClient.AddToUserAgent(config.UserAgent())
// 	return storageAccountsClient
// }