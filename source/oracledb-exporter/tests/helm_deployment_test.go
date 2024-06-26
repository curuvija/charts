package test

import (
	"fmt"
	"os"

	//"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"

	//corev1 "k8s.io/api/core/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
)

// https://github.com/gruntwork-io/terratest/blob/master/test/helm_basic_example_template_test.go

func TestDeployment(t *testing.T) {
	//t.Parallel()
	// Path to the helm chart we will test
	helmChartPath := os.Getenv("HELM_CHART_PATH")
	releaseName := os.Getenv("RELEASE_NAME")
	// require.NoError(t, err)

	// Setup the args. For this test, we will set the following input values:
	options := &helm.Options{
		SetValues: map[string]string{
			"image.repository": "ghcr.io/iamseth/oracledb-exporter",
			"image.tag":        "0.5.1",
		},
	}

	// Run RenderTemplate to render the template and capture the output.
	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/deployment.yaml"})
	fmt.Println(output)

	// Now we use kubernetes/client-go library to render the template output into the Pod struct. This will
	// ensure the Pod resource is rendered correctly.
	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)

	expectedContainerImage := "ghcr.io/iamseth/oracledb-exporter:0.5.1"
	deploymentContainers := deployment.Spec.Template.Spec.Containers
	require.Equal(t, len(deploymentContainers), 1)
	require.Equal(t, deploymentContainers[0].Image, expectedContainerImage)
}
