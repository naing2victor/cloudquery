package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource

	logger zerolog.Logger
	spec   *Spec
	*filetypes.Client
	writer *writers.StreamingBatchWriter

	storageClient *azblob.Client
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "azb").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal azblob spec: %w", err)
	}

	// if spec.WriteMode != specs.WriteModeAppend {
	//	return nil, fmt.Errorf("destination only supports append mode")
	// }

	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.spec.SetDefaults()

	filetypesClient, err := filetypes.NewClient(c.spec.FileSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to create filetypes client: %w", err)
	}
	c.Client = filetypesClient
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Azure credential: %w", err)
	}
	c.storageClient, err = azblob.NewClient("https://"+c.spec.StorageAccount+".blob.core.windows.net", cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Azure storage client: %w", err)
	}

	_, err = c.storageClient.UploadStream(ctx, c.spec.Container, "cq-test-file", strings.NewReader(""), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to write test file to Azure: %w", err)
	}

	c.writer, err = writers.NewStreamingBatchWriter(c, writers.WithStreamingBatchWriterBatchSizeRows(*c.spec.BatchSize), writers.WithStreamingBatchWriterBatchSizeBytes(*c.spec.BatchSizeBytes))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}
