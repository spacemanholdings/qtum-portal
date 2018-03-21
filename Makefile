.PHONY: assets cli

assets:
	(cd ui && yarn build)
	go-bindata -o ui/assets.go -pkg ui -prefix ui/build ui/build

cli:
	GOOS=darwin go build -o build/qtumportal-mac github.com/kfichter/qtum-portal/cli/qtumportal
	GOOS=linux go build -o build/qtumportal-linux github.com/kfichter/qtum-portal/cli/qtumportal