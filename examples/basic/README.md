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
| <a name="input_container_name"></a> [container\_name](#input\_container\_name) | n/a | `string` | n/a | yes |
| <a name="input_devices"></a> [devices](#input\_devices) | n/a | `map(any)` | `{}` | no |
| <a name="input_environment"></a> [environment](#input\_environment) | n/a | `map(string)` | `null` | no |
| <a name="input_gateway"></a> [gateway](#input\_gateway) | n/a | `string` | n/a | yes |
| <a name="input_host_paths"></a> [host\_paths](#input\_host\_paths) | n/a | `map(any)` | `{}` | no |
| <a name="input_ipv4_address"></a> [ipv4\_address](#input\_ipv4\_address) | n/a | `string` | n/a | yes |
| <a name="input_named_volumes"></a> [named\_volumes](#input\_named\_volumes) | n/a | `map(any)` | `{}` | no |
| <a name="input_network_name"></a> [network\_name](#input\_network\_name) | n/a | `string` | n/a | yes |
| <a name="input_ports"></a> [ports](#input\_ports) | n/a | `list(any)` | `null` | no |
| <a name="input_subnet"></a> [subnet](#input\_subnet) | n/a | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_devices"></a> [devices](#output\_devices) | n/a |
| <a name="output_environment"></a> [environment](#output\_environment) | n/a |
| <a name="output_volumes"></a> [volumes](#output\_volumes) | n/a |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
