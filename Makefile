GOCMD=go

test:
	cd tests; \
	$(GOCMD) test -v

release:
	npx github:escaletech/releaser --gpg-sign
