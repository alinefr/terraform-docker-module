package test

import (
	"testing"

	"github.com/Jeffail/gabs"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformDockerEnvironments(t *testing.T) {
	image := "nginx:latest"
	containerName := "terraform-test"
	hostname := "terratest"
	workingDir := "/tmp"
	restart_policy := "unless-stopped"
	privileged := false
	network_mode := "bridge"
	dns := []string{"1.1.1.1", "1.0.0.1"}
	entrypoint := []string{"/docker-entrypoint.sh"}
	command := []string{"nginx", "-g", "daemon off;"}
	ports := []map[string]interface{}{
		{
			"internal": 80,
			"external": 9999,
			"protocol": "tcp",
		},
	}
	named_volumes := map[string]interface{}{
		"nginx_volume": map[string]interface{}{
			"container_path": "/etc/nginx",
			"read_only":      false,
			"create":         true,
		},
		"mnt_volume": map[string]interface{}{
			"container_path": "/mnt",
			"read_only":      false,
			"create":         true,
		},
	}
	host_paths := map[string]interface{}{
		"/tmp": map[string]interface{}{
			"container_path": "/tmp",
			"read_only":      true,
		},
	}
	devices := map[string]interface{}{
		"/dev/null": map[string]interface{}{
			"container_path": "/dev/newnull",
			"permissions":    "rwm",
		},
	}
	capabilities := map[string]interface{}{
		"add":  []string{"NET_ADMIN"},
		"drop": []string{},
	}
	dockerNetworks := []map[string]interface{}{
		{
			"name": "nginx_network",
			"ipam_config": map[string]interface{}{
				"aux_address": map[string]interface{}{},
				"gateway":     "10.0.30.1",
				"subnet":      "10.0.30.0/24",
			},
		},
		{
			"name": "app_network",
			"ipam_config": map[string]interface{}{
				"aux_address": map[string]interface{}{},
				"gateway":     "10.0.55.1",
				"subnet":      "10.0.55.0/24",
			},
		},
	}
	networksAdvanced := []map[string]interface{}{
		{
			"name":         "nginx_network",
			"ipv4_address": "10.0.30.2",
		},
		{
			"name":         "app_network",
			"ipv4_address": "10.0.55.2",
		},
	}
	environment := map[string]interface{}{
		"ENV":  "test",
		"NAME": "nginx",
	}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",

		Vars: map[string]interface{}{
			"image":             image,
			"container_name":    containerName,
			"hostname":          hostname,
			"working_dir":       workingDir,
			"restart_policy":    restart_policy,
			"privileged":        privileged,
			"network_mode":      network_mode,
			"dns":               dns,
			"entrypoint":        entrypoint,
			"command":           command,
			"ports":             ports,
			"named_volumes":     named_volumes,
			"host_paths":        host_paths,
			"devices":           devices,
			"capabilities":      capabilities,
			"docker_networks":   dockerNetworks,
			"networks_advanced": networksAdvanced,
			"environment":       environment,
		},

		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	validateOutputs(t, terraformOptions)
}

func validateOutputs(t *testing.T, opts *terraform.Options) {
	jsonParsed, err := gabs.ParseJSON([]byte(terraform.OutputJson(t, opts, "")))
	if err != nil {
		panic(err)
	}

	// Image
	image, _ := jsonParsed.JSONPointer("/image/value")
	assert.Equal(t, "nginx:latest", image.Data().(string))

	// Container name
	container_name, _ := jsonParsed.JSONPointer("/container_name/value")
	assert.Equal(t, "terraform-test", container_name.Data().(string))

	// Hostname
	hostname, _ := jsonParsed.JSONPointer("/hostname/value")
	assert.Equal(t, "terratest", hostname.Data().(string))

	// Working Dir
	working_dir, _ := jsonParsed.JSONPointer("/working_dir/value")
	assert.Equal(t, "/tmp", working_dir.Data().(string))

	// Restart policy
	restart_policy, _ := jsonParsed.JSONPointer("/restart/value")
	assert.Equal(t, "unless-stopped", restart_policy.Data().(string))

	// Privileged
	privileged, _ := jsonParsed.JSONPointer("/privileged/value")
	assert.Equal(t, false, privileged.Data().(bool))

	// Network mode
	network_mode, _ := jsonParsed.JSONPointer("/network_mode/value")
	assert.Equal(t, "bridge", network_mode.Data().(string))

	// DNS
	dns1, _ := jsonParsed.JSONPointer("/dns/value/0")
	assert.Equal(t, "1.0.0.1", dns1.Data().(string))

	dns2, _ := jsonParsed.JSONPointer("/dns/value/1")
	assert.Equal(t, "1.1.1.1", dns2.Data().(string))

	// Entrypoint
	entrypoint, _ := jsonParsed.JSONPointer("/entrypoint/value/0")
	assert.Equal(t, "/docker-entrypoint.sh", entrypoint.Data().(string))

	// Command
	command1, _ := jsonParsed.JSONPointer("/command/value/0")
	assert.Equal(t, "nginx", command1.Data().(string))

	command2, _ := jsonParsed.JSONPointer("/command/value/1")
	assert.Equal(t, "-g", command2.Data().(string))

	command3, _ := jsonParsed.JSONPointer("/command/value/2")
	assert.Equal(t, "daemon off;", command3.Data().(string))

	// Ports
	ports_internal, _ := jsonParsed.JSONPointer("/ports/value/0/internal")
	assert.Equal(t, float64(80), ports_internal.Data().(float64))

	ports_external, _ := jsonParsed.JSONPointer("/ports/value/0/external")
	assert.Equal(t, float64(9999), ports_external.Data().(float64))

	ports_protocol, _ := jsonParsed.JSONPointer("/ports/value/0/protocol")
	assert.Equal(t, "tcp", ports_protocol.Data().(string))

	// Named volumes
	named_volumes1_name, _ := jsonParsed.JSONPointer("/volumes/value/0/volume_name")
	assert.Equal(t, "nginx_volume", named_volumes1_name.Data().(string))

	named_volumes1_container_path, _ := jsonParsed.JSONPointer("/volumes/value/0/container_path")
	assert.Equal(t, "/etc/nginx", named_volumes1_container_path.Data().(string))

	named_volumes1_host_path, _ := jsonParsed.JSONPointer("/volumes/value/0/host_path")
	assert.Equal(t, "", named_volumes1_host_path.Data().(string))

	docker_volumes1_name, _ := jsonParsed.JSONPointer("/docker_volumes/value")
	assert.Equal(t, true, docker_volumes1_name.Exists("nginx_volume"))

	named_volumes2_name, _ := jsonParsed.JSONPointer("/volumes/value/1/volume_name")
	assert.Equal(t, "mnt_volume", named_volumes2_name.Data().(string))

	named_volumes2_container_path, _ := jsonParsed.JSONPointer("/volumes/value/1/container_path")
	assert.Equal(t, "/mnt", named_volumes2_container_path.Data().(string))

	named_volumes2_host_path, _ := jsonParsed.JSONPointer("/volumes/value/1/host_path")
	assert.Equal(t, "", named_volumes2_host_path.Data().(string))

	docker_volumes2_name, _ := jsonParsed.JSONPointer("/docker_volumes/value")
	assert.Equal(t, true, docker_volumes2_name.Exists("mnt_volume"))

	// Host paths
	host_paths_container_path, _ := jsonParsed.JSONPointer("/volumes/value/2/container_path")
	assert.Equal(t, "/tmp", host_paths_container_path.Data().(string))

	host_paths_volume_name, _ := jsonParsed.JSONPointer("/volumes/value/2/volume_name")
	assert.Equal(t, "", host_paths_volume_name.Data().(string))

	// Devices
	devices_host_path, _ := jsonParsed.JSONPointer("/devices/value/0/host_path")
	assert.Equal(t, "/dev/null", devices_host_path.Data().(string))

	devices_container_path, _ := jsonParsed.JSONPointer("/devices/value/0/container_path")
	assert.Equal(t, "/dev/newnull", devices_container_path.Data().(string))

	// Capabilities
	capabilities, _ := jsonParsed.JSONPointer("/capabilities/value/0/add/0")
	assert.Equal(t, "NET_ADMIN", capabilities.Data().(string))

	// Networks advanced
	networks_advanced1_name, _ := jsonParsed.JSONPointer("/networks_advanced/value/0/name")
	assert.Equal(t, "nginx_network", networks_advanced1_name.Data().(string))

	networks_advanced1_ipv4_address, _ := jsonParsed.JSONPointer("/networks_advanced/value/0/ipv4_address")
	assert.Equal(t, "10.0.30.2", networks_advanced1_ipv4_address.Data().(string))

	networks_advanced2_name, _ := jsonParsed.JSONPointer("/networks_advanced/value/1/name")
	assert.Equal(t, "app_network", networks_advanced2_name.Data().(string))

	networks_advanced2_ipv4_address, _ := jsonParsed.JSONPointer("/networks_advanced/value/1/ipv4_address")
	assert.Equal(t, "10.0.55.2", networks_advanced2_ipv4_address.Data().(string))

	// Docker networks
	docker_networks1_name, _ := jsonParsed.JSONPointer("/docker_networks/value")
	assert.Equal(t, true, docker_networks1_name.Exists("nginx_network"))

	docker_networks1_gateway, _ := jsonParsed.JSONPointer("/docker_networks/value/nginx_network/ipam_config/0/gateway")
	assert.Equal(t, "10.0.30.1", docker_networks1_gateway.Data().(string))

	docker_networks2_name, _ := jsonParsed.JSONPointer("/docker_networks/value")
	assert.Equal(t, true, docker_networks2_name.Exists("app_network"))

	docker_networks2_gateway, _ := jsonParsed.JSONPointer("/docker_networks/value/app_network/ipam_config/0/gateway")
	assert.Equal(t, "10.0.55.1", docker_networks2_gateway.Data().(string))

	// Environment Validation
	environment_env, _ := jsonParsed.JSONPointer("/environment/value/0")
	environment_name, _ := jsonParsed.JSONPointer("/environment/value/1")
	assert.Equal(t, "ENV=test", environment_env.Data().(string))
	assert.Equal(t, "NAME=nginx", environment_name.Data().(string))
}
