REPO=github.com/sean9999/harebrain
SEMVER := $$(git tag --sort=-version:refname | head -n 1)

.PHONY: test

info:
	echo REPO is ${REPO} and SEMVER is ${SEMVER}

tidy:
	go mod tidy

test:
	go test ./...

clean:
	go clean

docs:
	pkgsite -open .

publish:
	GOPROXY=https://goproxy.io,direct go list -m ${REPO}@${SEMVER}
