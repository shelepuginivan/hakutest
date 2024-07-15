dev:
	gowatch

clean:
	go clean
	rm -rf ./target
	rm hakutest hakuctl

test:
	go test -cover ./...

web-minify:
	for f in web/css/*; do \
		minify -qo "$$f" "$$f" ; \
	done
	for f in web/js/*; do \
		minify -qo "$$f" "$$f" ; \
	done
	for f in internal/pkg/i18n/translations/*; do \
		minify -qo "$$f" "$$f" ; \
	done

web-vendor:
	wget -qO ./web/vendor/alpine-3.14.1.min.js "https://cdn.jsdelivr.net/npm/alpinejs@3.14.1/dist/cdn.min.js"
