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
| <a name="input_capabilities"></a> [capabilities](#input\_capabilities) | n/a | `map(any)` | `{}` | no |
| <a name="input_command"></a> [command](#input\_command) | n/a | `list(string)` | `null` | no |
| <a name="input_container_name"></a> [container\_name](#input\_container\_name) | n/a | `string` | n/a | yes |
| <a name="input_devices"></a> [devices](#input\_devices) | n/a | `map(any)` | `{}` | no |
| <a name="input_dns"></a> [dns](#input\_dns) | n/a | `list(string)` | `null` | no |
| <a name="input_docker_networks"></a> [docker\_networks](#input\_docker\_networks) | n/a | `list(any)` | `[]` | no |
| <a name="input_entrypoint"></a> [entrypoint](#input\_entrypoint) | n/a | `list(string)` | `null` | no |
| <a name="input_environment"></a> [environment](#input\_environment) | n/a | `map(string)` | `null` | no |
| <a name="input_host_paths"></a> [host\_paths](#input\_host\_paths) | n/a | `map(any)` | `{}` | no |
| <a name="input_hostname"></a> [hostname](#input\_hostname) | n/a | `string` | `null` | no |
| <a name="input_image"></a> [image](#input\_image) | n/a | `string` | n/a | yes |
| <a name="input_named_volumes"></a> [named\_volumes](#input\_named\_volumes) | n/a | `map(any)` | `{}` | no |
| <a name="input_network_mode"></a> [network\_mode](#input\_network\_mode) | n/a | `string` | `null` | no |
| <a name="input_networks_advanced"></a> [networks\_advanced](#input\_networks\_advanced) | n/a | `list(any)` | `[]` | no |
| <a name="input_ports"></a> [ports](#input\_ports) | n/a | `list(any)` | `null` | no |
| <a name="input_privileged"></a> [privileged](#input\_privileged) | n/a | `bool` | `false` | no |
| <a name="input_restart_policy"></a> [restart\_policy](#input\_restart\_policy) | n/a | `string` | `null` | no |
| <a name="input_working_dir"></a> [working\_dir](#input\_working\_dir) | n/a | `string` | `null` | no |

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
| <a name="output_image"></a> [image](#output\_image) | n/a |
| <a name="output_network_mode"></a> [network\_mode](#output\_network\_mode) | n/a |
| <a name="output_networks_advanced"></a> [networks\_advanced](#output\_networks\_advanced) | n/a |
| <a name="output_ports"></a> [ports](#output\_ports) | n/a |
| <a name="output_privileged"></a> [privileged](#output\_privileged) | n/a |
| <a name="output_restart"></a> [restart](#output\_restart) | n/a |
| <a name="output_volumes"></a> [volumes](#output\_volumes) | n/a |
| <a name="output_working_dir"></a> [working\_dir](#output\_working\_dir) | n/a |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
