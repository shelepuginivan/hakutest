CGO_ENABLED := 1
export CGO_ENABLED

install: build-linux
	mv ./target/linux/hakuctl /usr/local/bin
	mv ./target/linux/hakutest /usr/local/bin

build: build-linux build-windows

build-linux:
	GOARCH=amd64 \
	GOOS=linux   \
	go build -trimpath -o ./target/linux/hakutest ./cmd/hakutest
	GOARCH=amd64 \
	GOOS=linux   \
	go build -trimpath -o ./target/linux/hakuctl ./cmd/hakuctl

build-windows:
	CC=x86_64-w64-mingw32-gcc                             \
	GOARCH=amd64                                          \
	GOOS=windows                                          \
	PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/lib/pkgconfig \
	go build -ldflags "-H=windowsgui" -trimpath -o ./target/windows/hakutest.exe ./cmd/hakutest
	CGO_ENABLED=0                                         \
	GOARCH=amd64                                          \
	GOOS=windows                                          \
	go build -trimpath -o ./target/windows/hakuctl.exe ./cmd/hakuctl

clean:
	go clean
	rm -rf target hakutest hakuctl

dev:
	gowatch

docs-build:
	yarn --cwd ./docs build

docs-dev:
	yarn --cwd ./docs dev

docs-format:
	yarn --cwd ./docs format

docs-preview:
	yarn --cwd ./docs preview

package-linux-appimage:
	mkdir -p ./target/hakutest.AppDir/usr/bin
	cp ./build/appimage/AppRun ./target/hakutest.AppDir
	cp ./build/resources/hakutest.desktop ./target/hakutest.AppDir/hakutest.desktop
	cp ./build/resources/hakutest.svg ./target/hakutest.AppDir
	cp ./target/linux/hakutest ./target/hakutest.AppDir/usr/bin
	ARCH=x86_64 appimagetool ./target/hakutest.AppDir
	mv ./Hakutest-x86_64.AppImage ./target/hakutest.AppImage

package-linux-deb:
	mkdir -p ./target/hakutest/DEBIAN
	cp ./build/deb/control ./target/hakutest/DEBIAN
	mkdir -p ./target/hakutest/usr/bin
	cp ./target/linux/hakutest ./target/hakutest/usr/bin
	cp ./target/linux/hakuctl ./target/hakutest/usr/bin
	mkdir -p ./target/hakutest/usr/share/applications
	cp ./build/resources/hakutest.desktop ./target/hakutest/usr/share/applications
	mkdir -p ./target/hakutest/usr/share/icons
	cp ./build/resources/hakutest.svg ./target/hakutest/usr/share/icons
	dpkg --build ./target/hakutest
	rm -r ./target/hakutest

package-linux-tarball:
	tar -czf ./target/hakutest-linux-x86_64.tar.gz --transform 's/^./hakutest/' -C ./target/linux .

package-windows-zip:
	cp -r ./target/windows ./target/hakutest
	cd ./target && zip hakutest-win-x86_64.zip -r hakutest
	rm -r ./target/hakutest

test:
	go test -cover ./...

version-patch:
	semver up release > ./pkg/version/VERSION

version-minor:
	semver up minor > ./pkg/version/VERSION

version-major:
	semver up major > ./pkg/version/VERSION

web: web-vendor web-minify

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
	mkdir -p ./web/vendor
	wget -qO ./web/vendor/alpine-3.14.1.min.js "https://cdn.jsdelivr.net/npm/alpinejs@3.14.1/dist/cdn.min.js"
