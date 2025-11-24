// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_instance_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/compute_instance"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestComputeInstanceDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_instance.ComputeInstanceDataSourceModel)(nil)
	schema := compute_instance.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
