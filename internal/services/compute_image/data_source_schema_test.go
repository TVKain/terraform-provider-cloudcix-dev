// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_image_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/compute_image"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestComputeImageDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_image.ComputeImageDataSourceModel)(nil)
	schema := compute_image.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
