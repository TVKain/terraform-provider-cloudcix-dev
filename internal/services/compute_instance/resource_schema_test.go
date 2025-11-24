// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_instance_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/compute_instance"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestComputeInstanceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_instance.ComputeInstanceModel)(nil)
	schema := compute_instance.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
