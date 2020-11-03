data "docker_registry_image" "default" {
  name = var.docker_image
}

resource "docker_image" "default" {
  name          = data.docker_registry_image.default.name
  pull_triggers = [data.docker_registry_image.default.sha256_digest]
}

resource "docker_volume" "default" {
  count = var.docker_volumes != null ? length(var.docker_volumes) : 0
  name  = var.docker_volumes[count.index]
}

resource "docker_network" "default" {
  count = var.docker_networks != null ? length(var.docker_networks) : 0
  name  = var.docker_networks[count.index].name

  ipam_config {
    aux_address = var.docker_networks[count.index].ipam_config.aux_address
    gateway     = var.docker_networks[count.index].ipam_config.gateway
    subnet      = var.docker_networks[count.index].ipam_config.subnet
  }
}


resource "docker_container" "default" {
  name         = var.docker_name
  image        = docker_image.default.latest
  restart      = var.restart
  privileged   = var.privileged
  network_mode = var.network_mode
  working_dir  = var.working_dir
  dns          = var.dns
  command      = var.command
  env          = var.env

  dynamic "ports" {
    for_each = var.ports == null ? [] : var.ports
    content {
      internal = ports.value.internal
      external = ports.value.external
      protocol = ports.value.protocol
    }
  }

  dynamic "volumes" {
    for_each = var.volumes == null ? [] : var.volumes
    content {
      volume_name    = volumes.value.volume_name
      host_path      = volumes.value.host_path
      container_path = volumes.value.container_path
      read_only      = volumes.value.read_only
    }
  }

  dynamic "devices" {
    for_each = var.devices == null ? [] : var.devices
    content {
      host_path      = devices.value.host_path
      container_path = devices.value.container_path
      permissions    = devices.value.permissions
    }
  }

  dynamic "capabilities" {
    for_each = var.capabilities == null ? [] : list(var.capabilities)
    content {
      add  = var.capabilities.add
      drop = var.capabilities.drop
    }
  }

  dynamic "networks_advanced" {
    for_each = var.networks_advanced == null ? [] : list(var.networks_advanced)
    content {
      name         = var.networks_advanced.name
      ipv4_address = var.networks_advanced.ipv4_address
    }
  }

  dynamic "healthcheck" {
    for_each = var.healthcheck == null ? [] : list(var.healthcheck)
    content {
      interval     = var.healthcheck.interval
      retries      = var.healthcheck.retries
      start_period = var.healthcheck.start_period
      test         = var.healthcheck.test
      timeout      = var.healthcheck.timeout
    }
  }
}
