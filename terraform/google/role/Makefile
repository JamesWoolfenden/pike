clean:
	-rm -fr .terraform
	-rm .terraform.lock.hcl
	-rm terraform.tfstate
	-rm terraform.tfstate.backup
apply: init
	tofu apply -auto-approve

plan: init
	tofu plan

destroy: init
	tofu destroy -auto-approve

init:
	tofu init

role: FORCE
	GOOGLE_CLOUD_KEYFILE_JSON="" terraform  -chdir=apply -auto-approve

role_init:
	GOOGLE_CLOUD_KEYFILE_JSON="" terraform  -chdir= init

FORCE:
