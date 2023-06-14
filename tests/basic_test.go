package test

import (
  "testing"

  "github.com/gruntwork-io/terratest/modules/docker"
  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/stretchr/testify/assert"
)

func TestTerraformDockerSimple(t *testing.T) {
  containerName := "single_port"


  dockerNetworks := map[string]interface{}{
    "nginx_network": map[string]interface{}{
      "ipam_config": map[string]interface{}{
        "aux_address": map[string]interface{}{},
        "gateway": "10.0.30.1",
        "subnet": "10.0.30.0/24",
      },
    },
  }
  networksAdvanced := []map[string]interface{}{
    {
      "name": "nginx_network",
      "ipv4_address": "10.0.30.2",
    },
  }
  ports := []map[string]interface{}{
    {
      "internal": 80,
      "external": 80,
      "protocol": "tcp",
    },
  }
  namedVolumes := map[string]interface{}{
    "nginx_config": map[string]interface{}{
      "container_path": "/etc/nginx",
      "read_only":      false,
      "create":         true,
    },
  }
  hostPaths := map[string]interface{}{
    "/media/local": map[string]interface{}{
      "container_path": "/mnt",
      "read_only":      true,
    },
  }
  devices := map[string]interface{}{
    "/dev/null": map[string]interface{}{
      "container_path": "/dev/null",
      "permissions":    "rwm",
    },
  }

  terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../examples/basic",

    Vars: map[string]interface{}{
      "container_name": containerName,
      "docker_networks": dockerNetworks,
      "networks_advanced": networksAdvanced,
      "ports": ports,
      "named_volumes": namedVolumes,
      "host_paths": hostPaths,
      "devices": devices,
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
  dockerNetworks := map[string]interface{}{
    "nginx_network": map[string]interface{}{
      "ipam_config": map[string]interface{}{
        "aux_address": map[string]interface{}{},
        "gateway": "10.0.30.1",
        "subnet": "10.0.30.0/24",
      },
    },
  }
  networksAdvanced := []map[string]interface{}{
    {
      "name": "nginx_network",
      "ipv4_address": "10.0.30.2",
    },
  }
  ports := []map[string]interface{}{
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
  }
  namedVolumes := map[string]interface{}{
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
  }
  hostPaths := map[string]interface{}{
    "/media/local": map[string]interface{}{
      "container_path": "/mnt",
      "read_only":      true,
    },
    "/opt/test": map[string]interface{}{
      "container_path": "/opt/test",
      "read_only":      false,
    },
  }
  devices := map[string]interface{}{
    "/dev/null": map[string]interface{}{
      "container_path": "/dev/null",
      "permissions":    "rwm",
    },
    "/dev/random": map[string]interface{}{
      "container_path": "/dev/random",
      "permissions":    "rwm",
    },
  }

  terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../examples/basic",

    Vars: map[string]interface{}{
      "container_name": containerName,
      "docker_networks": dockerNetworks,
      "networks_advanced": networksAdvanced,
      "ports": ports,
      "named_volumes": namedVolumes,
      "host_paths": hostPaths,
      "devices": devices,
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
  dockerNetworks := map[string]interface{}{
    "nginx_network": map[string]interface{}{
      "ipam_config": map[string]interface{}{
        "aux_address": map[string]interface{}{},
        "gateway": "10.0.30.1",
        "subnet": "10.0.30.0/24",
      },
    },
  }
  networksAdvanced := []map[string]interface{}{
    {
      "name": "nginx_network",
      "ipv4_address": "10.0.30.2",
    },
  }

  terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../examples/basic",

    Vars: map[string]interface{}{
      "container_name": containerName,
      "docker_networks": dockerNetworks,
      "networks_advanced": networksAdvanced,
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
  dockerNetworks := map[string]interface{}{
    "nginx_network": map[string]interface{}{
      "ipam_config": map[string]interface{}{
        "aux_address": map[string]interface{}{},
        "gateway": "10.0.30.1",
        "subnet": "10.0.30.0/24",
      },
    },
  }
  networksAdvanced := []map[string]interface{}{
    {
      "name": "nginx_network",
      "ipv4_address": "10.0.30.2",
    },
  }
  environment := map[string]interface{}{
    "ENV":  "test",
    "NAME": "nginx",
  }

  terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../examples/basic",

    Vars: map[string]interface{}{
      "container_name": containerName,
      "docker_networks": dockerNetworks,
      "networks_advanced": networksAdvanced,
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
