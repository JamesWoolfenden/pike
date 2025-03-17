.PHONY:
TEST?=$$(go list ./... | grep -v 'vendor'| grep -v 'scripts'| grep -v 'version')
HOSTNAME=jameswoolfenden
FULL_PKG_NAME=github.com/jameswoolfenden/pike
VERSION_PLACEHOLDER=version.ProviderVersion
NAMESPACE=dev
BINARY=pike
OS_ARCH=darwin_amd64
TERRAFORM=./terraform/
TF_TEST=./terraform_test/

default:

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

test:
	go test $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m


destroy:
	cd $(TERRAFORM) && terraform destroy --auto-approve


BIN=$(CURDIR)/bin
$(BIN)/%:
	@echo "Installing tools from tools/tools.go"
	@cat tools/tools.go | grep _ | awk -F '"' '{print $$2}' | GOBIN=$(BIN) xargs -tI {} go install {}

generate-docs:
	echo "does nowt"

docs:


vet:
	go vet ./...

bump:
	git push
	$(eval VERSION=$(shell git describe --tags --abbrev=0 | awk -F. '{OFS="."; $$NF+=1; print $0}'))
	git tag -a $(VERSION) -m "new release"
	git push origin $(VERSION)

psbump:
	git push
	powershell -command "./bump.ps1"

update:
	go get -u
	go mod tidy
	pre-commit autoupdate

lint:
	golangci-lint run --fix

gci:
	gci -w .

fmt:
	gofumpt -l -w .

schema:
	wget -qO- https://schema.cloudformation.us-east-1.amazonaws.com/CloudformationSchema.zip  |tar xvz -C ./src/schema
