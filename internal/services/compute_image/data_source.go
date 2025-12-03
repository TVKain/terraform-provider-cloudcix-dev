// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_image

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/TVKain/cloudcix-go"
	"github.com/TVKain/cloudcix-go/option"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type ComputeImageDataSource struct {
	client *cloudcix.Client
}

var _ datasource.DataSourceWithConfigure = (*ComputeImageDataSource)(nil)

func NewComputeImageDataSource() datasource.DataSource {
	return &ComputeImageDataSource{}
}

func (d *ComputeImageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compute_image"
}

func (d *ComputeImageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudcix.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudcix.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *ComputeImageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ComputeImageDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ComputeImageContentDataSourceEnvelope{*data}
	_, err := d.client.Compute.Images.Get(
		ctx,
		data.Pk.ValueInt64(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Content

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
