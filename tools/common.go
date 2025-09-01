package tools

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type ClientRetriever interface {
	Get() (*azidentity.AzureCLICredential, error)
}

type ServiceClientRetriever struct {
}

func (retriever ServiceClientRetriever) Get() (*azidentity.AzureCLICredential, error) {
	client, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		return nil, fmt.Errorf("error creating az cli client: %v", err)
	}

	return client, nil

}
