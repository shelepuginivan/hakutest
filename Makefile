BINARY_NAME=hakutest
BUILD_DIR=./build
TARGET_DIR=./target

build: linux windows mac

linux: build-linux package-linux-appimage package-linux-deb package-linux-tarball cleanup-linux

build-linux:
	GOARCH=amd64 GOOS=linux go build -trimpath -o ${TARGET_DIR}/linux/${BINARY_NAME} ./cmd/hakutest
	GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -trimpath -o ${TARGET_DIR}/linux/${BINARY_NAME}-gtk ./cmd/hakutest-gtk

package-linux-appimage:
	mkdir -p ${TARGET_DIR}/hakutest.AppDir/usr/bin
	cp ${BUILD_DIR}/appimage/AppRun ${TARGET_DIR}/hakutest.AppDir
	cp ${BUILD_DIR}/resources/hakutest-gtk.desktop ${TARGET_DIR}/hakutest.AppDir/hakutest.desktop
	cp ${BUILD_DIR}/resources/hakutest.svg ${TARGET_DIR}/hakutest.AppDir
	cp ${TARGET_DIR}/linux/hakutest-gtk ${TARGET_DIR}/hakutest.AppDir/usr/bin
	NO_STRIP=true ${BUILD_DIR}/vendor/linuxdeploy-x86_64.AppImage --appdir ${TARGET_DIR}/hakutest.AppDir --plugin gtk --output appimage > /dev/null
	mv ./Hakutest-x86_64.AppImage ${TARGET_DIR}/hakutest.AppImage

package-linux-deb:
	mkdir -p ${TARGET_DIR}/${BINARY_NAME}/DEBIAN
	cp ${BUILD_DIR}/deb/control ${TARGET_DIR}/${BINARY_NAME}/DEBIAN
	cp ${BUILD_DIR}/deb/preinst ${TARGET_DIR}/${BINARY_NAME}/DEBIAN
	mkdir -p ${TARGET_DIR}/${BINARY_NAME}/usr/bin
	cp ${TARGET_DIR}/linux/${BINARY_NAME} ${TARGET_DIR}/${BINARY_NAME}/usr/bin
	cp ${TARGET_DIR}/linux/${BINARY_NAME}-gtk ${TARGET_DIR}/${BINARY_NAME}/usr/bin
	mkdir -p ${TARGET_DIR}/${BINARY_NAME}/usr/share/applications
	cp ${BUILD_DIR}/resources/hakutest-gtk.desktop ${TARGET_DIR}/${BINARY_NAME}/usr/share/applications
	mkdir -p ${TARGET_DIR}/${BINARY_NAME}/usr/share/icons
	cp ${BUILD_DIR}/resources/hakutest.svg ${TARGET_DIR}/${BINARY_NAME}/usr/share/icons
	dpkg --build ${TARGET_DIR}/${BINARY_NAME}

package-linux-tarball:
	tar -czf ${TARGET_DIR}/${BINARY_NAME}-linux-x86_64.tar.gz --transform 's/^./${BINARY_NAME}/' -C ${TARGET_DIR}/linux .

cleanup-linux:
	rm -rf ${TARGET_DIR}/${BINARY_NAME} ${TARGET_DIR}/${BINARY_NAME}.AppDir ${TARGET_DIR}/linux

windows:
	GOARCH=amd64 GOOS=windows go build -trimpath -o ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}.exe ./cmd/hakutest
	GOARCH=amd64 GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags "-H=windowsgui" -trimpath -o ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}-server.exe ./cmd/hakutest-server
	GOARCH=amd64 GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags "-H=windowsgui" -trimpath -o ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}-statistics.exe ./cmd/hakutest-statistics
	cp -r web ${TARGET_DIR}/${BINARY_NAME}
	cd ${TARGET_DIR}/${BINARY_NAME} && ../../scripts/generate-wxs.py
	cp assets/${BINARY_NAME}.ico ${TARGET_DIR}/${BINARY_NAME}
	cd ${TARGET_DIR}/${BINARY_NAME} && wixl ${BINARY_NAME}.wxs -o ../${BINARY_NAME}-win64.msi
	rm ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}.wxs ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}.ico
	cd ${TARGET_DIR} && zip ${BINARY_NAME}-win64.zip -r ${BINARY_NAME}
	rm -r ${TARGET_DIR}/${BINARY_NAME}

mac:
	GOARCH=amd64 GOOS=darwin go build -trimpath -o ${TARGET_DIR}/mac/${BINARY_NAME} ./cmd/hakutest
	cp -r web ${TARGET_DIR}/mac
	tar -czf ${TARGET_DIR}/${BINARY_NAME}-macos.tar.gz --transform 's/^./${BINARY_NAME}/' -C ${TARGET_DIR}/mac .
	rm -r ${TARGET_DIR}/mac

manual:
	mkdir -p ${TARGET_DIR}
	cp -r docs ${TARGET_DIR}/man
	gzip --recursive ${TARGET_DIR}/man
	tar -czf ${TARGET_DIR}/${BINARY_NAME}-manual.tar.gz --transform 's/^./man/' -C ${TARGET_DIR}/man .
	rm -r ${TARGET_DIR}/man

clean:
	go clean
	rm -r ${TARGET_DIR}

test:
	go test -cover ./...
