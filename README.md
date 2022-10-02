# Protohacker solutions in Golang

Solutions for [Protohackers](https://protohackers.com) in Golang, with a
sprinkling of Terraform.

## Prerequisites

### Packages

* Golang (REQUIRED) - language used to solve protohackers problems
* Taskfile (optional) - used as a taskrunner
* Terraform (optional) - used to create public endpoint for testing code
* jq (optional) - used to get instance slugs from DO API and some 1password bits
* direnv (optional) - helpful tool for managing env variables

Using homebrew they may be installed as follows:

```bash
brew install go-task/tap/go-task
brew install terraform jq golang terraform-docs direnv
```

## Creating public endpoint for testing code

The following task will create a droplet and drop you in a prompt:

```bash
task terraform:ssh
```

In order for this Terraform to work *as is* you will need to have configured a Digital
Ocean account and remote storage using the [S3 Terraform backend](https://www.terraform.io/language/settings/backends/s3)

### Credentials

* [Create a Digital Ocean API Key](https://docs.digitalocean.com/reference/api/create-personal-access-token/)
* [Add an SSH key to your Digital Ocean account](https://docs.digitalocean.com/products/droplets/how-to/add-ssh-keys/to-team/)

### Terraform Backend

This code is configured to use my [minio](https://min.io) installation, however
this could easily be reconfigured to use AWS S3 or [Digital Ocean's spaces](https://dev.to/jmarhee/digitalocean-spaces-as-a-terraform-backend-3lck).

### Managing Secrets

I use [direnv](https://direnv.net) to manage the environment variables:

> Note I'm using `1password` to manage my secrets, so my [.envrc](.envrc) file may be commited
> safely to this repository

> If you do not want to use `direnv` you may export these variables directly in
> your shell

My `.envrc` file in the root dir for this project is as follows:

```bash
cat .envrc

export AWS_S3_ENDPOINT="https://minio.internal.tld"
export AWS_ACCESS_KEY_ID=$(op item get adf3123 --fields access-key)
export AWS_SECRET_ACCESS_KEY=$(op item get qwerty1234 --fields secret-key)
export TF_VAR_do_token=$(op item get 12345qwerty --fields credential)
export TF_VAR_do_ssh_key="$(op item get 12345qwerty --fields name"
```

GOOS=linux GOARCH=amd64 go build

```bash
curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer $TF_VAR_do_token" "https://api.digitalocean.com/v2/sizes" | jq '.sizes |.[] |.slug'
```
