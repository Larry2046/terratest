// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformAzureLoadBalancerExample(t *testing.T) {
	t.Parallel()

	// loadbalancer::tag::1:: Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-azure-loadbalancer-example",
	}

	// config
	FrontendIPAllocationMethod := "Dynamic"

	// loadbalancer::tag::4:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// loadbalancer::tag::2:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// loadbalancer::tag::3:: Run `terraform output` to get the values of output variables
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	loadBalancer01Name := terraform.Output(t, terraformOptions, "loadbalancer01_name")
	loadBalancer02Name := terraform.Output(t, terraformOptions, "loadbalancer02_name")

	frontendIPConfigForLB01 := terraform.Output(t, terraformOptions, "lb01_feconfig")
	publicIPAddressForLB01 := terraform.Output(t, terraformOptions, "pip_forlb01")

	frontendIPConfigForLB02 := terraform.Output(t, terraformOptions, "feIPConfig_forlb02")
	frontendIPAllocForLB02 := "Static"
	frontendSubnetID := terraform.Output(t, terraformOptions, "feSubnet_forlb02")

	// loadbalancer::tag::5 Set expected variables for test

	// happy path tests
	t.Run("Load Balancer 01", func(t *testing.T) {
		// load balancer 01 (with Public IP) exists
		lb01Exists, err := azure.LoadBalancerExistsE(loadBalancer01Name, resourceGroupName, "")
		assert.NoError(t, err, "Load Balancer error.")
		assert.True(t, lb01Exists)

	})

	t.Run("Frontend Config for LB01", func(t *testing.T) {
		// Read the LB information
		lb01, err := azure.GetLoadBalancerE(loadBalancer01Name, resourceGroupName, "")
		require.NoError(t, err)
		lb01Props := lb01.LoadBalancerPropertiesFormat
		fe01Config := (*lb01Props.FrontendIPConfigurations)[0]
		//fe01Props := *fe01Config.FrontendIPConfigurationPropertiesFormat

		// Verify settings
		assert.Equal(t, frontendIPConfigForLB01, *fe01Config.Name, "LB01 Frontend IP config name")
	})

	t.Run("IP Checks for LB01", func(t *testing.T) {
		// Read the LB information
		lb01, err := azure.GetLoadBalancerE(loadBalancer01Name, resourceGroupName, "")
		require.NoError(t, err)
		lb01Props := lb01.LoadBalancerPropertiesFormat
		fe01Config := (*lb01Props.FrontendIPConfigurations)[0]
		fe01Props := *fe01Config.FrontendIPConfigurationPropertiesFormat

		// Ensure PrivateIPAddress is nil for LB01
		assert.Nil(t, fe01Props.PrivateIPAddress, "LB01 shouldn't have PrivateIPAddress")

		// Ensure PublicIPAddress Resource exists, no need to check PublicIPAddress value
		publicIPAddressResource, err := azure.GetPublicIPAddressE(resourceGroupName, publicIPAddressForLB01, "")
		require.NoError(t, err)
		assert.NotNil(t, publicIPAddressResource, fmt.Sprintf("Public IP Resource for LB01 Frontend: %s", publicIPAddressForLB01))

		// Verify that expected PublicIPAddressResource is assigned to Load Balancer
		pipResourceName, err := GetSliceLastValueLocal(*fe01Props.PublicIPAddress.ID, "/")
		require.NoError(t, err)
		assert.Equal(t, publicIPAddressForLB01, pipResourceName, "LB01 Public IP Address Resource Name")

		assert.Equal(t, FrontendIPAllocationMethod, string(fe01Props.PrivateIPAllocationMethod), "LB01 Frontend IP allocation method")
		assert.Nil(t, fe01Props.Subnet, "LB01 shouldn't have Subnet")
	})

	t.Run("Load Balancer 02", func(t *testing.T) {
		// load balancer 02 (with Private IP on vnet/subnet) exists
		lb02Exists, err := azure.LoadBalancerExistsE(loadBalancer02Name, resourceGroupName, "")
		assert.NoError(t, err, "Load Balancer error.")
		assert.True(t, lb02Exists)
	})

	t.Run("IP Check for Load Balancer 02", func(t *testing.T) {
		// Read LB02 information
		lb02, err := azure.GetLoadBalancerE(loadBalancer02Name, resourceGroupName, "")
		require.NoError(t, err)
		lb02Props := lb02.LoadBalancerPropertiesFormat
		fe02Config := (*lb02Props.FrontendIPConfigurations)[0]
		fe02Props := *fe02Config.FrontendIPConfigurationPropertiesFormat

		assert.Equal(t, frontendIPConfigForLB02, *fe02Props.PrivateIPAddress, "LB02 Frontend IP address")
		assert.Equal(t, frontendIPAllocForLB02, string(fe02Props.PrivateIPAllocationMethod), "LB02 Frontend IP allocation method")
		subnetID, err := GetSliceLastValueLocal(*fe02Props.Subnet.ID, "/")
		require.NoError(t, err, "LB02 Frontend subnet not found")
		frontendSubnetID, err := GetSliceLastValueLocal(frontendSubnetID, "/")
		assert.Equal(t, frontendSubnetID, subnetID, "LB02 Frontend subnet ID")
	})

}

// GetSliceLastValue will take a source string and returns the last value when split by the seperaror char
func GetSliceLastValueLocal(source string, seperator string) (string, error) {
	if !(len(source) == 0 || len(seperator) == 0 || !strings.Contains(source, seperator)) {
		tmp := strings.Split(source, seperator)
		return tmp[len(tmp)-1], nil
	}
	return "", errors.New("invalid input or no slice available")
}
