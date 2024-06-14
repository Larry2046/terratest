package k8s

import (
	"net/url"
	"os"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/shell"
	"github.com/gruntwork-io/terratest/modules/testing"
)

// RunKubectl will call kubectl using the provided options and args, failing the test on error.
func RunKubectl(t testing.TestingT, options *KubectlOptions, args ...string) {
	require.NoError(t, RunKubectlE(t, options, args...))
}

// RunKubectlE will call kubectl using the provided options and args.
func RunKubectlE(t testing.TestingT, options *KubectlOptions, args ...string) error {
	_, err := RunKubectlAndGetOutputE(t, options, args...)
	return err
}

// RunKubectlWithRetry will call kubectl using the provided options and args, failing the test after errors have been retried.
func RunKubectlWithRetry(t testing.TestingT, options *KubectlOptions, retries int, sleepBetweenRetries time.Duration, args ...string) {
	err := RunKubectlWithRetryE(t, options, retries, sleepBetweenRetries, args...)
	require.NoError(t, err)
}

// RunKubectlWithRetryE will call kubectl using the provided options and args, reytring on error.
func RunKubectlWithRetryE(t testing.TestingT, options *KubectlOptions, retries int, sleepBetweenRetries time.Duration, args ...string) error {
	_, err := retry.DoWithRetryE(
		t,
		"",
		retries,
		sleepBetweenRetries,
		func() (string, error) {
			return "", RunKubectlE(t, options, args...)
		},
	)
	return err
}

// RunKubectlAndGetOutputE will call kubectl using the provided options and args, returning the output of stdout and
// stderr.
func RunKubectlAndGetOutputE(t testing.TestingT, options *KubectlOptions, args ...string) (string, error) {
	cmdArgs := []string{}
	if options.ContextName != "" {
		cmdArgs = append(cmdArgs, "--context", options.ContextName)
	}
	if options.ConfigPath != "" {
		cmdArgs = append(cmdArgs, "--kubeconfig", options.ConfigPath)
	}
	if options.Namespace != "" {
		cmdArgs = append(cmdArgs, "--namespace", options.Namespace)
	}
	cmdArgs = append(cmdArgs, args...)
	command := shell.Command{
		Command: "kubectl",
		Args:    cmdArgs,
		Env:     options.Env,
		Logger:  options.Logger,
	}
	return shell.RunCommandAndGetOutputE(t, command)
}

// RunKubectlAndGetOutputWithRetry will call kubectl using the provided options and args, returning the output of stdout and
// stderr, failing the test after errors have been retried.
func RunKubectlAndGetOutputWithRetry(t testing.TestingT, options *KubectlOptions, retries int, sleepBetweenRetries time.Duration, args ...string) string {
	output, err := RunKubectlAndGetOutputWithRetryE(t, options, retries, sleepBetweenRetries, args...)
	require.NoError(t, err)
	return output
}

// RunKubectlAndGetOutputWithRetryE will call kubectl using the provided options and args, returning the output of stdout and
// stderr, reytring on error.
func RunKubectlAndGetOutputWithRetryE(t testing.TestingT, options *KubectlOptions, retries int, sleepBetweenRetries time.Duration, args ...string) (string, error) {
	output, err := retry.DoWithRetryE(
		t,
		"",
		retries,
		sleepBetweenRetries,
		func() (string, error) {
			return RunKubectlAndGetOutputE(t, options, args...)
		},
	)
	return output, err
}

// KubectlDelete will take in a file path and delete it from the cluster targeted by KubectlOptions. If there are any
// errors, fail the test immediately.
func KubectlDelete(t testing.TestingT, options *KubectlOptions, configPath string) {
	require.NoError(t, KubectlDeleteE(t, options, configPath))
}

// KubectlDeleteE will take in a file path and delete it from the cluster targeted by KubectlOptions.
func KubectlDeleteE(t testing.TestingT, options *KubectlOptions, configPath string) error {
	return RunKubectlE(t, options, "delete", "-f", configPath)
}

// KubectlDeleteFromKustomize will take in a kustomization directory path and delete it from the cluster targeted by KubectlOptions. If there are any
// errors, fail the test immediately.
func KubectlDeleteFromKustomize(t testing.TestingT, options *KubectlOptions, configPath string) {
	require.NoError(t, KubectlDeleteFromKustomizeE(t, options, configPath))
}

// KubectlDeleteFromKustomizeE will take in a kustomization directory path and delete it from the cluster targeted by KubectlOptions.
func KubectlDeleteFromKustomizeE(t testing.TestingT, options *KubectlOptions, configPath string) error {
	return RunKubectlE(t, options, "delete", "-k", configPath)
}

// KubectlDeleteFromString will take in a kubernetes resource config as a string and delete it on the cluster specified
// by the provided kubectl options.
func KubectlDeleteFromString(t testing.TestingT, options *KubectlOptions, configData string) {
	require.NoError(t, KubectlDeleteFromStringE(t, options, configData))
}

// KubectlDeleteFromStringE will take in a kubernetes resource config as a string and delete it on the cluster specified
// by the provided kubectl options. If it fails, this will return the error.
func KubectlDeleteFromStringE(t testing.TestingT, options *KubectlOptions, configData string) error {
	tmpfile, err := StoreConfigToTempFileE(t, configData)
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile)
	return KubectlDeleteE(t, options, tmpfile)
}

// KubectlApply will take in a file path and apply it to the cluster targeted by KubectlOptions. If there are any
// errors, fail the test immediately.
func KubectlApply(t testing.TestingT, options *KubectlOptions, configPath string) {
	require.NoError(t, KubectlApplyE(t, options, configPath))
}

// KubectlApplyE will take in a file path and apply it to the cluster targeted by KubectlOptions.
func KubectlApplyE(t testing.TestingT, options *KubectlOptions, configPath string) error {
	return RunKubectlE(t, options, "apply", "-f", configPath)
}

// KubectlApplyFromKustomize will take in a kustomization directory path and apply it to the cluster targeted by KubectlOptions. If there are any
// errors, fail the test immediately.
func KubectlApplyFromKustomize(t testing.TestingT, options *KubectlOptions, configPath string) {
	require.NoError(t, KubectlApplyFromKustomizeE(t, options, configPath))
}

// KubectlApplyFromKustomizeE will take in a kustomization directory path and apply it to the cluster targeted by KubectlOptions.
func KubectlApplyFromKustomizeE(t testing.TestingT, options *KubectlOptions, configPath string) error {
	return RunKubectlE(t, options, "apply", "-k", configPath)
}

// KubectlApplyFromString will take in a kubernetes resource config as a string and apply it on the cluster specified
// by the provided kubectl options.
func KubectlApplyFromString(t testing.TestingT, options *KubectlOptions, configData string) {
	require.NoError(t, KubectlApplyFromStringE(t, options, configData))
}

// KubectlApplyFromStringE will take in a kubernetes resource config as a string and apply it on the cluster specified
// by the provided kubectl options. If it fails, this will return the error.
func KubectlApplyFromStringE(t testing.TestingT, options *KubectlOptions, configData string) error {
	tmpfile, err := StoreConfigToTempFileE(t, configData)
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile)
	return KubectlApplyE(t, options, tmpfile)
}

// KubectlApplyFromStringWithRetry will take in a kubernetes resource config as a string and apply it on the cluster specified
// by the provided kubectl options. The test fails after errors have been retried.
func KubectlApplyFromStringWithRetry(t testing.TestingT, options *KubectlOptions, retries int, sleepBetweenRetries time.Duration, configData string) {
	err := KubectlApplyFromStringWithRetryE(t, options, retries, sleepBetweenRetries, configData)
	require.NoError(t, err)
}

// KubectlApplyFromStringWithRetryE will take in a kubernetes resource config as a string and apply it on the cluster specified
// by the provided kubectl options. Errors are retried.
func KubectlApplyFromStringWithRetryE(t testing.TestingT, options *KubectlOptions, retries int, sleepBetweenRetries time.Duration, configData string) error {
	_, err := retry.DoWithRetryE(
		t,
		"",
		retries,
		sleepBetweenRetries,
		func() (string, error) {
			return "", KubectlApplyFromStringE(t, options, configData)
		},
	)
	return err
}

// StoreConfigToTempFile will store the provided config data to a temporary file created on the os and return the
// filename.
func StoreConfigToTempFile(t testing.TestingT, configData string) string {
	out, err := StoreConfigToTempFileE(t, configData)
	require.NoError(t, err)
	return out
}

// StoreConfigToTempFileE will store the provided config data to a temporary file created on the os and return the
// filename, or error.
func StoreConfigToTempFileE(t testing.TestingT, configData string) (string, error) {
	escapedTestName := url.PathEscape(t.Name())
	tmpfile, err := os.CreateTemp("", escapedTestName)
	if err != nil {
		return "", err
	}
	defer tmpfile.Close()

	_, err = tmpfile.WriteString(configData)
	return tmpfile.Name(), err
}
