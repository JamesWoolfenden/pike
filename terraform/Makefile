clean:
	-rm -fr .terraform
	-rm .terraform.lock.hcl
	-rm terraform.tfstate
	-rm terraform.tfstate.backup
apply: init
	terraform apply -auto-approve

init:
	terraform init

role: FORCE
	terraform  -chdir=./role apply -auto-approve

FORCE:
