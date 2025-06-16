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

upgrade:
	tofu init --upgrade

role: FORCE
	GOOGLE_CLOUD_KEYFILE_JSON="" tofu  -chdir=./role apply -auto-approve

role_init:
	GOOGLE_CLOUD_KEYFILE_JSON="" tofu  -chdir=./role init

FORCE:
