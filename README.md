# Terraform Docker Module

## Overview

Terraform docker module is a module to help docker maintenance over terraform. 
It should replace other means of docker maintenance like docker-compose.

There are several advantages of maintaining docker on terraform.

* Infrastructure as code. You don't need to manually ssh into servers
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
  version = "2.1.0"

  image = "masnagam/nginx-proxy:latest"
  name = "proxy"
  restart_policy = "always"
  docker_networks = [
    {
      name = "proxy-tier"
      ipam_config = {
        aux_address = {}
        gateway = "10.0.20.1"
        subnet = "10.0.20.0/24"
      }
    }
  ]
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
  volumes = [
    {
      volume_name = "nginx_confs"
      container_path = "/etc/nginx/conf.d"
      host_path = null
      read_only = false
      create = true
    },
    {
      volume_name = "nginx_html"
      container_path = "/var/www/html"
      host_path = null
      read_only = false
      create = true
    },
    {
      volume_name = null
      container_path = "/etc/nginx/certs"
      host_path = "/media/letsencrypt/etc/letsencrypt/live"
      read_only = false
      create = false
    },
    {
      volume_name = null
      container_path = "/etc/nginx/archive"
      host_path = "/media/letsencrypt/etc/letsencrypt/archive"
      read_only = false
      create = false
    },
    {
      volume_name = null
      container_path = "/tmp/docker.sock"
      host_path = "/var/run/docker.sock"
      read_only = true
      create = true
    }
  ]
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
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.13 |
| docker | ~> 2.7 |

## Providers

| Name | Version |
|------|---------|
| docker | ~> 2.7 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| capabilities | Add or drop container capabilities | <pre>object({<br>    add  = list(string)<br>    drop = list(string)<br>  })</pre> | `null` | no |
| command | Override the default command | `list(string)` | `null` | no |
| devices | Device mappings | <pre>list(object({<br>    host_path      = string<br>    container_path = string<br>    permissions    = string<br>  }))</pre> | `null` | no |
| dns | Set custom dns servers for the container | `list(string)` | `null` | no |
| docker\_networks | List of custom networks to create | <pre>list(object({<br>    name = string<br>    ipam_config = object({<br>      aux_address = map(string)<br>      gateway     = string<br>      subnet      = string<br>    })<br>  }))</pre> | `null` | no |
| environment | Add environment variables | `list(string)` | `null` | no |
| healthcheck | Test to check if container is healthy | <pre>object({<br>    interval     = string<br>    retries      = number<br>    start_period = string<br>    test         = list(string)<br>    timeout      = string<br>  })</pre> | `null` | no |
| image | Specify the image to start the container from. Can either be a repository/tag or a partial image ID | `string` | n/a | yes |
| name | Custom container name | `string` | `null` | no |
| network\_mode | Specify a custom network mode | `string` | `null` | no |
| networks\_advanced | Advanced network options for the container | <pre>object({<br>    name         = string<br>    aliases      = string<br>    ipv4_address = string<br>    ipv6_address = string<br>  })</pre> | `null` | no |
| ports | Expose ports | <pre>list(object({<br>    internal = number<br>    external = number<br>    protocol = string<br>  }))</pre> | `null` | no |
| privileged | Give extended privileges to this container | `bool` | `false` | no |
| restart\_policy | Restart policy. Default: no | `string` | `"no"` | no |
| volumes | Mount host paths or named volumes, specified as sub-options to a service | <pre>list(object({<br>    volume_name    = string<br>    container_path = string<br>    host_path      = string<br>    read_only      = bool<br>    create         = bool<br>  }))</pre> | `null` | no |
| working\_dir | Working directory inside the container | `string` | `null` | no |

## Outputs

No output.

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
