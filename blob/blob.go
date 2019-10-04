package blob

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/storage"
)

func IsValid() bool {
	return storage.IsValidStorageAccount("blob")
}

func CreateBlockBlobFromScratch() error {
	c := getEmptyBlob()
	fmt.Println(c)
	b := fillInBlobOptions(c)
	fmt.Println(b)
	var err error
	err = b.CreateBlockBlobFromReader(nil, &storage.PutBlobOptions{
		Timeout: 30,
	})
	return err
}


// foo does foo stuff
func Foo() string {
	return "blob"
}

func getEmptyBlob() *storage.Container {
	var accountName = "clmgodevstgacc"
	var accountKey = ""
	var blobServiceBaseURL = "core.windows.net"
	var apiVersion = "1.0"
	var useHTTPS = false
	client, _ := storage.NewClient(accountName, accountKey, blobServiceBaseURL, apiVersion, useHTTPS)
	fmt.Println("client: ", client)	
	bsc := client.GetBlobService()
	fmt.Println("bsc: ", bsc)
	c := bsc.GetContainerReference("clm-test-container")
	fmt.Println("c: ", c)

	// c.GetBlobService("blob")
	return c
}

func fillInBlobOptions(pc *storage.Container) *storage.Blob {
	blob := pc.GetBlobReference("clmtestblob")
	return blob
}