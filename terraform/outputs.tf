output "droplet_ipv4_address" {
  value       = digitalocean_droplet.web.ipv4_address
  description = "The public IP address of the droplet."

}
