<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

No requirements.

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_docker"></a> [docker](#module\_docker) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_capabilities"></a> [capabilities](#input\_capabilities) | Add or drop container capabilities | <pre>object({<br>    add  = list(string)<br>    drop = list(string)<br>  })</pre> | `null` | no |
| <a name="input_command"></a> [command](#input\_command) | Override the default command | `list(string)` | `null` | no |
| <a name="input_container_name"></a> [container\_name](#input\_container\_name) | Custom container name | `string` | `null` | no |
| <a name="input_devices"></a> [devices](#input\_devices) | Device mappings | <pre>map(object({<br>    container_path = string<br>    permissions    = string<br>  }))</pre> | `{}` | no |
| <a name="input_dns"></a> [dns](#input\_dns) | Set custom dns servers for the container | `list(string)` | `null` | no |
| <a name="input_docker_networks"></a> [docker\_networks](#input\_docker\_networks) | List of custom networks to create<pre>hcl<br>docker_networks = [<br>  {<br>    name = "proxy-tier"<br>    ipam_config = {<br>      aux_address = {}<br>      gateway     = "10.0.0.1"<br>      subnet      = "10.0.0.0/24"<br>    }<br>  }<br>]</pre> | `any` | `[]` | no |
| <a name="input_entrypoint"></a> [entrypoint](#input\_entrypoint) | Override the default entrypoint | `list(string)` | `null` | no |
| <a name="input_environment"></a> [environment](#input\_environment) | Add environment variables | `map(string)` | `null` | no |
| <a name="input_existing_image"></a> [existing\_image](#input\_existing\_image) | Specify an existing image from another module | `string` | `null` | no |
| <a name="input_healthcheck"></a> [healthcheck](#input\_healthcheck) | Test to check if container is healthy | <pre>object({<br>    interval     = string<br>    retries      = number<br>    start_period = string<br>    test         = list(string)<br>    timeout      = string<br>  })</pre> | `null` | no |
| <a name="input_host_paths"></a> [host\_paths](#input\_host\_paths) | Mount host paths | <pre>map(object({<br>    container_path = string<br>    read_only      = bool<br>  }))</pre> | `{}` | no |
| <a name="input_hostname"></a> [hostname](#input\_hostname) | Set docker hostname | `string` | `null` | no |
| <a name="input_image"></a> [image](#input\_image) | Specify the image to start the container from. Can either be a repository/tag or a partial image ID | `string` | `null` | no |
| <a name="input_init"></a> [init](#input\_init) | If init process should be used as the PID 1 in the container | `bool` | `false` | no |
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
| <a name="output_capabilities"></a> [capabilities](#output\_capabilities) | n/a |
| <a name="output_command"></a> [command](#output\_command) | n/a |
| <a name="output_container_name"></a> [container\_name](#output\_container\_name) | n/a |
| <a name="output_devices"></a> [devices](#output\_devices) | n/a |
| <a name="output_dns"></a> [dns](#output\_dns) | n/a |
| <a name="output_docker_networks"></a> [docker\_networks](#output\_docker\_networks) | n/a |
| <a name="output_docker_volumes"></a> [docker\_volumes](#output\_docker\_volumes) | n/a |
| <a name="output_entrypoint"></a> [entrypoint](#output\_entrypoint) | n/a |
| <a name="output_environment"></a> [environment](#output\_environment) | n/a |
| <a name="output_healthcheck"></a> [healthcheck](#output\_healthcheck) | n/a |
| <a name="output_hostname"></a> [hostname](#output\_hostname) | n/a |
| <a name="output_image_name"></a> [image\_name](#output\_image\_name) | n/a |
| <a name="output_network_mode"></a> [network\_mode](#output\_network\_mode) | n/a |
| <a name="output_networks_advanced"></a> [networks\_advanced](#output\_networks\_advanced) | n/a |
| <a name="output_ports"></a> [ports](#output\_ports) | n/a |
| <a name="output_privileged"></a> [privileged](#output\_privileged) | n/a |
| <a name="output_restart"></a> [restart](#output\_restart) | n/a |
| <a name="output_volumes"></a> [volumes](#output\_volumes) | n/a |
| <a name="output_working_dir"></a> [working\_dir](#output\_working\_dir) | n/a |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
