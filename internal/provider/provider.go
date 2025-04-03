package provider

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/provider"
    "github.com/hashicorp/terraform-plugin-framework/provider/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
    _ provider.Provider = &CredsFetchProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
    return func() provider.Provider {
        return &CredsFetchProvider{
            version: version,
        }
    }
}

// CredsFetchProvider is the provider implementation.
type CredsFetchProvider struct {
    // version is set to the provider version on release, "dev" when the
    // provider is built and ran locally, and "test" when running acceptance
    // testing.
    version string
}

// Metadata returns the provider type name.
func (p *CredsFetchProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
    resp.TypeName = "credsfetch"
    resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *CredsFetchProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
    resp.Schema = schema.Schema{}
}

// Configure prepares a AWSCredentials API client for data sources and resources.
func (p *CredsFetchProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *CredsFetchProvider) DataSources(_ context.Context) []func() datasource.DataSource {
    return []func() datasource.DataSource{
        func() datasource.DataSource {
            return NewExampleDataSource()
        },
    }
}

// Resources defines the resources implemented in the provider.
func (p *CredsFetchProvider) Resources(_ context.Context) []func() resource.Resource {
    return nil
}
