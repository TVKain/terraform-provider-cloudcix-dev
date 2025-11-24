// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_backup_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/compute_backup"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestComputeBackupModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_backup.ComputeBackupModel)(nil)
	schema := compute_backup.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
