// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func TargetSslProxies() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_target_ssl_proxies",
		Resolver:  fetchTargetSslProxies,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "certificate_map",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateMap"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "proxy_header",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProxyHeader"),
			},
			{
				Name:     "service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Service"),
			},
			{
				Name:     "ssl_certificates",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SslCertificates"),
			},
			{
				Name:     "ssl_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SslPolicy"),
			},
		},
	}
}

func fetchTargetSslProxies(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListTargetSslProxiesRequest{
		Project: c.ProjectId,
	}
	it := c.Services.ComputeTargetSslProxiesClient.List(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp

	}
	return nil
}
