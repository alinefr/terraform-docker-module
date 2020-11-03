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
| capabilities | n/a | <pre>object({<br>    add  = list(string)<br>    drop = list(string)<br>  })</pre> | `null` | no |
| command | n/a | `list(string)` | `null` | no |
| devices | n/a | <pre>list(object({<br>    host_path      = string<br>    container_path = string<br>    permissions    = string<br>  }))</pre> | `null` | no |
| dns | n/a | `list(string)` | `null` | no |
| docker\_image | n/a | `string` | n/a | yes |
| docker\_name | n/a | `string` | n/a | yes |
| docker\_networks | n/a | <pre>list(object({<br>    name = string<br>    ipam_config = object({<br>      aux_address = map(string)<br>      gateway     = string<br>      subnet      = string<br>    })<br>  }))</pre> | `null` | no |
| docker\_volumes | n/a | `list(string)` | `null` | no |
| env | n/a | `list(string)` | `null` | no |
| healthcheck | n/a | <pre>object({<br>    interval     = string<br>    retries      = number<br>    start_period = string<br>    test         = list(string)<br>    timeout      = string<br>  })</pre> | `null` | no |
| network\_mode | n/a | `string` | `null` | no |
| networks\_advanced | n/a | <pre>object({<br>    name         = string<br>    ipv4_address = string<br>  })</pre> | `null` | no |
| ports | n/a | <pre>list(object({<br>    internal = number<br>    external = number<br>    protocol = string<br>  }))</pre> | `null` | no |
| privileged | n/a | `bool` | `false` | no |
| restart | n/a | `string` | `"unless-stopped"` | no |
| volumes | n/a | <pre>list(object({<br>    volume_name    = string<br>    container_path = string<br>    host_path      = string<br>    read_only      = bool<br>  }))</pre> | `null` | no |
| working\_dir | n/a | `string` | `null` | no |

## Outputs

No output.

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
