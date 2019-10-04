// run these tests by executing the following command from the command line:
// go test 
package storage

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"blobstorehelper/config"
)

var _ = Describe("Storage Account", func() {
	// const timeout = time.Second * 180

	BeforeEach(func() {
		// TODO
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		// TODO
		// Add any teardown steps that needs to be executed after each test
	})

	Context("Create and Delete Storage Accounts", func() {

		It("should create and delete storage account in azure", func() {
			config.LoadSettings()
			ctx := context.Background()
		
			storageAccountName := "clmteststorageaccount"
			resourceGroup := "clm-go-dev"

			_, err := CreateStorageAccount(ctx, storageAccountName, resourceGroup)

			Expect(err).NotTo(HaveOccurred())

			err = DeleteStorageAccount(ctx, storageAccountName, resourceGroup)

			Expect(err).NotTo(HaveOccurred())
		})
	})
})

