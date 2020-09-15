# Terraform Azure Log Analytics Example

This folder contains a simple Terraform module that deploys resources in [Azure](https://azure.microsoft.com/) to demonstrate how you can use Terratest to write automated tests for your Azure Terraform code. This module deploys the following:


* A [Log Analytics Workspace](https://docs.microsoft.com/en-us/azure/azure-monitor/platform/log-analytics-agent) that gives the module the following:
    * [Name](https://docs.microsoft.com/en-us/azure/azure-monitor/learn/quick-create-workspace#:~:text=%20Create%20a%20Log%20Analytics%20workspace%20in%20the,and%20region%20as%20in%20the%20deleted...%20More%20)  with the value specified in the `loganalytics_workspace_name`  output variable.
    * [Sku](https://docs.microsoft.com/en-us/azure/azure-monitor/learn/quick-create-workspace#:~:text=%20Create%20a%20Log%20Analytics%20workspace%20in%20the,and%20region%20as%20in%20the%20deleted...%20More%20)  with the value specified in the `loganalytics_workspace_sku`  output variable.
    [RentionPeriodInDays](https://docs.microsoft.com/en-us/azure/azure-monitor/learn/quick-create-workspace#:~:text=%20Create%20a%20Log%20Analytics%20workspace%20in%20the,and%20region%20as%20in%20the%20deleted...%20More%20)  with the value specified in the `loganalytics_workspace_retention`  output variable.

Check out [test/azure/terraform_azure_@module@_test.go](/test/azure/terraform_azure_@module@_test.go) to see how you can write
automated tests for this module.

Note that the resources deployed in this module don't actually do anything; it just runs the resources for
demonstration purposes.

**WARNING**: This module and the automated tests for it deploy real resources into your Azure account which can cost you
money. The resources are all part of the [Azure Free Account](https://azure.microsoft.com/en-us/free/), so if you haven't used that up,
it should be free, but you are completely responsible for all Azure charges.

## Running this module manually

1. Sign up for [Azure](https://azure.microsoft.com/).
1. Configure your Azure credentials using one of the [supported methods for Azure CLI
   tools](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest).
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Run `terraform init`.
1. Run `terraform apply`.
1. When you're done, run `terraform destroy`.

## Running automated tests against this module

1. Sign up for [Azure](https://azure.microsoft.com/).
1. Configure your Azure credentials using one of the [supported methods for Azure CLI
   tools](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest).
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. [Review environment variables](#review-environment-variables).
1. Install [Golang](https://golang.org/) and make sure this code is checked out into your `GOPATH`.
1. `cd test`
1. Make sure [the azure-sdk-for-go versions match](#check-go-dependencies) in [/test/go.mod](/test/go.mod) and in [test/azure/terraform_azure_@module@_test.go](/test/terraform_azure_nic_test.go).
1. `go build terraform_azure_@module@_test.go`
1. `go test -v -run TestTerraformAzure@Module@Example`

## Module test APIs


- func LogAnalyticsWorkspaceExists(workspaceName, resourceGroupName, subscriptionID string) bool
    </br><font color="green">LogAnalyticsWorkspaceExists indicates whether the operatonal insights
    workspaces exists; otherwise false or error</font>

- func GetLogAnalyticsWorkspaceSku(workspaceName, resourceGroupName, subscriptionID string) string
    </br><font color="green">GetLogAnalyticsWorkspaceSku return the log analytics workspace SKU as string
    as one of the following: Free, Standard, Premium, PerGB2018,
    CapacityReservation; otherwise empty string ""</font>

- func GetLogAnalyticsWorkspaceRetentionPeriodDays(workspaceName, resourceGroupName, subscriptionID string) int32
    </br><font color="green">GetLogAnalyticsWorkspaceRetentionPeriodDays returns the log analytics
    workspace retention period in days; otherwise -1</font>

- func GetLogAnalyticsWorkspaceE(workspaceName, resoureGroupName, subscriptionID string) (*operationalinsights.Workspace, error)
    </br><font color="green">KeyVaultCertificateExistsE returns true if the Azure Key Vault Certificate exists</font>
- func GetLogAnalyticsWorkspacesClientE(subscriptionID string) (*operationalinsights.WorkspacesClient, error)
    </br><font color="green">GetLogAnalyticsWorkspacesClientE return workspaces client; otherwise error</font>

## Check Go Dependencies

Check that the `github.com/Azure/azure-sdk-for-go` version in your generated `go.mod` for this test matches the version in the terratest [go.mod](https://github.com/gruntwork-io/terratest/blob/master/go.mod) file.  

> This was tested with **go1.14.4**.

### Check Azure-sdk-for-go version

Let's make sure [go.mod](https://github.com/gruntwork-io/terratest/blob/master/go.mod) includes the appropriate [azure-sdk-for-go version](https://github.com/Azure/azure-sdk-for-go/releases/tag/v38.1.0):

```go
require (
    ...
    github.com/Azure/azure-sdk-for-go v38.1.0+incompatible
    ...
)
```

If we make changes to either the **go.mod** or the **go test file**, we should make sure that the go build command works still.

```powershell
go build terraform_azure_@module_test.go
```

## Review Environment Variables

As part of configuring terraform for Azure, we'll want to check that we have set the appropriate [credentials](https://docs.microsoft.com/en-us/azure/terraform/terraform-install-configure?toc=https%3A%2F%2Fdocs.microsoft.com%2Fen-us%2Fazure%2Fterraform%2Ftoc.json&bc=https%3A%2F%2Fdocs.microsoft.com%2Fen-us%2Fazure%2Fbread%2Ftoc.json#set-up-terraform-access-to-azure) and also that we set the [environment variables](https://docs.microsoft.com/en-us/azure/terraform/terraform-install-configure?toc=https%3A%2F%2Fdocs.microsoft.com%2Fen-us%2Fazure%2Fterraform%2Ftoc.json&bc=https%3A%2F%2Fdocs.microsoft.com%2Fen-us%2Fazure%2Fbread%2Ftoc.json#configure-terraform-environment-variables) on the testing host.

```bash
export ARM_CLIENT_ID=your_app_id
export ARM_CLIENT_SECRET=your_password
export ARM_SUBSCRIPTION_ID=your_subscription_id
export ARM_TENANT_ID=your_tenant_id
```

Note, in a Windows environment, these should be set as **system environment variables**.  We can use a PowerShell console with administrative rights to update these environment variables:

```powershell
[System.Environment]::SetEnvironmentVariable("ARM_CLIENT_ID",$your_app_id,[System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("ARM_CLIENT_SECRET",$your_password,[System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("ARM_SUBSCRIPTION_ID",$your_subscription_id,[System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("ARM_TENANT_ID",$your_tenant_id,[System.EnvironmentVariableTarget]::Machine)
```






 