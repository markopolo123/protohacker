# Terraform 4 protohackers

Just enough Terraform to create a public endpoint for testing protohackers code

<!-- BEGIN_TF_DOCS -->

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_digitalocean"></a> [digitalocean](#requirement\_digitalocean) | ~> 2.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_digitalocean"></a> [digitalocean](#provider\_digitalocean) | 2.22.3 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [digitalocean_droplet.web](https://registry.terraform.io/providers/digitalocean/digitalocean/latest/docs/resources/droplet) | resource |
| [digitalocean_ssh_key.my_key](https://registry.terraform.io/providers/digitalocean/digitalocean/latest/docs/data-sources/ssh_key) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_do_droplet_size"></a> [do\_droplet\_size](#input\_do\_droplet\_size) | slug for droplet size | `string` | `"s-1vcpu-512mb-10gb"` | no |
| <a name="input_do_ssh_key"></a> [do\_ssh\_key](#input\_do\_ssh\_key) | name of your DO ssh key, as it appears in the web console | `string` | `""` | no |
| <a name="input_do_token"></a> [do\_token](#input\_do\_token) | DO API token, scoped with write access | `string` | `""` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_droplet_ipv4_address"></a> [droplet\_ipv4\_address](#output\_droplet\_ipv4\_address) | The public IP address of the droplet. |
<!-- END_TF_DOCS -->
