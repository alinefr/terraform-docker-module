package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformDockerSimple(t *testing.T) {
	containerName := "single_port"
	gateway := "10.0.30.1"
	ipv4Address := "10.0.30.2"
	subnet := "10.0.30.0/24"
	networkName := "nginx_network"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",

		Vars: map[string]interface{}{
			"container_name": containerName,
			"gateway":        gateway,
			"ipv4_address":   ipv4Address,
			"subnet":         subnet,
			"network_name":   networkName,
			"ports": []map[string]interface{}{
				{
					"internal": 80,
					"external": 80,
					"protocol": "tcp",
				},
			},
			"named_volumes": map[string]interface{}{
				"nginx_config": map[string]interface{}{
					"container_path": "/etc/nginx",
					"read_only":      false,
					"create":         true,
				},
			},
			"host_paths": map[string]interface{}{
				"/media/local": map[string]interface{}{
					"container_path": "/mnt",
					"read_only":      true,
				},
			},
			"devices": map[string]interface{}{
				"/dev/null": map[string]interface{}{
					"container_path": "/dev/null",
					"permissions":    "rwm",
				},
			},
		},

		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	inspectOutput := docker.Inspect(
		t,
		containerName,
	)

	expectedPorts := []docker.Port{
		{
			HostPort:      80,
			ContainerPort: 80,
			Protocol:      "tcp",
		},
	}
	actualPorts := inspectOutput.Ports

	expectedVolumes := []map[string]interface{}{
		{
			"container_path": "/etc/nginx",
			"from_container": "",
			"host_path":      "",
			"read_only":      false,
			"volume_name":    "nginx_config",
		},
		{
			"container_path": "/mnt",
			"from_container": "",
			"host_path":      "/media/local",
			"read_only":      true,
			"volume_name":    "",
		},
	}
	actualVolumes := terraform.OutputListOfObjects(t, terraformOptions, "volumes")

	expectedDevices := []map[string]interface{}{
		{
			"container_path": "/dev/null",
			"host_path":      "/dev/null",
			"permissions":    "rwm",
		},
	}
	actualDevices := terraform.OutputListOfObjects(t, terraformOptions, "devices")

	assert.Equal(t, expectedPorts, actualPorts)
	assert.Equal(t, expectedVolumes, actualVolumes)
	assert.Equal(t, expectedDevices, actualDevices)
}

func TestTerraformDockerMultiple(t *testing.T) {
	containerName := "multiple_ports"
	gateway := "10.0.30.1"
	ipv4Address := "10.0.30.2"
	subnet := "10.0.30.0/24"
	networkName := "nginx_network"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",

		Vars: map[string]interface{}{
			"container_name": containerName,
			"gateway":        gateway,
			"ipv4_address":   ipv4Address,
			"subnet":         subnet,
			"network_name":   networkName,
			"ports": []map[string]interface{}{
				{
					"internal": 80,
					"external": 80,
					"protocol": "tcp",
				},
				{
					"internal": 443,
					"external": 443,
					"protocol": "tcp",
				},
			},
			"named_volumes": map[string]interface{}{
				"nginx_config": map[string]interface{}{
					"container_path": "/etc/nginx",
					"read_only":      false,
					"create":         true,
				},
				"nginx_html": map[string]interface{}{
					"container_path": "/var/www/html",
					"read_only":      false,
					"create":         true,
				},
			},
			"host_paths": map[string]interface{}{
				"/media/local": map[string]interface{}{
					"container_path": "/mnt",
					"read_only":      true,
				},
				"/opt/test": map[string]interface{}{
					"container_path": "/opt/test",
					"read_only":      false,
				},
			},
			"devices": map[string]interface{}{
				"/dev/null": map[string]interface{}{
					"container_path": "/dev/null",
					"permissions":    "rwm",
				},
				"/dev/random": map[string]interface{}{
					"container_path": "/dev/random",
					"permissions":    "rwm",
				},
			},
		},

		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := docker.Inspect(
		t,
		containerName,
	)

	expectedPorts := []docker.Port{
		{
			HostPort:      443,
			ContainerPort: 443,
			Protocol:      "tcp",
		},
		{
			HostPort:      80,
			ContainerPort: 80,
			Protocol:      "tcp",
		},
	}
	actualPorts := output.Ports

	expectedVolumes := []map[string]interface{}{
		{
			"container_path": "/etc/nginx",
			"from_container": "",
			"host_path":      "",
			"read_only":      false,
			"volume_name":    "nginx_config",
		},
		{
			"container_path": "/mnt",
			"from_container": "",
			"host_path":      "/media/local",
			"read_only":      true,
			"volume_name":    "",
		},
		{
			"container_path": "/opt/test",
			"from_container": "",
			"host_path":      "/opt/test",
			"read_only":      false,
			"volume_name":    "",
		},
		{
			"container_path": "/var/www/html",
			"from_container": "",
			"host_path":      "",
			"read_only":      false,
			"volume_name":    "nginx_html",
		},
	}
	actualVolumes := terraform.OutputListOfObjects(t, terraformOptions, "volumes")

	expectedDevices := []map[string]interface{}{
		{
			"container_path": "/dev/null",
			"host_path":      "/dev/null",
			"permissions":    "rwm",
		},
		{
			"container_path": "/dev/random",
			"host_path":      "/dev/random",
			"permissions":    "rwm",
		},
	}
	actualDevices := terraform.OutputListOfObjects(t, terraformOptions, "devices")

	assert.Equal(t, expectedPorts, actualPorts)
	assert.Equal(t, expectedVolumes, actualVolumes)
	assert.Equal(t, expectedDevices, actualDevices)
}

func TestTerraformDockerWithoutPortsVolumesDevices(t *testing.T) {
	containerName := "without"
	gateway := "10.0.30.1"
	ipv4Address := "10.0.30.2"
	subnet := "10.0.30.0/24"
	networkName := "nginx_network"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",

		Vars: map[string]interface{}{
			"container_name": containerName,
			"gateway":        gateway,
			"ipv4_address":   ipv4Address,
			"subnet":         subnet,
			"network_name":   networkName,
		},

		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	inspectOutput := docker.Inspect(
		t,
		containerName,
	)

	expectedPorts := []docker.Port(nil)
	actualPorts := inspectOutput.Ports

	expectedVolumes := []map[string]interface{}(nil)
	actualVolumes := terraform.OutputListOfObjects(t, terraformOptions, "volumes")

	expectedDevices := []map[string]interface{}(nil)
	actualDevices := terraform.OutputListOfObjects(t, terraformOptions, "devices")

	assert.Equal(t, expectedPorts, actualPorts)
	assert.Equal(t, expectedVolumes, actualVolumes)
	assert.Equal(t, expectedDevices, actualDevices)
}

func TestTerraformDockerEnvironments(t *testing.T) {
	containerName := "environments"
	gateway := "10.0.30.1"
	ipv4Address := "10.0.30.2"
	subnet := "10.0.30.0/24"
	networkName := "nginx_network"
	environment := map[string]interface{}{
		"ENV":  "test",
		"NAME": "nginx",
	}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",

		Vars: map[string]interface{}{
			"container_name": containerName,
			"gateway":        gateway,
			"ipv4_address":   ipv4Address,
			"subnet":         subnet,
			"network_name":   networkName,
			"environment":    environment,
		},

		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	expectedEnv := []string{
		"ENV=test",
		"NAME=nginx",
	}
	actualEnv := terraform.OutputList(t, terraformOptions, "environment")
	assert.Equal(t, expectedEnv, actualEnv)
}
