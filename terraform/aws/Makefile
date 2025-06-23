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

upgrade:
	terraform init --upgrade

update: upgrade

role: FORCE
	terraform  -chdir=./role apply -auto-approve

FORCE:

sleep:
	timeout 5

up: role sleep apply

mac: role init
	sleep 5
	terraform apply -auto-approve

bump:
	$(eval VERSION=$(shell git describe --tags --abbrev=0 | awk -F. '{OFS="."; $$NF+=1; print $0}'))
	git tag -a $(VERSION) -m "new release"
	git push origin $(VERSION)

psbump:
	powershell -command "./bump.ps1"
