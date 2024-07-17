CGO_ENABLED := 1
export CGO_ENABLED

build: build-linux build-windows

build-linux:
	GOARCH=amd64 \
	GOOS=linux   \
	go build -trimpath -o ./target/linux/hakutest ./cmd/hakutest

build-windows:
	CC=x86_64-w64-mingw32-gcc                             \
	GOARCH=amd64                                          \
	GOOS=windows                                          \
	PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/lib/pkgconfig \
	go build -ldflags "-H=windowsgui" -trimpath -o ./target/windows/hakutest.exe ./cmd/hakutest

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
