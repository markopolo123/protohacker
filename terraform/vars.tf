# Set the variable value in *.tfvars file
# or using -var="do_token=..." CLI option
variable "do_token" {
  type        = string
  default     = ""
  description = "DO API token, scoped with write access"
}

variable "do_ssh_key" {
  type        = string
  default     = ""
  description = "name of your DO ssh key, as it appears in the web console"
}

variable "do_droplet_size" {
  type        = string
  default     = "s-1vcpu-512mb-10gb"
  description = "slug for droplet size"
}
