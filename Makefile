dev:
	gowatch

clean:
	go clean
	rm -rf ./target
	rm hakutest hakuctl

minify:
	for f in web/css/*; do \
		minify -qo "$$f" "$$f" ; \
	done
	for f in web/js/*; do \
		minify -qo "$$f" "$$f" ; \
	done
	for f in internal/pkg/i18n/translations/*; do \
		minify -qo "$$f" "$$f" ; \
	done

test:
	go test -cover ./...
