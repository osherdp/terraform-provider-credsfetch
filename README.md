# Terraform Provider CredsFetch

This is a helper provider for an issue `terraform-provider-vault` (and maybe other providers) has, where native support for SSO or IRSA is not fully implemented.

This gives the option to fetch AWS credentials for a specific profile, and use it for the AWS authentication. For example:
```
terraform {
  required_providers {
    vault = {
      source  = "hashicorp/vault"
      version = "~> 4.5.0"
    }
    credsfetch = {
      source  = "osherdp/credsfetch"
      version = "~> 0.1.0"
    }
  }

data "credsfetch_credentials" "staging" {
  profile = "staging"
}

provider "vault" {
  address = "<some-address>"
  auth_login_aws {
    aws_access_key_id     = data.credsfetch_credentials.staging.access_key_id
    aws_secret_access_key = data.credsfetch_credentials.staging.secret_access_key
    aws_session_token     = data.credsfetch_credentials.staging.session_token
    role                  = "<some role>"
  }
}
```

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.23

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

Fill this in for each provider

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
