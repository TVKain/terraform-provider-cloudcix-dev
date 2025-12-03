// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_volume_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/storage_volume"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestStorageVolumeDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*storage_volume.StorageVolumeDataSourceModel)(nil)
	schema := storage_volume.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
