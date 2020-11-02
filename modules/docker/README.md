# Terraform Docker Module

This module allows to maintain docker images, volumes, networks
& containers.

It replaces docker-compose for docker management.

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.13 |

## Providers

| Name | Version |
|------|---------|
| docker | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| capabilities | Additional capabilities | <pre>object({<br>    add  = list(string)<br>    drop = list(string)<br>  })</pre> | `null` | no |
| command | Custom command | `list(string)` | `null` | no |
| devices | Exported devices | <pre>list(object({<br>    host_path      = string<br>    container_path = string<br>    permissions    = string<br>  }))</pre> | `null` | no |
| dns | Custom dns for container | `list(string)` | `null` | no |
| docker\_image | Remote docker image | `string` | n/a | yes |
| docker\_name | Docker container name | `string` | n/a | yes |
| docker\_networks | Docker networks to create | <pre>list(object({<br>    name = string<br>    ipam_config = object({<br>      aux_address = map(string)<br>      gateway     = string<br>      subnet      = string<br>    })<br>  }))</pre> | `null` | no |
| docker\_volumes | Docker volumes to create | `list(string)` | `null` | no |
| env | Environment variables | `list(string)` | `null` | no |
| healthcheck | Healthcheck | <pre>object({<br>    interval     = string<br>    retries      = number<br>    start_period = string<br>    test         = list(string)<br>    timeout      = string<br>  })</pre> | `null` | no |
| network\_mode | Network mode | `string` | `null` | no |
| networks\_advanced | Custom network | <pre>object({<br>    name         = string<br>    ipv4_address = string<br>  })</pre> | `null` | no |
| ports | Exported ports | <pre>list(object({<br>    internal = number<br>    external = number<br>    protocol = string<br>  }))</pre> | `null` | no |
| privileged | If we need full privileges | `bool` | `false` | no |
| restart | Restart policy | `string` | `"unless-stopped"` | no |
| volumes | Mounted volumes | <pre>list(object({<br>    volume_name    = string<br>    container_path = string<br>    host_path      = string<br>    read_only      = bool<br>  }))</pre> | `null` | no |
| working\_dir | Working dir | `string` | `null` | no |

## Outputs

No output.
