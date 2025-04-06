// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCredentialsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// assert that the data source has some values for the sensitive attributes
			{
				Config: testAccCredentialsDataSourceConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectSensitiveValue(
						"data.credsfetch_credentials.test",
						tfjsonpath.New("access_key_id"),
					),
					statecheck.ExpectSensitiveValue(
						"data.credsfetch_credentials.test",
						tfjsonpath.New("secret_access_key"),
					),
					statecheck.ExpectSensitiveValue(
						"data.credsfetch_credentials.test",
						tfjsonpath.New("session_token"),
					),
				},
			},
		},
	})
}

const testAccCredentialsDataSourceConfig = `
data "credsfetch_credentials" "test" {}
`
