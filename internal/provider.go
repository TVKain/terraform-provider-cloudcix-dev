// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/TVKain/cloudcix-go"
	"github.com/TVKain/cloudcix-go/option"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/compute_backup"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/compute_gpu"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/compute_image"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/compute_instance"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/compute_snapshot"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/network_firewall"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/network_ip_group"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/network_router"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/network_vpn"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/project"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/storage_volume"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.ProviderWithConfigValidators = (*CloudcixDevProvider)(nil)

// CloudcixDevProvider defines the provider implementation.
type CloudcixDevProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CloudcixDevProviderModel describes the provider data model.
type CloudcixDevProviderModel struct {
	BaseURL types.String `tfsdk:"base_url" json:"base_url,optional"`
	APIKey  types.String `tfsdk:"api_key" json:"api_key,optional"`
}

func (p *CloudcixDevProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudcix-dev"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *CloudcixDevProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *CloudcixDevProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data CloudcixDevProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("CLOUDCIX_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.APIKey.IsNull() && !data.APIKey.IsUnknown() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	} else if o, ok := os.LookupEnv("CLOUDCIX_API_KEY"); ok {
		opts = append(opts, option.WithAPIKey(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing api_key value",
			"The api_key field is required. Set it in provider configuration or via the \"CLOUDCIX_API_KEY\" environment variable.",
		)
		return
	}

	client := cloudcix.NewClient(
		opts...,
	)

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *CloudcixDevProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *CloudcixDevProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		compute_backup.NewResource,
		compute_gpu.NewResource,
		compute_instance.NewResource,
		compute_snapshot.NewResource,
		network_firewall.NewResource,
		network_ip_group.NewResource,
		network_router.NewResource,
		network_vpn.NewResource,
		project.NewResource,
		storage_volume.NewResource,
	}
}

func (p *CloudcixDevProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		compute_backup.NewComputeBackupDataSource,
		compute_gpu.NewComputeGPUDataSource,
		compute_image.NewComputeImageDataSource,
		compute_instance.NewComputeInstanceDataSource,
		compute_snapshot.NewComputeSnapshotDataSource,
		network_firewall.NewNetworkFirewallDataSource,
		network_ip_group.NewNetworkIPGroupDataSource,
		network_router.NewNetworkRouterDataSource,
		network_vpn.NewNetworkVpnDataSource,
		project.NewProjectDataSource,
		storage_volume.NewStorageVolumeDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudcixDevProvider{
			version: version,
		}
	}
}
