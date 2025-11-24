// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/project"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestProjectDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*project.ProjectDataSourceModel)(nil)
	schema := project.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
