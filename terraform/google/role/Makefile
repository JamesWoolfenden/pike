clean:
	-rm -fr .terraform
	-rm .terraform.lock.hcl
	-rm terraform.tfstate
	-rm terraform.tfstate.backup
apply: init
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu apply -auto-approve

plan: init
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu plan

destroy: init
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu destroy -auto-approve

init:
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu init

role: FORCE
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu apply -auto-approve

role_init:
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu  init

update:
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu init --upgrade

upgrade:update

FORCE:
