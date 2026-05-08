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

apply-wait: init
	go run ../../tools/retry-iam tofu apply -auto-approve

plan-wait: init
	go run ../../tools/retry-iam tofu plan

destroy-wait: init
	go run ../../tools/retry-iam tofu destroy -auto-approve

init:
	tofu init

upgrade:
	tofu init --upgrade

update: upgrade

role: FORCE
	tofu -chdir=./role apply -auto-approve

FORCE:

sleep:
	timeout 5

up: role sleep apply

mac: role init
	sleep 5
	tofu apply -auto-approve

bump:
	$(eval VERSION=$(shell git describe --tags --abbrev=0 | awk -F. '{OFS="."; $$NF+=1; print $0}'))
	git tag -a $(VERSION) -m "new release"
	git push origin $(VERSION)

psbump:
	powershell -command "./bump.ps1"
