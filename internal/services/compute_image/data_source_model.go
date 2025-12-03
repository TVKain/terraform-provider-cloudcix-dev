// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_image

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeImageContentDataSourceEnvelope struct {
	Content ComputeImageDataSourceModel `json:"content,computed"`
}

type ComputeImageDataSourceModel struct {
	ID        types.Int64  `tfsdk:"id" path:"id,required"`
	Filename  types.String `tfsdk:"filename" json:"filename,computed"`
	OsVariant types.String `tfsdk:"os_variant" json:"os_variant,computed"`
	SKUName   types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
