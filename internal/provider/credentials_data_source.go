// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &CredentialsDataSource{}

func NewExampleDataSource() datasource.DataSource {
	return &CredentialsDataSource{}
}

// CredentialsDataSource defines the data source implementation.
type CredentialsDataSource struct {
	client *http.Client
}

// ExampleDataSourceModel describes the data source data model.
type ExampleDataSourceModel struct {
	Profile         types.String `tfsdk:"profile"`
	AccessKeyID     types.String `tfsdk:"access_key_id"`
	SecretAccessKey types.String `tfsdk:"secret_access_key"`
	SessionToken    types.String `tfsdk:"session_token"`
}

func (d *CredentialsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_credentials"
}

func (d *CredentialsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example data source",

		Attributes: map[string]schema.Attribute{
			"profile": schema.StringAttribute{
				MarkdownDescription: "Example configurable attribute",
				Optional:            true,
			},
			"access_key_id": schema.StringAttribute{
				MarkdownDescription: "AWS access key ID",
				Computed:            true,
				Sensitive:           true,
			},
			"secret_access_key": schema.StringAttribute{
				MarkdownDescription: "AWS secret access key",
				Computed:            true,
				Sensitive:           true,
			},
			"session_token": schema.StringAttribute{
				MarkdownDescription: "AWS session token",
				Computed:            true,
				Sensitive:           true,
			},
		},
	}
}

func (d *CredentialsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*http.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *CredentialsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ExampleDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	profile := data.Profile.ValueString()

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error loading AWS config",
			fmt.Sprintf("Unable to load AWS config for profile %s: %v", profile, err),
		)
		return
	}

	credentials, err := cfg.Credentials.Retrieve(context.TODO())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error retrieving AWS credentials",
			fmt.Sprintf("Unable to retrieve AWS credentials for profile %s: %v", profile, err),
		)
		return
	}

	data.AccessKeyID = types.StringValue(credentials.AccessKeyID)
	data.SecretAccessKey = types.StringValue(credentials.SecretAccessKey)
	data.SessionToken = types.StringValue(credentials.SessionToken)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
