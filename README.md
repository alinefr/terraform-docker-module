# Terraform Docker Module

## Overview

Terraform docker module is a module to help docker maintenance over terraform. 
It should replace other means of docker maintenance like docker-compose.

There are several advantages of maintaining docker on terraform.

* Infrastructure as code. You don't need to manually ssh into servers
* CI/CD. Many CI tools offers some way to automate terraform execution.
* Remote execution. You don't need to manually ssh into servers.

This module uses under the hood [Docker Provider](https://www.terraform.io/docs/providers/docker/index.html).

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
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
| capabilities | Add or drop container capabilities | <pre>object({<br>    add  = list(string)<br>    drop = list(string)<br>  })</pre> | `null` | no |
| command | Override the default command | `list(string)` | `null` | no |
| devices | Device mappings | <pre>list(object({<br>    host_path      = string<br>    container_path = string<br>    permissions    = string<br>  }))</pre> | `null` | no |
| dns | Set custom dns servers for the container | `list(string)` | `null` | no |
| docker\_networks | List of custom networks to create | <pre>list(object({<br>    name = string<br>    ipam_config = object({<br>      aux_address = map(string)<br>      gateway     = string<br>      subnet      = string<br>    })<br>  }))</pre> | `null` | no |
| docker\_volumes | List of named volumes to create | `list(string)` | `null` | no |
| environment | Add environment variables | `list(string)` | `null` | no |
| healthcheck | Test to check if container is healthy | <pre>object({<br>    interval     = string<br>    retries      = number<br>    start_period = string<br>    test         = list(string)<br>    timeout      = string<br>  })</pre> | `null` | no |
| image | Specify the image to start the container from. Can either be a repository/tag or a partial image ID | `string` | n/a | yes |
| name | Custom container name | `string` | `null` | no |
| network\_mode | Specify a custom network mode | `string` | `null` | no |
| networks\_advanced | Advanced network options for the container | <pre>object({<br>    name         = string<br>    aliases      = string<br>    ipv4_address = string<br>    ipv6_address = string<br>  })</pre> | `null` | no |
| ports | Expose ports | <pre>list(object({<br>    internal = number<br>    external = number<br>    protocol = string<br>  }))</pre> | `null` | no |
| privileged | Give extended privileges to this container | `bool` | `false` | no |
| restart\_policy | Restart policy. Default: no | `string` | `"no"` | no |
| volumes | Mount host paths or named volumes, specified as sub-options to a service | <pre>list(object({<br>    volume_name    = string<br>    container_path = string<br>    host_path      = string<br>    read_only      = bool<br>  }))</pre> | `null` | no |
| working\_dir | Working directory inside the container | `string` | `null` | no |

## Outputs

No output.

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
