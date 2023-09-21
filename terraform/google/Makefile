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
	GOOGLE_CLOUD_KEYFILE_JSON="" terraform  -chdir=./role apply -auto-approve

FORCE:
