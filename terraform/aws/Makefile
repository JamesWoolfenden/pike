clean:
	-rm -fr .terraform
	-rm .terraform.lock.hcl
	-rm terraform.tfstate
	-rm terraform.tfstate.backup
apply: init
	terraform apply -auto-approve

plan: init
	terraform plan

destroy: init
	terraform destroy -auto-approve

init:
	terraform init

role: FORCE
	terraform  -chdir=./role apply -auto-approve

FORCE:

sleep:
	timeout 5
up: role sleep apply

mac: role init
	sleep 5
	terraform apply -auto-approve
