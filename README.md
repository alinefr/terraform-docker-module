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
| container\_name | Custom container name | `string` | `null` | no |
| devices | Device mappings | <pre>map(object({<br>    container_path = string<br>    permissions    = string<br>  }))</pre> | `{}` | no |
| dns | Set custom dns servers for the container | `list(string)` | `null` | no |
| docker\_networks | List of custom networks to create | <pre>map(object({<br>    ipam_config = object({<br>      aux_address = map(string)<br>      gateway     = string<br>      subnet      = string<br>    })<br>  }))</pre> | `{}` | no |
| environment | Add environment variables | `list(string)` | `null` | no |
| healthcheck | Test to check if container is healthy | <pre>object({<br>    interval     = string<br>    retries      = number<br>    start_period = string<br>    test         = list(string)<br>    timeout      = string<br>  })</pre> | `null` | no |
| host\_paths | Mount host paths | <pre>map(object({<br>    container_path = string<br>    read_only      = bool<br>  }))</pre> | `{}` | no |
| hostname | Set docker hostname | `string` | `null` | no |
| image | Specify the image to start the container from. Can either be a repository/tag or a partial image ID | `string` | n/a | yes |
| named\_volumes | Mount named volumes | <pre>map(object({<br>    container_path = string<br>    read_only      = bool<br>    create         = bool<br>  }))</pre> | `{}` | no |
| network\_mode | Specify a custom network mode | `string` | `null` | no |
| networks\_advanced | Advanced network options for the container | <pre>object({<br>    name         = string<br>    aliases      = list(string)<br>    ipv4_address = string<br>    ipv6_address = string<br>  })</pre> | `null` | no |
| ports | Expose ports | <pre>list(object({<br>    internal = number<br>    external = number<br>    protocol = string<br>  }))</pre> | `null` | no |
| privileged | Give extended privileges to this container | `bool` | `false` | no |
| restart\_policy | Restart policy. Default: no | `string` | `"no"` | no |
| volumes\_from\_containers | Mount volumes from another container | `list` | `null` | no |
| working\_dir | Working directory inside the container | `string` | `null` | no |

## Outputs

No output.

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
