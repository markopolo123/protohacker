data "digitalocean_ssh_key" "my_key" {
  name = var.do_ssh_key
}

# Create a new Web Droplet in the amsterdam region
resource "digitalocean_droplet" "web" {
  image    = "ubuntu-22-04-x64"
  name     = "protohacker-test"
  region   = "ams3"
  ssh_keys = tolist([data.digitalocean_ssh_key.my_key.id])
  size     = var.do_droplet_size
}
