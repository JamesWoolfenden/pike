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

show: init
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu show

init:
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu init

upgrade:
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu init --upgrade

rm:
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu state rm $(RESOURCE)

update: upgrade


role: FORCE
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu  -chdir=./role apply -auto-approve

role_init:
	GOOGLE_BACKEND_CREDENTIALS=$(PIKE_GOOGLE_BACKEND_CREDENTIALS) tofu  -chdir=./role init

scan:
	GCP_PROJECT=pike-477416 pike scan -w -i -e -d ./ -outfile "./role/google_project_iam_custom_role.tf" -policyName terraform_pike

FORCE:
