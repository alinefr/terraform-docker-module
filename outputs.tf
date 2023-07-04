output "image_name" {
  value = var.existing_image != null ? var.existing_image : docker_image.default[var.image].name
}

output "image_id" {
  value = var.existing_image != null ? var.existing_image : docker_image.default[var.image].image_id
}

output "container_name" {
  value = docker_container.default.name
}

output "hostname" {
  value = docker_container.default.hostname
}

output "working_dir" {
  value = docker_container.default.working_dir
}

output "restart" {
  value = docker_container.default.restart
}

output "privileged" {
  value = docker_container.default.privileged
}

output "network_mode" {
  value = docker_container.default.network_mode
}

output "dns" {
  value = docker_container.default.dns
}

output "entrypoint" {
  value = docker_container.default.entrypoint
}

output "command" {
  value = docker_container.default.command
}

output "ports" {
  value = docker_container.default.ports
}

output "volumes" {
  value = docker_container.default.volumes
}

output "docker_volumes" {
  value = docker_volume.default
}

output "devices" {
  value = docker_container.default.devices
}

output "capabilities" {
  value = docker_container.default.capabilities
}

output "networks_advanced" {
  value = docker_container.default.networks_advanced
}

output "docker_networks" {
  value = docker_network.default
}

output "healthcheck" {
  value = docker_container.default.healthcheck
}

output "environment" {
  value = docker_container.default.env
}
