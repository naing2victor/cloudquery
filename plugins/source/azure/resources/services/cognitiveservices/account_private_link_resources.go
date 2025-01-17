package cognitiveservices

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func accountPrivateLinkResources() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cognitiveservices_account_private_link_resources",
		Resolver:             fetchAccountPrivateLinkResources,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/private-link-resources/list?tabs=HTTP#privatelinkresource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_cognitiveservices_account_private_link_resources", client.Namespacemicrosoft_cognitiveservices),
		Transform:            transformers.TransformWithStruct(&armcognitiveservices.PrivateLinkResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchAccountPrivateLinkResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armcognitiveservices.Account)
	cl := meta.(*client.Client)
	svc, err := armcognitiveservices.NewPrivateLinkResourcesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	resp, err := svc.List(ctx, group, *p.Name, nil)
	if err != nil {
		return err
	}
	res <- resp.Value
	return nil
}
