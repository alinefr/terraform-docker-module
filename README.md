# Terraform Docker Module

## Overview

Terraform docker module is a module to help docker maintenance over terraform. 
It should replace other means of docker maintenance like docker-compose.

There are several advantages of maintaining docker on terraform.

* Infrastructure as code.
* CI/CD. Many CI tools offers some way to automate terraform execution.
* Remote execution. You don't need to manually ssh into servers.

This module uses under the hood [Docker Provider](https://www.terraform.io/docs/providers/docker/index.html).

Example:

```hcl
provider "docker" {
  host = "tcp://192.168.0.100:2375/"
}

module "proxy" {
  source = "alinefr/module/docker"
  version = "<add latest version>"

  image = "masnagam/nginx-proxy:latest"
  container_name = "proxy"
  restart_policy = "always"
  docker_networks = {
    "proxy-tier" = {
      ipam_config = {
        aux_address = {}
        gateway = "10.0.20.1"
        subnet = "10.0.20.0/24"
      }
    }
  }
  ports = [
    {
      internal = 80
      external = 80
      protocol = "tcp"
    },
    {
      internal = 443
      external = 443
      protocol = "tcp"
    }
  ]
  named_volumes = {
    "nginx_confs" = {
      container_path = "/etc/nginx/conf.d"
      read_only = false
      create = true
    },
    "nginx_html" = {
      container_path = "/var/www/html"
      read_only = false
      create = true
    }
  }
  host_paths = {
    "/media/letsencrypt/etc/letsencrypt/live" = {
      container_path = "/etc/nginx/certs"
      read_only = false
    },
    "/media/letsencrypt/etc/letsencrypt/archive" = {
      container_path = "/etc/nginx/archive"
      read_only = false
    },
    "/var/run/docker.sock" = {
      container_path = "/tmp/docker.sock"
      read_only = true
    }
  }
  capabilities = {
    add = ["NET_ADMIN"]
    drop = []
  }
  networks_advanced = {
    name = "proxy-tier"
    ipv4_address = "10.0.20.100"
    ipv6_address = null
    aliases = null
  }
}

module "letsencrypt-companion" {
  source = "alinefr/module/docker"
  version = "<add latest version>"

  image = "jrcs/letsencrypt-nginx-proxy-companion"
  container_name = "letsencrypt-companion"
  restart_policy = "always"
  volumes_from_containers = [
      "proxy"
  ]
  host_paths = {
    "/var/run/docker.sock" = {
      container_path = "/var/run/docker.sock"
      read_only = true
    }
  }
  networks_advanced = {
    name = "proxy-tier"
    ipv4_address = "10.0.20.101"
    ipv6_address = null
    aliases = null
  }
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.4 |
| <a name="requirement_docker"></a> [docker](#requirement\_docker) | ~> 3.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_docker"></a> [docker](#provider\_docker) | ~> 3.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [docker_container.default](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/container) | resource |
| [docker_image.default](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/image) | resource |
| [docker_network.default](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/network) | resource |
| [docker_volume.default](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/volume) | resource |
| [docker_registry_image.default](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/data-sources/registry_image) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_capabilities"></a> [capabilities](#input\_capabilities) | Add or drop container capabilities | <pre>object({<br>    add  = list(string)<br>    drop = list(string)<br>  })</pre> | `null` | no |
| <a name="input_command"></a> [command](#input\_command) | Override the default command | `list(string)` | `null` | no |
| <a name="input_container_name"></a> [container\_name](#input\_container\_name) | Custom container name | `string` | `null` | no |
| <a name="input_devices"></a> [devices](#input\_devices) | Device mappings | <pre>map(object({<br>    container_path = string<br>    permissions    = string<br>  }))</pre> | `{}` | no |
| <a name="input_dns"></a> [dns](#input\_dns) | Set custom dns servers for the container | `list(string)` | `null` | no |
| <a name="input_docker_networks"></a> [docker\_networks](#input\_docker\_networks) | List of custom networks to create<pre>hcl<br>docker_networks = [<br>  {<br>    name = "proxy-tier"<br>    ipam_config = {<br>      aux_address = {}<br>      gateway     = "10.0.0.1"<br>      subnet      = "10.0.0.0/24"<br>    }<br>  }<br>]</pre> | `any` | `null` | no |
| <a name="input_entrypoint"></a> [entrypoint](#input\_entrypoint) | Override the default entrypoint | `list(string)` | `null` | no |
| <a name="input_environment"></a> [environment](#input\_environment) | Add environment variables | `map(string)` | `null` | no |
| <a name="input_healthcheck"></a> [healthcheck](#input\_healthcheck) | Test to check if container is healthy | <pre>object({<br>    interval     = string<br>    retries      = number<br>    start_period = string<br>    test         = list(string)<br>    timeout      = string<br>  })</pre> | `null` | no |
| <a name="input_host_paths"></a> [host\_paths](#input\_host\_paths) | Mount host paths | <pre>map(object({<br>    container_path = string<br>    read_only      = bool<br>  }))</pre> | `{}` | no |
| <a name="input_hostname"></a> [hostname](#input\_hostname) | Set docker hostname | `string` | `null` | no |
| <a name="input_image"></a> [image](#input\_image) | Specify the image to start the container from. Can either be a repository/tag or a partial image ID | `string` | n/a | yes |
| <a name="input_named_volumes"></a> [named\_volumes](#input\_named\_volumes) | Mount named volumes | <pre>map(object({<br>    container_path = string<br>    read_only      = bool<br>    create         = bool<br>  }))</pre> | `{}` | no |
| <a name="input_network_mode"></a> [network\_mode](#input\_network\_mode) | Specify a custom network mode | `string` | `null` | no |
| <a name="input_networks_advanced"></a> [networks\_advanced](#input\_networks\_advanced) | Advanced network options for the container<pre>hcl<br>networks_advanced = [<br>  {<br>    name         = "proxy-tier"<br>    ipv4_address = "10.0.0.14"<br>  },<br>  {<br>    name         = "media-tier"<br>    ipv4_address = "172.0.0.14"<br>  }<br>]</pre> | `any` | `null` | no |
| <a name="input_ports"></a> [ports](#input\_ports) | Expose ports | <pre>list(object({<br>    internal = number<br>    external = number<br>    protocol = string<br>  }))</pre> | `null` | no |
| <a name="input_privileged"></a> [privileged](#input\_privileged) | Give extended privileges to this container | `bool` | `false` | no |
| <a name="input_restart_policy"></a> [restart\_policy](#input\_restart\_policy) | Restart policy. Default: no | `string` | `"no"` | no |
| <a name="input_volumes_from_containers"></a> [volumes\_from\_containers](#input\_volumes\_from\_containers) | Mount volumes from another container | `list(any)` | `null` | no |
| <a name="input_working_dir"></a> [working\_dir](#input\_working\_dir) | Working directory inside the container | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_devices"></a> [devices](#output\_devices) | n/a |
| <a name="output_environment"></a> [environment](#output\_environment) | n/a |
| <a name="output_volumes"></a> [volumes](#output\_volumes) | n/a |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
