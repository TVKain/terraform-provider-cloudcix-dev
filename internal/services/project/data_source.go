// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

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

type ProjectDataSource struct {
	client *cloudcix.Client
}

var _ datasource.DataSourceWithConfigure = (*ProjectDataSource)(nil)

func NewProjectDataSource() datasource.DataSource {
	return &ProjectDataSource{}
}

func (d *ProjectDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (d *ProjectDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ProjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ProjectDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ProjectContentDataSourceEnvelope{*data}
	_, err := d.client.Project.Get(
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
	data.ID = data.Pk

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
