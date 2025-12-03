// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_gpu_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/compute_gpu"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestComputeGPUModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_gpu.ComputeGPUModel)(nil)
	schema := compute_gpu.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
