.PHONY: clean  docs
TEST?=$$(go list ./... | grep -v 'vendor'| grep -v 'scripts'| grep -v 'version')
HOSTNAME=jameswoolfenden
FULL_PKG_NAME=github.com/jameswoolfenden/pike
VERSION_PLACEHOLDER=version.ProviderVersion
NAMESPACE=dev
BINARY=pike
OS_ARCH=darwin_amd64
TERRAFORM=./terraform/
TF_TEST=./terraform_test/

default: install

build:
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

check: install purge_tf purge_state
	cd $(TERRAFORM) && terraform init
	cd $(TERRAFORM) && terraform plan

check_tf: install
	cd $(TF_TEST) && rm .terraform.*
	cd $(TF_TEST) && terraform init
	cd $(TF_TEST) && terraform plan

apply: install purge_tf purge_state
	cd $(TERRAFORM) && terraform init
	cd $(TERRAFORM) && terraform apply --auto-approve

plan: install purge_tf
	cd $(TERRAFORM) && terraform init
	cd $(TERRAFORM) && terraform plan

update: install
	cd $(TERRAFORM) && terraform init
	cd $(TERRAFORM) && terraform apply --auto-approve

destroy:
	cd $(TERRAFORM) && terraform destroy --auto-approve

clean: purge_tf
	-rm -rf ./bin

purge_state:
	-rm $(TERRAFORM)terraform.tfstate
	-rm $(TERRAFORM)terraform.tfstate.backup

purge_tf:
	-rm -fr $(TERRAFORM).terraform
	-rm $(TERRAFORM).terraform.lock.hcl

BIN=$(CURDIR)/bin
$(BIN)/%:
	@echo "Installing tools from tools/tools.go"
	@cat tools/tools.go | grep _ | awk -F '"' '{print $$2}' | GOBIN=$(BIN) xargs -tI {} go install {}

generate-docs: $(BIN)/tfplugindocs
	-rm -fr templates-backup
	go run -ldflags="-X $(FULL_PKG_NAME)/$(VERSION_PLACEHOLDER)=$(shell git describe --tags --always --abbrev=0)" scripts/generate-docs.go -tfplugindocsPath=$(BIN)/tfplugindocs


docs: $(BIN)/tfplugindocs
	-rm -fr templates-backup
	go run scripts/generate-docs.go -tfplugindocsPath=$(BIN)/tfplugindocs

validate-docs: $(BIN)/tfplugindocs
	$(BIN)/tfplugindocs validate

fmt:
	go fmt ./...

vet:
	go vet ./...
