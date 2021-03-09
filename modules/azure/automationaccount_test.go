// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when methods to create and delete automation account resources are added, these tests can be extended.
*/

func TestGetAutomationAccountClientE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	client, err := GetAutomationAccountClientE(subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}

func TestAutomationAccountExistsE(t *testing.T) {
	t.Parallel()

	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := AutomationAccountExistsE(t, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
}

func TestAutomationAccountDscExistsE(t *testing.T) {
	t.Parallel()

	dscConfiguraitonName := ""
	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := AutomationAccountDscExistsE(t, dscConfiguraitonName, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
}

func TestAutomationAccountDscCompiledE(t *testing.T) {
	t.Parallel()

	dscConfiguraitonName := ""
	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := AutomationAccountDscCompiledE(t, dscConfiguraitonName, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
}

func TestAutomationAccountRunAsCertificateThumbprintMatchesE(t *testing.T) {
	t.Parallel()

	runAsCertificateThumbprint := ""
	runAsCertificateName := ""
	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := AutomationAccountRunAsCertificateThumbprintMatchesE(t, runAsCertificateThumbprint, runAsCertificateName, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
}

func TestAutomationAccountRunAsConnectionValidatesE(t *testing.T) {
	t.Parallel()

	automationAccountrunAsAccountName := ""
	runAsConnectionType := ""
	runAsCertificateThumbprint := ""
	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := AutomationAccountRunAsConnectionValidatesE(t, automationAccountrunAsAccountName, runAsConnectionType, runAsCertificateThumbprint, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
}

func TestAutomationAccountDscAppliedSuccessfullyToVME(t *testing.T) {
	t.Parallel()

	dscConfiguraitonName := ""
	vmName := ""
	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := AutomationAccountDscAppliedSuccessfullyToVME(t, dscConfiguraitonName, vmName, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
}

func TestGetAutomationAccountE(t *testing.T) {
	t.Parallel()

	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	client, err := GetAutomationAccountE(t, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}

func TestGetAutomationAccountDscConfigurationE(t *testing.T) {
	t.Parallel()

	dscConfigurationName := ""
	resourceGroupName := ""
	automationAccountName := ""
	subscriptionID := ""

	dscConfiguration, err := GetAutomationAccountDscConfigurationE(t, dscConfigurationName, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *dscConfiguration)
}

func TestAutomationAccountDscCompileJobStatusE(t *testing.T) {
	t.Parallel()

	dscConfigurationName := ""
	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	status, err := AutomationAccountDscCompileJobStatusE(t, dscConfigurationName, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, status)
}

func TestGetAutomationAccountCertificateE(t *testing.T) {
	t.Parallel()

	automationAccountCertificateName := ""
	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	certificate, err := GetAutomationAccountCertificateE(t, automationAccountCertificateName, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *certificate)
}

func TestGetAutomationAccountDscNodeConfigurationE(t *testing.T) {
	t.Parallel()

	dscConfiguraitonName := ""
	vmName := ""
	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	dscNodeConfig, err := GetAutomationAccountDscNodeConfigurationE(t, dscConfiguraitonName, vmName, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *dscNodeConfig)
}

func TestGetAutomationAccountRunAsConnectionE(t *testing.T) {
	t.Parallel()

	automationAccountRunAsConnectionName := ""
	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	connection, err := GetAutomationAccountRunAsConnectionE(t, automationAccountRunAsConnectionName, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *connection)
}

func TestGetCertificateClientE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	client, err := GetCertificateClientE(subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}

func TestGetDscConfigurationClientE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	client, err := GetDscConfigurationClientE(subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}

func TestGetDscCompilationJobClientE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	client, err := GetDscCompilationJobClientE(subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}

func TestGetAutomationAccountCertficateClientE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	client, err := GetAutomationAccountCertficateClientE(subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}

func TestGetAutomationAccountRunAsConnectionClientE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	client, err := GetAutomationAccountRunAsConnectionClientE(subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}

func TestGetAutomationAccountDscNodeConfigClientE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	client, err := GetAutomationAccountDscNodeConfigClientE(subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}
