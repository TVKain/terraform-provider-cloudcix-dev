// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/TVKain/cloudcix-go"
	"github.com/TVKain/cloudcix-go/auth"
	"github.com/TVKain/cloudcix-go/config"
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
	BaseURL      types.String `tfsdk:"base_url" json:"base_url,optional"`
	APIKey       types.String `tfsdk:"api_key" json:"api_key,optional"`
	SettingsFile types.String `tfsdk:"settings_file" json:"settings_file,optional"`
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
			"settings_file": schema.StringAttribute{
				Description: "Path to a settings file containing CloudCIX credentials.",
				Optional:    true,
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

	if resp.Diagnostics.HasError() {
		return
	}

	// Load settings from file if provided, otherwise from environment
	var settingsFile string
	if !data.SettingsFile.IsNull() && !data.SettingsFile.IsUnknown() {
		settingsFile = data.SettingsFile.ValueString()
	}

	settings, err := config.LoadSettings(settingsFile)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to load settings",
			"Could not load CloudCIX settings: "+err.Error(),
		)
		return
	}

	// Allow provider config to override settings
	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		settings.CLOUDCIX_API_URL = data.BaseURL.ValueString()
	} else if o, ok := os.LookupEnv("CLOUDCIX_BASE_URL"); ok {
		// Backward compatibility for CLOUDCIX_BASE_URL
		settings.CLOUDCIX_API_URL = o
	}

	if !data.APIKey.IsNull() && !data.APIKey.IsUnknown() {
		settings.CLOUDCIX_API_KEY = data.APIKey.ValueString()
	}

	opts := []option.RequestOption{}

	// Determine authentication method
	// 1. Auto Auth (Username + Password + API Key)
	if settings.CLOUDCIX_API_USERNAME != "" && settings.CLOUDCIX_API_PASSWORD != "" && settings.CLOUDCIX_API_KEY != "" {
		tokenManager := auth.NewTokenManager(settings)
		opts = append(opts, auth.WithAutoAuth(tokenManager))
	} else if settings.CLOUDCIX_API_KEY != "" {
		// 2. Static Token Auth (API Key treated as Session Token)
		// This supports the legacy behavior where api_key was the session token
		opts = append(opts, option.WithAPIKey(settings.CLOUDCIX_API_KEY))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing credentials",
			"The provider requires either a settings file/env vars with full credentials (username, password, api_key) for auto-auth, or a static session token via api_key.",
		)
		return
	}

	// Set Base URL
	// Use the Compute URL from settings which handles the subdomain logic
	opts = append(opts, option.WithBaseURL(settings.ComputeURL()))

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
