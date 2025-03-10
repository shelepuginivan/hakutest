CGO_ENABLED := 1
export CGO_ENABLED

install: build-linux package-linux-prepare
	cp ./target/linux/usr/bin/{hakuctl,hakutest} /usr/local/bin
	cp ./target/linux/usr/share/icons/hakutest.svg /usr/local/share/icons
	mkdir -p /usr/local/share/licenses/hakutest
	cp ./target/linux/usr/share/licenses/hakutest/LICENSE.md /usr/local/share/licenses/hakutest
	cp ./target/linux/usr/share/applications/hakutest.desktop /usr/local/share/applications

all: test web build package
	rm -rf ./target/{hakutest.AppDir,linux,windows}
	git restore web internal

patch: version-patch all

minor: version-minor all

major: version-major all

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

package: package-linux-prepare package-linux-appimage package-linux-deb package-linux-tarball package-windows-zip

package-linux-prepare:
	mkdir -p \
		./target/linux/usr/bin \
		./target/linux/usr/share/icons \
		./target/linux/usr/share/applications \
		./target/linux/usr/share/licenses/hakutest
	mv ./target/linux/{hakutest,hakuctl} ./target/linux/usr/bin
	cp ./build/resources/hakutest.desktop ./target/linux/usr/share/applications
	cp ./build/resources/hakutest.svg ./target/linux/usr/share/icons
	cp ./LICENSE.md ./target/linux/usr/share/licenses/hakutest

package-linux-appimage:
	mkdir -p ./target/hakutest.AppDir/usr
	cp -r ./target/linux/usr/bin ./target/hakutest.AppDir/usr
	cp ./build/appimage/AppRun ./target/hakutest.AppDir
	cp ./build/resources/hakutest.desktop ./target/hakutest.AppDir/hakutest.desktop
	cp ./build/resources/hakutest.svg ./target/hakutest.AppDir
	ARCH=x86_64 appimagetool ./target/hakutest.AppDir
	mv ./Hakutest-x86_64.AppImage ./target/hakutest.AppImage

package-linux-deb:
	cp -r ./target/linux ./target/hakutest
	mkdir ./target/hakutest/DEBIAN
	cp ./build/deb/control ./target/hakutest/DEBIAN
	dpkg --build ./target/hakutest
	rm -r ./target/hakutest

package-linux-tarball:
	mkdir ./target/hakutest
	cp -r ./target/linux/usr/{bin,share} ./target/hakutest
	tar -czf ./target/hakutest-linux-x86_64.tar.gz --transform 's/^./hakutest/' -C ./target/hakutest .
	rm -rf ./target/hakutest

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
		go tool minify -qo "$$f" "$$f" ; \
	done
	for f in web/js/*; do \
		go tool minify -qo "$$f" "$$f" ; \
	done
	for f in internal/pkg/i18n/translations/*; do \
		go tool minify -qo "$$f" "$$f" ; \
	done

web-vendor:
	mkdir -p ./web/vendor
	wget -qO ./web/vendor/alpine-3.14.1.min.js "https://cdn.jsdelivr.net/npm/alpinejs@3.14.1/dist/cdn.min.js"
