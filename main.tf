locals {
  image          = var.image
  img_nouser     = replace(local.image, "/", "") != local.image ? split("/", local.image)[1] : split("/", local.image)[0]
  container_name = split(":", local.img_nouser)[0]
}

data "docker_registry_image" "default" {
  name = var.image
}

resource "docker_image" "default" {
  name          = data.docker_registry_image.default.name
  pull_triggers = [data.docker_registry_image.default.sha256_digest]
}

resource "docker_volume" "default" {
  for_each = var.named_volumes
  name     = each.key
}

resource "docker_network" "default" {
  for_each = var.docker_networks
  name     = each.key

  ipam_config {
    aux_address = each.value.ipam_config.aux_address
    gateway     = each.value.ipam_config.gateway
    subnet      = each.value.ipam_config.subnet
  }
}


resource "docker_container" "default" {
  name         = var.container_name != null ? var.container_name : local.container_name
  image        = docker_image.default.image_id
  hostname     = var.hostname
  restart      = var.restart_policy
  privileged   = var.privileged
  network_mode = var.network_mode
  working_dir  = var.working_dir
  dns          = var.dns
  command      = var.command
  env          = var.environment != null ? [for k, v in var.environment : "${k}=${v}"] : null

  dynamic "ports" {
    for_each = var.ports == null ? [] : var.ports
    content {
      internal = ports.value.internal
      external = ports.value.external
      protocol = ports.value.protocol
    }
  }

  dynamic "volumes" {
    for_each = var.named_volumes
    content {
      volume_name    = volumes.key
      container_path = volumes.value.container_path
      read_only      = volumes.value.read_only
    }
  }

  dynamic "volumes" {
    for_each = var.host_paths
    content {
      host_path      = volumes.key
      container_path = volumes.value.container_path
      read_only      = volumes.value.read_only
    }
  }

  dynamic "volumes" {
    for_each = var.volumes_from_containers == null ? [] : var.volumes_from_containers
    content {
      from_container = volumes.value
    }
  }

  dynamic "devices" {
    for_each = var.devices
    content {
      host_path      = devices.key
      container_path = devices.value.container_path
      permissions    = devices.value.permissions
    }
  }

  dynamic "capabilities" {
    for_each = var.capabilities == null ? [] : [var.capabilities]
    content {
      add  = var.capabilities.add
      drop = var.capabilities.drop
    }
  }

  dynamic "networks_advanced" {
    for_each = var.networks_advanced == null ? [] : [var.networks_advanced]
    content {
      name         = var.networks_advanced.name
      ipv4_address = var.networks_advanced.ipv4_address
      ipv6_address = var.networks_advanced.ipv6_address
      aliases      = var.networks_advanced.aliases
    }
  }

  dynamic "healthcheck" {
    for_each = var.healthcheck == null ? [] : [var.healthcheck]
    content {
      interval     = var.healthcheck.interval
      retries      = var.healthcheck.retries
      start_period = var.healthcheck.start_period
      test         = var.healthcheck.test
      timeout      = var.healthcheck.timeout
    }
  }
}
